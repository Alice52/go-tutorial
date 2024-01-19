## introduce

1. syntax

   - key word
   - var
   - basic data-type

     - string:
     - int:
     - float32/64
     - complex64/128
     - rune: **rune is an alias for int32 and is equivalent to int32 in all ways, used to handle Unicode and UTF-8 char**
     - byte: **byte is an alias for int8 and is equivalent to int8 in all ways, used to handle ASCII characters**
     - bool

   - operator
   - process control
     - if
     - for
     - switch
     - goto
   - Array
   - Slice
   - Map
   - Struct
   - Function
   - Interface
   - Pointer
   - Channel

2. feature

   - GC
   - `Simple thinking`: no extends, polymorphism, class
   - Elegant `concurrency` using GoRoutines and Channels
   - Multi-core support
   - Very good `raw CPU-and-memory-bound performance`
   - Types aren’t cumbersome
   - Built in testing framework
   - Static code analysis
   - Compilation / deployment
   - `Cross compile`
   - 扩展性
     1. reflect
     2. generic
     3. ast

3. 优点

   - 自动立即回收
   - 更丰富的内置类型
   - `函数多返回值`
   - 错误处理
   - 匿名函数和闭包
   - 类型和接口
   - 并发编程
   - 反射
   - 语言交互性

4. `& vs *`

   - `& 是取地址符号`: 取值的地址`{&变量值 就是 指针}`
   - `* 是指针运算符`: 取指针地址中存的值`{*指针 就是变量值}`
     1. 表示一个变量是指针类型
     2. 表示一个指针变量所指向的存储单元{地址所存储的值}
   - 参数等传递都是按值传递

     1. 传实例: 真正传递的是实例副本的地址 + 不可以修改实例的字段
     2. 传指针: 指针的副本的地址(指向该实例) + 只可以修改实例的字段
     3. 传指针的指针: 指针的副本的地址(指向前一个指针) + 可以对实例重新赋值(`*p=nil`)

     ```go
     emial := config.Email{}  // 实例
     ept := &emial            // 指针
     eptpt := &ept            // 指针的指针
     ```

5. 执行顺序

   - var 的全局变量
   - init() 函数
   - main() 函数
