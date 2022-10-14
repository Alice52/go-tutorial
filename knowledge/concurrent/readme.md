## 相关概念

1. 串行 & 并发 & 并行
2. 进程 & 线程 & 协程(纤程)
3. 并发模型

   - 线程 & 锁模型: `通过共享内存通信`
   - Actor 模型
   - **CSP 模型**: `通过通信共享内存`
   - Fork&Join 模型

4. goroutine: `用户态+小栈`

   - 原理

5. channel

   - 缓存
   - 方向
   - 使用(关闭 & range)

6. select

   - 多路复用
   - 优先级

7. 并发

   - 互斥锁: sync.Mutex
   - 读写锁: sync.RWMutex
   - sync.waitgroup: CountDownLatch
   - 执行一次: sync.Once
   - sync.Map: 线程安全的 Map
   - 原子包: `sync/atomic`
