## 定义变量

1. 基本类型变量

   ```go
   var[int] a int  // declare and init as 0
   var[int] b int = 100  // declare and init as 0
   var[int] c = 100
   d := 100  // 只能在函数内
   v1, _ := foo() // 匿名变量(不占用空间, 不分配内存)
   ```

2. 数组变量: 复制后修改彼此不影响

   - `*new`: 值
   - 无长度初始化: `[2]int{2, 5}`
   - 可变长度初始化
   - 声明: 会被初始化为 zero-value

   ```go
   var a1 [2]int // declare and init as zero
   var a2 = [2]int{1,2}                        // declare and init
   var a3 = [5]string{3: "ab", 4: "cd"}
   var a4 = [...]int{1,2}                      // variable-length array declaration
   var a5 = *new([5]int) // a5 is array
   ```

3. slice 变量: 定义声明时无长度

   - make: 返回引用本身
   - new: 返回指针
   - 无长度初始化: 对象本身, `[]int{2, 5}`
   - 截取数组
   - 声明: 会被初始化为 nil

   ```go
   var s1 []int // declare and init as nil
   var s3 []int = make([]int, 0, 10)
   var s3 = make([]int, 0, 10)
   // s6 := *new([]int) // init as nil
   var s6 = new([5]int) // s6 is slice and pointer: 修改后都会变

   s4 := a2[1:4]
   s5 := []int{2, 5}
   ```

4. map 变量: `map[KeyType]ValueType`

   - 声明
   - 初始化: 对象本身
   - make: 引用本身
   - ~~new: 返回指针~~

   ```go
   // 1. declare
   var nmap map[string]int // value is nil

   // 2. declare and init
   var nmap map[string]int = make(map[string]int, 10)
   nmap["zack"] = 100

   // 3. init
   userInfo := map[string]string{
       "username": "prof.cn",
       "password": "123456",
   }

   // 4. 任意类型
   any := map[string]interface{}{
      "name": "zack",
      "age":  18,
   }
   ```

5. struct 变量

   - `var 结构体实例 结构体类型`
   - new 会直接分配为 z-v: **因为值 struct 是值类型**
   - 初始化: 值 || 指针

     ```go
     // 1. var
     var person Person
     person.age = 15
     person.name = "zack"

     // 2. pointer
     p := new(Person)
     p.name = "kayla"
     p.age = 15

     // 3. initialize
     p5 := Person{
         name: "prof.cn",
         city: "BeiJing",
         age:  18,
     }
     p6 := &p5

     // 4. initialize and pointer
     p6 := &Person{
         "prof.cn",
         "BeiJing",
         18,
     }
     ```

6. 函数变量

   - 类型: `type cal func(int, int) int`
   - 变量: `var cal func(int, int) int`
   - 参数: `cal func(int, int) int`
   - 返回值: `func(int, int) int`

7. 指针变量
8. _接口变量_

## 初始化

1. var: 默认会被初始化为 zero-value

   - map 会被初始化为 nil
   - `*指针` 会被初始化为 nil

2. `func init()` 会被优先执行

## 任意类型 Map

1. 初始化

   ```go
   m := map[string]interface{} {
   "name": "zack",
   "age": 18
   }
   ```

2. 使用

   ```go
   // fing
   v, ok := m["name"]
   if ok {
     log.Println(v)
   } else {
     log.Println("not find")
   }

   // travse
   for k, v := range m {
     log.Println(k, v)
   }
   ```

3. `interface{}`

   ```go
   switch value := v.(type) {
   case string:
   	fmt.Printf("string type, and value is %s", value)
   default:
   	fmt.Println(value)
   }
   ```

## 继承

1. go 语言本身为了简洁性, 所以不提供对继承的支持
2. 使用组合代替继承

   ```go
   package main

   import "fmt"

   type Skills  []string

   type person struct {
     name string
     age  int
     weight int
   }

   type Student struct {
     person    //继承
     Skills
     int
     spe string
   }

   func main() {
     //方式一,全部指定
     xuxu := Student{person{"xuxu",25,68}, []string{"anatomy"}, 1, "boy"}
     //方式二,指哪打哪
     jane := Student{person:person{"Jane",25,100}, spe:"Biology"}

     fmt.Printf("His name is %s", jane.name)
     fmt.Printf("His name is %s", xuxu.name)
   }
   ```

## 重载

1. go 语言本身为了简洁性, 所以不提供对重载的支持
2. 一般解决方法

   ```go
   // 01. 需求
   func Handler()  {}
   func Handler(timeOut time.Duration) {}
   func Handler(timeOut time.Duration, retry int) {}

   // 02. 解决: 不好
   func Handler(op ...interface{}) {}

   // 03. 解决: 不好
   type Op struct {
       TimeOut time.Duration
       Retry   int
   }
   func Handler(op *Op) {}

   // 04. 解决: 推荐
   type Options struct {
       TimeOut     time.Duration
       RetryMaxNum int
   }

   type Option func(*Options)
   func loadOp(option ...Option) *Options {
       options := new(Options)
       for _, e := range option {
           e(options)
       }
       return options
   }

   func Handler(option ...Option) {
       op := loadOp(option ...)
   }

   func main() {
       Handler()
       Handler(func(options *Options) { options.TimeOut = time.Millisecond})

       Handler(func(options *Options) { options.RetryMaxNum = 1 })

       Handler(func(options *Options) { options.RetryMaxNum = 1 }, func(options *Options) { options.TimeOut = time.Millisecond })
   }
   ```

## printf

1. %c: 字符
2. %s: 字符创
3. %d: 数字
4. %T: 类型
5. %x: 16 进制
6. %o: 8 进制
7. %b: 2 进制
8. %v: 变量值
9. %p: 变量地址
