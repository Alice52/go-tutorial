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
