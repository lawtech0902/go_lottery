/*
微信摇一摇
基础功能：
/lucky 只有一个抽奖的接口
*/

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// 奖品类型，枚举值iota
const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 不同的券
	giftTypeCouponFix        // 相同的券
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

type gift struct {
	id       int      // 奖品ID
	name     string   // 奖品名称
	pic      string   // 奖品图片
	link     string   // 奖品链接
	gtype    int      // 奖品类型
	data     string   // 奖品数据（特定的配置信息）
	dataList []string // 奖品数据集合（不同的优惠券的编码）
	total    int      // 总数，0 为无限
	left     int      // 剩余数量
	isUse    bool     // 是否可用
	rate     int      // 中奖概率，万分之 N, 0-9999
	rateMin  int      // 大于等于最小中奖编码
	rateMax  int      // 小于中奖编码
}

// 最大的中奖号码
const rateMax = 10000

var (
	logger   *log.Logger
	giftList []*gift // 奖品列表
	mu       sync.Mutex
)

type lotteryController struct {
	Ctx iris.Context
}

// 初始化日志
func initLog() {
	f, _ := os.Create("./demo/3wechatShake/log/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

// 初始化奖品
func initGift() {
	giftList = make([]*gift, 5)

	giftList[0] = &gift{
		id:       1,
		name:     "手机大奖",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealLarge,
		data:     "",
		dataList: nil,
		total:    10,
		left:     10,
		isUse:    true,
		rate:     10,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[1] = &gift{
		id:       2,
		name:     "充电器",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealSmall,
		data:     "",
		dataList: nil,
		total:    5,
		left:     5,
		isUse:    true,
		rate:     10,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[2] = &gift{
		id:       3,
		name:     "优惠券满200减50元",
		pic:      "",
		link:     "",
		gtype:    giftTypeCouponFix,
		data:     "mall-coupon-2019",
		dataList: nil,
		total:    50,
		left:     50,
		isUse:    true,
		rate:     500,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[3] = &gift{
		id:       4,
		name:     "直降优惠券50元",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoupon,
		data:     "",
		dataList: []string{"c01", "c02", "c03", "c04", "c05"},
		total:    10,
		left:     10,
		isUse:    true,
		rate:     100,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[4] = &gift{
		id:       5,
		name:     "虚拟币",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoin,
		data:     "10金币",
		dataList: nil,
		total:    5,
		left:     5,
		isUse:    true,
		rate:     5000,
		rateMin:  0,
		rateMax:  0,
	}

	// 数据整理，中奖区间数据
	rateStart := 0
	for _, data := range giftList {
		if !data.isUse {
			continue
		}

		data.rateMin = rateStart
		data.rateMax = rateStart + data.rate

		if data.rateMax >= rateMax {
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
}

func newApp() *iris.Application {
	app := iris.Default()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()
	initGift()

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 奖品数量信息 GET http://localhost:8080/
func (c *lotteryController) Get() string {
	count := 0
	total := 0
	for _, data := range giftList {
		if data.isUse && (data.total == 0 || data.total > 0 && data.left > 0) {
			count++
			total += data.left
		}
	}

	return fmt.Sprintf("当前有效奖品种类数量：%d，限量奖品总数量：%d\n", count, total)
}

// 抽奖 GET http://localhost:8080/lucky
func (c *lotteryController) GetLucky() map[string]interface{} {
	mu.Lock()
	defer mu.Unlock()

	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok

	for _, data := range giftList {
		if !data.isUse || (data.total > 0 && data.left <= 0) {
			continue
		}

		// 中奖了，抽奖编码在奖品编码范围内
		if data.rateMin <= int(code) && int(code) <= data.rateMax {
			// 开始发奖
			sendData := ""
			switch data.gtype {
			case giftTypeCoin:
				ok, sendData = sendCoin(data)
			case giftTypeCoupon:
				ok, sendData = sendCoupon(data)
			case giftTypeCouponFix:
				ok, sendData = sendCouponFix(data)
			case giftTypeRealSmall:
				ok, sendData = sendRealSmall(data)
			case giftTypeRealLarge:
				ok, sendData = sendRealLarge(data)
			}

			if ok {
				// 中奖成功，成功得到奖品，生成中奖记录
				saveLuckyData(code, sendData, data)
				result["success"] = ok
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				break
			}
		}
	}

	if v, ok := result["success"]; ok && v == false {
		result["data"] = "没有中奖"
	}

	return result
}

// 返回一个随机数
func luckyCode() int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(rateMax)
}

// 虚拟币
func sendCoin(g *gift) (bool, string) {
	if g.total == 0 {
		// 数量无限
		return true, g.data
	} else if g.left > 0 {
		// 仍有剩余
		g.left--
		return true, g.data
	} else {
		return true, "奖品已发完"
	}
}

// 不同值的优惠券
func sendCoupon(g *gift) (bool, string) {
	if g.left > 0 {
		// 仍有剩余
		g.left--

		if g.left > len(g.dataList) {
			return true, "奖品已发完"
		}
		return true, g.dataList[g.left]
	}
	return true, "奖品已发完"
}

var sendCouponFix = sendCoin
var sendRealSmall = sendCoin
var sendRealLarge = sendCoin

// 记录用户获奖信息
func saveLuckyData(code int32, sendData string, g *gift) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s, left=%d\n",
		code, g.id, g.name, g.link, sendData, g.left)
}
