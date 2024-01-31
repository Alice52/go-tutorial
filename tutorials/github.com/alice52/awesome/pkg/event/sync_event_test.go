package event

import (
	"fmt"
	"github.com/gookit/event"
	"testing"
	"time"
)

var listener event.ListenerFunc = func(e event.Event) error {
	fmt.Printf("handle event: %s\n", e.Name())
	return nil
}

func TestQuickStart(t *testing.T) {
	// 1. 定义处理器
	hanFunc := listener

	// 2. 注册事件监听器
	event.On("evt1", event.ListenerFunc(hanFunc), event.Normal)
	event.On("evt1", event.ListenerFunc(hanFunc), event.High)

	// 3.  触发事件
	event.MustFire("evt1", event.M{"arg0": "val0", "arg1": "val1"})
	event.FireEvent(event.NewBasic("evt1", event.M{"arg0": "val0", "arg1": "val1"}))
}

// Deprecated TestModeSimple 是默认模式, 注册事件监听器和名称以通配符 * 结尾
// "*" only allow one and must at end
func TestModeSimple(t *testing.T) {
	// 1. 定义处理器
	hanFunc := listener

	// 2. 注册事件监听器
	event.On("app.db.*", hanFunc, event.Normal)

	// 3. Trigger event
	event.MustFire("app.db.create", event.M{"arg0": "val0", "arg1": "val1"})
	event.MustFire("app.db.update", event.M{"arg0": "val0"})
}

// TestModePath
// * 只匹配一段非 . 的字符, 可以进行更精细的监听匹配
// ** 则匹配任意多个字符, 并且只能用于开头或结尾
func TestModePath(t *testing.T) {
	// 1. 定义处理器
	hanFunc := listener

	// 2. 创建事件管理器: 默认是 default
	em := event.NewManager("test", event.UsePathMode)
	// event.Config(event.UsePathMode)

	// 3. 注册事件监听器
	em.On("app.**", hanFunc)
	em.On("app.db.*", hanFunc)
	em.On("app.*.create", hanFunc)
	em.On("app.*.update", hanFunc)

	// 4. 触发事件: 只有 app.** | app.db.* | app.*.create 会被触发
	em.MustFire("app.db.create", event.M{"arg0": "val0", "arg1": "val1"})
}

// TestConcurrentEventAsync
// 同步下只有一个协程在执行任务: 内部指向handle之后在触发下一个任务
func TestConcurrentEventSync(t *testing.T) {
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
	em.On("app.*", listener, event.Normal)

	// 4. 触发事件(同步)
	for i := 0; i < 10; i++ {
		if em.FireEvent(event.New("app.evt1", event.M{"arg0": "val2"})) != nil {
			fmt.Println("occurs error when handle event!")
		}
	}

	fmt.Println("publish event finished!")
}
