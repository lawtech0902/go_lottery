package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type gift struct {
	id      int    // 奖品ID
	name    string // 奖品名称
	pic     string // 奖品图片
	link    string // 奖品链接
	isUse   bool   // 是否可用
	rate    int    // 中奖概率，万分之 N, 0-9999
	rateMin int    // 大于等于最小中奖编码
	rateMax int    // 小于中奖编码
}

// 最大的中奖号码
const rateMax = 10

// 初始化奖品
func newGift() *[5]gift {
	giftList := new([5]gift)

	giftList[0] = gift{
		id:      1,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[1] = gift{
		id:      2,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[2] = gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[3] = gift{
		id:      4,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[4] = gift{
		id:      5,
		name:    "敬业福",
		pic:     "敬业福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	return giftList
}

func giftRange(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rate)

	rateStart := 0
	for i := range giftList {
		if !giftList[i].isUse {
			continue
		}

		grate := 0

		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}

		giftList[i].rate = grate

		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate

		if giftList[i].rateMax >= rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	return giftList
}

type lotteryController struct {
	Ctx iris.Context
}

var logger *log.Logger

// 初始化日志
func initLog() {
	f, _ := os.Create("./demo/4alipayFu/log/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.Default()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 奖品数量信息 GET http://localhost:8080/
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	giftList := giftRange(rate)

	result := ""
	for _, data := range giftList {
		result += fmt.Sprintf("%v\n", data)
	}

	return result
}

// 抽奖 GET http://localhost:8080/lucky
func (c *lotteryController) GetLucky() map[string]interface{} {
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	code := luckyCode()
	giftList := giftRange(rate)

	result := map[string]interface{}{}
	result["success"] = false

	for _, data := range giftList {
		if !data.isUse {
			continue
		}

		// 中奖了，抽奖编码在奖品编码范围内
		if data.rateMin <= int(code) && int(code) <= data.rateMax {
			sendData := data.pic
			saveLuckyData(code, sendData, &data)
			result["success"] = true
			result["uid"] = uid
			result["id"] = data.id
			result["name"] = data.name
			result["link"] = data.link
			result["data"] = sendData
			break
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

// 记录用户获奖信息
func saveLuckyData(code int32, sendData string, g *gift) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s\n",
		code, g.id, g.name, g.link, sendData)
}
