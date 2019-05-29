/*
curl http://localhost:8080
curl --data "users=law1,law2" http://localhost:8080/import
curl http://localhost:8080/lucky
*/

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	userList []string
	mu       sync.Mutex
)

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	var (
		app *iris.Application
	)

	app = iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	mu = sync.Mutex{}
	return app
}

func main() {
	var (
		app *iris.Application
	)

	app = newApp()
	userList = []string{}

	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	var (
		count int
	)

	count = len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数：%d\n", count)
}

/*
POST http://localhost:8080/import
params: users

curl --data "users=law1,law2"
*/
func (c *lotteryController) PostImport() string {
	var (
		strUsers string
		users    []string
		count1   int
		count2   int
		u        string
	)

	strUsers = c.Ctx.FormValue("users")
	users = strings.Split(strUsers, ",")

	mu.Lock()
	defer mu.Unlock()

	count1 = len(userList)

	for _, u = range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}

	count2 = len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数：%d，成功导入的用户数：%d\n", count1, count2)
}

/*
GET http://localhost:8080/lucky
*/
func (c *lotteryController) GetLucky() string {
	var (
		count int
		seed  int64
		index int32
		user  string
	)

	mu.Lock()
	defer mu.Unlock()

	count = len(userList)
	if count > 1 {
		seed = time.Now().UnixNano()
		index = rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user = userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else if count == 1 {
		user = userList[0]
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintf("已经没有用户参与，请先通过 /import 导入用户\n")
	}
}
