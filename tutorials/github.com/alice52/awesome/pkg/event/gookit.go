// Package event: https://github.com/gookit/event
package event

/**
1. On/Listen(name string, listener Listener, priority ...int) 注册事件监听
2. Subscribe/AddSubscriber(sbr Subscriber) 订阅，支持注册多个事件监听
3. Trigger/Fire(name string, params M) (error, Event) 触发事件
4. MustTrigger/MustFire(name string, params M) Event 触发事件，有错误则会panic
5. FireEvent(e Event) (err error) 根据给定的事件实例，触发事件
6. FireBatch(es ...interface{}) (ers []error) 一次触发多个事件
7. Async/FireC(name string, params M) 投递事件到 chan，异步消费处理
8. FireAsync(e Event) 投递事件到 chan，异步消费处理
9. AsyncFire(e Event) 简单的通过 go 异步触发事件
*/
