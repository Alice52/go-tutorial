package goroutine

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func DeferRecoverError() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover inner panic:%v\n", r)
			}
		}()
		fmt.Println("in goroutine....")
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("exit")
}

func GroupConError() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"http://www.yixieqitawangzhi.com",
		"http://www.liwenzhou.com",
	}
	for _, url := range urls {
		url := url // 注意此处声明新的变量
		// 第一个返回非零错误的调用将取消该Group
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("所有 goroutine 均成功")
	return nil
}
