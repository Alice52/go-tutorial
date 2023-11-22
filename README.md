## go-tutorial

1. This repository records golang learning process
2. pros: 简洁 & 高效

## blemish

1. without syntax annotation
2. without aop/validate
3. without named function and default args
4. without agent enhance

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
