## Test(TDD)[https://www.liwenzhou.com/posts/Go/unit-test/]

1. 分类

   - 测试函数: 前缀为 Test
   - 基准函数: 前缀为 Benchmark
   - 示例函数: 前缀为 Example(示例文档)

2. 演进: `TestSplit`

   - 单个测试用例: 实际结果与预期比较
   - 多个测试用例: 测试组
   - 多个测试用例区分: map 命令的自测试
   - 基准测试: 一定负载下的测试, `内存占用 | 内存申请`
   - 性能测试: 同一函数不同数据的测试 || 不同函数同一数据的测试
   - 示例函数: 提供文档和使用示例

3. 相关命令

   - go test -v
   - go test -run
   - go test -cover
   - go test -cover -coverprofile=c.out
