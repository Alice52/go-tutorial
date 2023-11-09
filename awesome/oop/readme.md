## go 中面向对象编程的实践

1. 对象 - go 没有 class 但是可以使用 struct 代替

   ```go
   type Person struct {
     Name string
     Age uint8
   }
   ```

2. 构造函数

   ```go
   //  struct 是值类型
   // 一般比较复杂, 值拷贝开销大, 所以一般返回指针
   func NewPerson(name string, age uint8) *Person {
     return &Person{
       Name: name,
       Age:  age,
     }
   }
   ```

3. 定义行为方法: **函数接收方**

   - 非本地类型不能定义方法: 不能给别的包的类型定义方法

   ```go
   func (p *Person) Saying(word string) *Person {
     fmt.Printf("%s say %s", p.name, word)
     return p
   }
   ```

4. 对象嵌套 & 匿名字段

   ```go
   // 匿名字段默认使用类型名作为字段名
   func AsAnonymousFile() {
     type User struct {
       Name    string
       Gender  string
       Address // 匿名字段
     }

     var user2 User
     user2.Address.province = "山东"
     user2.province = "山东" // 匿名字段可以省略
   }
   ```

5. `{go 为了简介没有重载}`[重载](https://juejin.cn/post/7109097301574385701):

   - [不好]通过可变参数的 any 实现: 缺失约束
   - [可以]通过定义结构体做参数实现: 参数使用复杂
   - [推荐]Options 模式: struct(setXxx) + **可变**函数做参数

   ```go
   // 01. 需求
   RpcMethod()
   RpcMethod(timeout int)
   RpcMethod(timeout int, host string)
   RpcMethod(timeout int, host string, cluster string)

   // 02. 解决: 不好(约束缺失)
   func RpcMethod(op ...interface{}) {}

   // 03. 解决: 不好
   type RpcConfig struct {
     timeout time.Duration
     cluster string
     host    string
   }
   func Handler(op *Op) {}

   // 04. 解决: 推荐
   type RpcConfig struct {
     timeout time.Duration
     cluster string
     host    string
   }

   // func as setXx
   func Timeout(duration time.Duration) Option {
     return func(config *RpcConfig) {
       config.timeout = duration
     }
   }
   func Cluster(cluster string) Option {
     return func(config *RpcConfig) {
       config.cluster = cluster
     }
   }
   func Host(host string) Option {
     return func(config *RpcConfig) {
       config.host = host
     }
   }

   func RpcMethod(ops ...Option) {
     // build args by executing args-func
     var rpcConf RpcConfig
     for _, op := range ops {
       op(&rpcConf)
     }
     // invoke rpc method
   }

   func OverrideUsage() {
     RpcMethod()
     RpcMethod(Timeout(100))
     RpcMethod(Timeout(100), Host("127.0.0.1"))
   }
   ```

6. `{go 为了简介没有继承}`对象继承: `通过嵌套(组合)匿名结构体实现继承`

   ```go
   type Student struct {
     *Person // 匿名嵌套对象(实现继承)
     school  string
     grade   string
   }

   func (s *Student) Learning() {
     fmt.Printf("%s is learning.", s.name)
   }

   // 继承了属性及相关方法
   func TestInherit(t *testing.T) {
      stu := &Student{
        Person: p1,
        grade:  "2",
        school: "nt",
      }

      stu.Saying("student")
      stu.Learning()
   }
   ```

7. 接口 & _实现_

   - 接口一组接口的集合, 是一种抽象类型
   - 实现接口的所有方法就是实现类: duck typing

     ```go
     // 1. 定义接口
     type Flyable interface {
       Fly()
     }

     // 2. 定义对象 & 对象实现接口
     type Bird struct {
       Name string
     }
     func (b *Bird) Fly() {
       fmt.Printf("%v is flying\n", b.Name)
     }

     // 4. compiler checker
     var _ Flyable = &Bird{}

     func test() {
       // 5. 使用接口接受实现
       var bird Flyable = &Bird{
         Name: "bird",
       }
       bird.Fly()
     }
     ```

8. 泛型(实现)-应用(继承)

   - 泛型

     ```go
     // type Set[T Animal | Dog] interface {
     type Set[T any] interface {
       Put(val T)
       All() []T
     }

     type HashSet[T any] struct {
       Elements []T
     }
     func (t *HashSet[T]) Put(val T) {
       fmt.Printf("call method of put(%v)\n", val)
       t.Elements = append(t.Elements, val)
     }
     func (t *HashSet[T]) All() []T {
       fmt.Printf("call method of all(): %v\n", t.Elements)
       return t.Elements
     }
     ```

   - 应用

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

     // 4. 使用
        // 4.1 创建容器
     var set Set[Animal] = &HashSet[Animal]{}
        // 4.2 向Set中添加Dog(匿名属性初始化)
     dog := &Dog{
       &Nameable{
         Name: "dog",
       },
     }
        // 4.3 使用
     set.Put(dog)
     ```
