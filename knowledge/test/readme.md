## Test(TDD)[https://www.liwenzhou.com/posts/Go/unit-test/]

1. 分类

   - 测试函数: 前缀为 Test
   - 基准函数: 前缀为 Benchmark
   - 示例函数: 前缀为 Example(示例文档)

### 测试函数

1. 测试用例: `*testing.T`
2. 测试组: []struct{input/sep/expect} + for + 函数执行
   - 多个测试用例
3. 子测试: map[name]struct
   - 测试用例比较多的时无法直接识别哪个测试用例失败
   - 单独运行某个测试用例
4. 测试覆盖率

   ```go
   go test -cover

   go test -cover -coverprofile=c.out
   ```

5. [sample](/tutorials/cn.edu.ntu.awesome/test/split_test.go#L8)

### 基准测试

1. 在一定负载下测试性能: `*testing.B`
2. b.N 的值是系统根据实际情况去调整的: 保证测试的稳定性
3. command

   ```go
   func BenchmarkSplit(b *testing.B) {
      for i := 0; i < b.N; i++ {
         Split("沙河有沙又有河", "沙")
      }
   }

   go test -bench=Split -benchmem
   // 2-core                   执行次数               平均每次执行时间         每次执行消耗空间      每次执行分配内存次数
   // BenchmarkSplit-2         5017052               236.9 ns/op           112 B/op          3 allocs/op
   ```

4. 性能比较函数:

   - 比较同一个函数处理 1000 个元素的耗时与处理 1 万甚至 100 万个元素的耗时的差别
   - 对于同一个任务究竟使用哪种算法性能最佳

     ```go
     func benchmarkFib(b *testing.B, n int) {
        for i := 0; i < b.N; i++ {
           Fib(n)
        }
     }

     func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
     func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }

     // go test -bench=. -benchtime=20s // 增加最小基准时间
     ```

5. 重置时间: `b.ResetTimer() // 重置计时器`
6. 并行测试
   ```go
   func BenchmarkSplitParallel(b *testing.B) {
      // b.SetParallelism(1) // 设置使用的CPU数
      b.RunParallel(func(pb *testing.PB) {
         for pb.Next() {
            Split("沙河有沙又有河", "沙")
         }
      })
   }
   // go test -bench=. -cpu 1
   ```

### Setup 与 TearDown

1. TestMain 函数来可以在测试之前进行额外的设置(setup)或在测试之后进行拆卸(teardown)操作
2. `*testing.M`
3. sample

   ```go
   // 测试集的Setup与Teardown
   func setupTestCase(t *testing.T) func(t *testing.T) {
      t.Log("如有需要在此执行:测试之前的setup")
      return func(t *testing.T) {
         t.Log("如有需要在此执行:测试之后的teardown")
      }
   }

   defer setupTestCase(t)(t)
   ```

### 示例函数
