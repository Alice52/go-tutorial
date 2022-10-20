## Test(TDD)[https://www.liwenzhou.com/posts/Go/unit-test/]

1. 分类: `*_test.go`

   - 测试函数: 前缀为 Test
   - 基准函数: 前缀为 Benchmark
   - 示例函数: 前缀为 Example(示例文档)

2. 演进: `TestSplit`

   - 单个测试用例: 实际结果与预期比较
   - 多个测试用例: 测试组
   - 多个测试用例区分: 子测试, map 命令的自测试
   - 基准测试: 一定负载下的测试, `内存占用 | 内存申请`
   - 性能测试: 同一函数不同数据的测试 || 不同函数同一数据的测试
   - 示例函数: 提供文档和使用示例

3. 相关命令

   - go test ./... -v: `./...`(该目录及子目录) + `输出完整的测试结果`
   - go test -run REG: `运行函数名匹配的`
   - `go test -bench=REG -benchmem`: 运行基准测试
   - go test -short: 跳过测试
   - go test -cover
   - go test -cover -coverprofile=c.out: html 打开覆盖说明
   - `gotests -all -w split.go`: 生成测试代码(只需要填充测试用例)

4. 打桩(stub)

   - 屏蔽: 不想在单元测试用引入数据库连接等重资源
   - 补齐: 依赖的上下游函数或方法还未实现

5. 其他

   - 回归测试
   - 跳过指定测试
   - 子测试: ` t.Run("c-name", func(t *testing.T){...})`
   - 测试组
   - 测试覆盖率
   - **表格驱动测试**: 是编写更清晰测试的一种方式和视角
   - 生成测试用例: `go get xxx/gotests` + `gotests -all -w split.go`

## Best Pratice

1. 单元测试

   - goconvey: webui &

2. 基准测试

   - 想办法是否能调用单元测试

3. 示例函数
4. 工具

   - net: httptest{server} + **gock{client}**
   - redis: miniredis
   - mysql: `go-sqlmock{不真实存储数据}`
   - 打桩: monkey

5. 相关原则

   - 剔除干扰因素
   - 接口抽象进行解耦
   - 依赖注入代替隐式依赖: 减少使用全局变量
   - dpp: solid

---

## issues

1. 如何在基准测试中调用单元测试
