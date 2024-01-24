// chan
package event

import (
	"fmt"
	"github.com/gookit/event"
	"testing"
	"time"
)

func TestChanEvent(t *testing.T) {
	defer event.CloseWait()

	// 1. 注册事件监听器
	event.On("app.evt1", listener, event.Normal)

	// 2. 触发事件(异步)
	event.Async("app.evt1", event.M{"arg0": "val0", "arg1": "val1"})
	event.FireAsync(event.New("app.evt1", event.M{"arg0": "val2"}))
}

func TestChanEvent1(t *testing.T) {
	defer event.CloseWait()

	// 1. 定义事件管理器
	event.Config(func(o *event.Options) {
		o.ChannelSize = 0
		o.ConsumerNum = 10
	})

	// 2. 注册事件监听器
	event.On("app.evt1", listener, event.Normal)

	// 3. 触发事件(异步)
	event.Async("app.evt1", event.M{"arg0": "val0", "arg1": "val1"})
	event.FireAsync(event.New("app.evt1", event.M{"arg0": "val2"}))
}

func TestChanEvent2(t *testing.T) {
	// 1. 定义事件管理器
	var em = event.NewManager("default", func(o *event.Options) {
		o.ConsumerNum = 10
		o.EnableLock = false // 加锁后只有一个 routine 在执行
	})
	defer em.CloseWait()

	// 2. 定义事件处理器
	var listener event.ListenerFunc = func(e event.Event) error {
		time.Sleep(1 * time.Second)
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}

	// 3. 注册事件监听器
	em.On("app.evt1", listener, event.Normal)

	// 4. 触发事件(异步)
	for i := 0; i < 100; i++ {
		em.FireAsync(event.New("app.evt1", event.M{"arg0": "val2"}))
	}

	fmt.Println("publish event finished!")
}
