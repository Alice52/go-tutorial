## go-tutorial

1. This repository records golang learning process
2. pros: 简洁 & 高效

## blemish

1. without syntax annotation
2. without aop/validate
3. without named function and default args
4. without agent enhance
5. 在 Go 语言中, 方法的选择是在编译时进行的, 而不是在运行时(优先考虑传接口而不是对象)

   - caller 接口的方法集在编译时就已经确定, 而不会根据实例的动态类型变化(既要实现约束-又要调用子类方法)

     ```go
     type LimitInterface interface {
       GenerationKey(c *gin.Context) string
       CheckOrMark(key string, expire int, limit int) error
     }

     type LimitBase struct {
       LimitInterface
     }

     func (l *LimitBase) Process(xxx) string {
       l.CheckOrMark(xxx) // 此时会调用LimitBase的CheckOrMark(没有实现-报错): 解决办法是将其作为函数传递进来
     }
     ```

     - 通过传递函数调用子类的实现

       ```go
       type CheckFn = func(key string, expire int, limit int) (err error)
       func (l *LimitBase) Process(subCheck CheckFn, xxxx) string {
         subCheck(xxxx) // 此时调用传递进来的实现
       }

       func main() {
        // 子类定义 LocalLimit
        type LocalLimit struct { // todo: 实现相关接口
          LimitBase
        }
        localLimit.Process(limiter.CheckOrMark, xxxx)
       }
       ```

   - 方法参数可以传递接口的实现对象调用

     ```go
     type Allower interface {
       Allow() bool
     }
     func NewErrorLimiter(limit Allower)  {
       limit.Allow()  // 此时会调用传进来的实现
     }

     func main() {
       limit := rate.NewLimiter(rate.Every(time.Minute), 1) // 实现了Allower
       NewErrorLimiter(limit) // 此时会调用 NewLimiter 的 Allow() 方法
     }
     ```

6. confuse parse

   - 转换

     ```go
     type FuncJob func()
     func (f FuncJob) Run() { f() }

     func (c *Cron) AddFunc(spec string, , name string) {
       c.AddJob(spec, FuncJob(cmd), name)  // FuncJob(cmd) 是将 cmd 转换为 FuncJob 类型
     }
     ```

   - 接口实现约束

     ```go
     // 约束 T 实现了 Trace
     var _ T = (*Trace)(nil) // 将 nil 强制转换为 *Trace 类型(T类型)

     type T interface {}
     type Trace struct {}
     ```

## core point

1. func can be use as type, arg, var, return

   ```go
   // 1. as type
   type Op func(int, int) int

   // 2. as arg
   func calc(x, y int, op Op) Op {

     // 3. as var
     opFunc := op
     opFunc(x, y)

     return func(a, b int) int {
       return a + b
     }
   }

   func main() {
     // as arg
     add := func(a, b int) int {
       return a + b
     }

     // excute func
     add(1, 2)
     // pass as arg
     calc(10, 20, add)
   }
   ```

2. oop & option 模式

   ```go
   // 1. 接口
   type Animal interface {
     GetName()
   }

   // 2. 接口实现
   type Nameable struct {
     Name string
   }
   func (t *Nameable) GetName() {
     fmt.Printf("name: %v\n", t.Name)
   }

   // 3. 继承
   type Dog struct {
     *Nameable
   }
   var _ Animal = &Dog{}  // impl check for compiler
   ```

3. generic
