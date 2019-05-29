/*
压力测试 wrk -t10 -c100 -d5 "http://localhost:8080/prize"
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type lotteryController struct {
	Ctx iris.Context
}

// 奖品中奖概率
type Prate struct {
	Rate  int    // 万分之N的中奖概率
	Total int    // 总数量限制，0表示无数量限制
	CodeA int    // 中奖概率起始编码
	CodeB int    // 中奖概率终止编码
	Left  *int32 // 剩余数
}

// 奖品列表
var prizeList = []string{
	"一等奖，火星单程船票",
	"二等奖，凉飕飕南极之旅",
	"三等奖，iPhone一部",
	"", // 没有中奖
}

var left int32 = 1000

// 奖品的中奖概率设置，与上面的prizeList对应的设置
var rateList = []Prate{
	Prate{100, 1000, 0, 9999, &left}, // 压力测试配置
	//Prate{1, 1, 0, 0, 1},
	//Prate{2, 2, 1, 2, 2},
	//Prate{5, 10, 3, 5, 10},
	//Prate{100, 0, 0, 9999, 0},
}

var logger *log.Logger

// 初始化日志
func initLog() {
	f, _ := os.Create("./demo/6wheel/log/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.Default()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 首页 奖品信息 GET http://localhost:8080/
func (c *lotteryController) Get() string {
	c.Ctx.Header("Content-type", "text/html")
	return fmt.Sprintf("大转盘奖品列表：<br/>%s", strings.Join(prizeList, "<br/>"))
}

// 奖品概率 GET http://localhost:8080/debug
func (c *lotteryController) GetDebug() string {
	c.Ctx.Header("Content-type", "text/html")
	return fmt.Sprintf("获奖概率：<br/>%s",
		strings.Join(func(rateList []Prate) []string {
			var res []string
			for _, v := range rateList {
				res = append(res, fmt.Sprintf("%+v", v))
			}
			return res
		}(rateList), "<br/>"))
}

// 抽奖 GET http://localhost:8080/prize
func (c *lotteryController) GetPrize() string {
	// 第一步，抽奖，根据随机数匹配奖品
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Intn(10000)

	var (
		myPrize   string
		prizeRate *Prate
	)

	// 从奖品列表匹配是否中奖
	for i, prize := range prizeList {
		rate := &rateList[i]

		// 满足中奖条件
		if code >= rate.CodeA && code <= rate.CodeB {
			myPrize = prize
			prizeRate = rate
			break
		}
	}

	if myPrize == "" {
		return "很遗憾，再来一次吧"
	}

	// 第二步，中奖后开始发奖
	// 无限量奖品
	if prizeRate.Total == 0 {
		//logger.Print("奖品", myPrize)
		fmt.Println("奖品", myPrize)
		return myPrize
	} else if *prizeRate.Left > 0 {
		left := atomic.AddInt32(prizeRate.Left, -1)

		if left >= 0 {
			logger.Println("奖品", myPrize)
			return myPrize
		} else {
			return "很遗憾，再来一次吧"
		}
	} else {
		return "很遗憾，再来一次吧"
	}
}
