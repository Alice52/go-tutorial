## go-tutorial

1. This repository records golang learning process

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
