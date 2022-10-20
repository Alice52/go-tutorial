package redis

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func TestOp(t *testing.T) {
	// 1. mock一个redis server
	mock, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mock.Close()

	// 2. 准备数据
	mock.Set("q1mi", "liwenzhou.com")
	mock.SAdd(KeyValidWebsite, "q1mi")

	// 3. 连接mock的redis server
	rdb := redis.NewClient(&redis.Options{
		Addr: mock.Addr(), // mock redis server的地址
	})

	// 调用函数
	ok := GetKey(rdb, "q1mi")
	if !ok {
		t.Fatal()
	}

	// 可以手动检查redis中的值是否复合预期
	if got, err := mock.Get("blog"); err != nil || got != "liwenzhou.com" {
		t.Fatalf("'blog' has the wrong value")
	}
	// 也可以使用帮助工具检查
	mock.CheckGet(t, "blog", "liwenzhou.com")

	// 过期检查
	mock.FastForward(5 * time.Second) // 快进5秒
	if mock.Exists("blog") {
		t.Fatal("'blog' should not have existed anymore")
	}
}
