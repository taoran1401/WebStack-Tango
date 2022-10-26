package command

import (
	"fmt"
	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

type CollyCmd struct {
	Log *zap.SugaredLogger
}

func NewCollyCmd(log *zap.SugaredLogger) *CollyCmd {
	return &CollyCmd{Log: log}
}

func (this *CollyCmd) Handle() {
	url := "https://blog.csdn.net/Xiang_lhh/article/details/113769478"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		//请求之前调用
		fmt.Println("请求之前调用")
		//过滤url，只访问符合正则表达式的url
		//colly.URLFilters()
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("请求过程中发生错误调用")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("收到响应后调用")
	})

	/*c.OnHTML("a", func(e *colly.HTMLElement) {
		//fmt.Println("收到html之后调用")

		fmt.Println(e.Attr("href"))
	})*/

	c.OnHTML("#rewardNew .reward-close img", func(e *colly.HTMLElement) {
		//fmt.Println("收到html之后调用")

		fmt.Println(e.Attr("src"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("回调之后调用")
	})

	c.Visit(url)
}
