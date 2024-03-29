## important

1. **for 循环下的 v 值是不一样的, 但是地址是一直复用的**

   - 因此, 尽量不要使用其引用`&`

   ```go
   m := make(map[string]*student)
   students := []student{
       {name: "prof.cn", age: 18},
       {name: "testing", age: 23},
       {name: "blog", age: 28},
   }

   for _, stu := range students {
       m[stu.name] = &stu
   }
   for k, v := range m {
       // nmae will always return blog
       fmt.Println(k, "=>", v.name)
   }
   ```

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

     // 5. empty struct: 不占内存, 不分配空间
     var es struct{} // es is nil
     c := &es{} // c is not nil

     // 6. 匿名结构体: 临时数据结构
     func AnonymousStruct() {
        var user struct {
            Name string
            Age  int
        }
        user.Name = "zack"
        user.Age = 15
     }

     // 7. 批量定义 + 结构体标签
     type (
         userpw struct {
             UserName     string `json:"username"`
             Password     string `json:"password"`
             RefreshToken string `json:"refreshToken"`
         }
     )
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
2. 使用组合代替继承: `通过嵌套匿名结构体实现继承`
3. [sample](/tutorials/cn.edu.ntu.awesome/syntax/oop/oop.go)

## 重载

1. go 语言本身为了简洁性, 所以不提供对重载的支持
2. 一般解决方法

## printf

1. %c: 字符
2. %s: 字符创
3. %d: 数字
4. %T: 类型 - `x.(T)` - `x.(type)`
5. %x: 16 进制
6. %o: 8 进制
7. %b: 2 进制
8. %v: 变量值
9. %p: 变量地址

## 空间: `unsafe.Sizeof()`

1. 空间对齐: 伪共享
2. 不分配空间内存
   - `_` 匿名变量
   - `[空slice会分配空间]`空结构体: `var v struct{}`
   - channel 的零值: `var ch chan int`
