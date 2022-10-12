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

5. `{go 为了简介没有继承}`对象继承: `通过嵌套匿名结构体实现继承`

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

6. 接口 & 实现
