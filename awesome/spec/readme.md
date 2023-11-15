[toc]

## go coding spec

1. naming var | func | struct etcs
2. naming module
3. naming package & file
4. naming project struct(open source best practice)
5. release

## naming module

1. 中划线分隔(redis-go)
2. 命名尽量在三个单词以内
3. 命名是项目功能的描述 | 代号(础组件或者开源项目)

## naming package&file: `目录不一样,包名不会冲突`

1. 简短、清晰且富有描述性: 全小写|全字母|无保留字

   - bad: computeServiceClient, priority_queue
   - good: time, list, http
   - good(可理解性的缩写): strconv, syscall, fmt
   - good(全小写): stringset, helloworld
   - good(通用缩写): `certificate->crt | strings->strx`

2. 尽量不使用无意义的包名

   - bad: util, common, misc
   - good: util.SortStringSet() -> stringset.Sort()

3. 主动分包(目录清晰): 功能和领域

   - bad: util.NewStringSet() & util.SortStringSet()
   - good: stringset.New() & stringset.Sort()

4. 尽量不用通用词(api, interface, types, models) & 不要刻意避免包名冲突

   - good: runtime/pprof | net/http/pprof
   - bad: api/user.go | api/post.go

5. 文件命名

   - 小写字母和下划线: `http_server.go`
   - 描述性名称
   - 测试文件: `http_server_test.go`
   - 平台特定文件: `http_server_windows.go`

6. **函数**要避免与包名重复(与标准库冲突的情况可以在包名后增加 x 表示扩展)

   - good: http.Server | ring.New() | `net -> netx`
   - bad: http.HttpServer | ring.NewRing()

7. 函数名字简化

   - package 内含一个类型: list.New(), time.Parse()
   - package 内含多个类型: time.NewTicker(), context.WithTimeout

## naming var | func | struct etcs

1. 驼峰法命名: 大写可见 + 单词缩写默认全大写(全小写) + 保持简短

   - bad: userId || baiduCdn || index/reader/buffer
   - good: userID || baiduCDN || i/r/b

2. 结构体|接口: 驼峰法命名 + 大写可见

   - 名词或名词短语
   - 命名不能与包名称相同
   - 接口只有一个方法, 默认为方法名+er 来命名接口: `type Reader interface {}`
   - 自定义 error: `名称Error`

3. 方法|函数: 驼峰法命名 + 大写可见

   - 动词或动词短语
   - 注意换行: 80 字符上限
   - bad: p.GetFirstName()
   - good: p.FirstName()

## short

1. 常见的缩写

   ```go
   src = source
   srv = server
   arg = argument
   conn = connect, connection
   attr = attribute
   abs = absolute
   min = minimum
   len = length
   auth = authenticate
   buf = buffer
   ctl = control
   ctx = context
   str = string
   msg = message
   fmt = format
   dest = destination
   diff = difference
   orig = original
   recv = receive
   ref = reference
   repo = repository
   util = utility
   fmt = format
   ```

## naming project struct(open source best practice)

1. canonical

   ```go
   /myapp
   ├── cmd                                      // 主要入口代码(可多个)
   │   ├── server
   │   │   ├── main.go
   │   │   ├── server_test.go
   │   └── client
   │       ├── main.go
   │       └── client_test.go
   ├── pkg                                      // 可以被其他应用程序重用的代码
   │   ├── http
   │   │   ├── server.go
   │   │   ├── server_test.go
   │   │   ├── client.go
   │   │   └── client_test.go
   │   ├── database
   │   │   ├── connection.go
   │   │   └── connection_test.go
   │   └── util
   │       ├── converter.go
   │       └── converter_test.go
   └── internal                               // 只能被当前应用程序使用的代码
       ├── config
       │   ├── config.go
       │   └── config_test.go
       └── model
           ├── user.go
           └── user_test.go
   ```

## release

1. directory sample

   ```go
   .
   ├── LICENSE
   ├── README.md
   ├── go.mod
   └── stringutil
       ├── reverse.go
       └── reverse_test.go
   ```

2. 注意点

   - 写详细的 README
   - 具有准确有效的 COMMIT
   - 符合相关代码规范(上面的所有)
   - 代码同级目录下提供相关测试用例
   - 不要对已发版做任何修改: 应该迭代代码发布新版本

3. release(push tag 即可)

   - go mod tidy: 去除不需要的依赖
   - **git tag|push 规范标签**

     ![avatar](/static/image/release-version-number.png)

     | version stage       | example       | message to developers                  |
     | :------------------ | :------------ | :------------------------------------- |
     | in development      | v0.x.x        | unstable and no backward compatibility |
     | major version       | v1.x.x        | backward compatibility                 |
     | minor version       | vx.4.x        | backward compatibility                 |
     | patch version       | vx.x.1        | fix bug and backward compatibility     |
     | pre-release version | vx.x.x-beta.2 | unstable and pre-release milestone     |

4. version flow

   - pseudo-version number: `v0.0.0-20170915032832-14c0d48ead0c`
   - v0 number: `v0.x.x`
   - pseudo-version number: `vx.0.0`
   - pre-release version: `vx.x.x-beta.2`
   - minor version: `vx.4.x`
   - patch version: `vx.x.1`
   - major version: `v1.x.x`

---

## reference

1. [go-module version](https://go.dev/doc/modules/version-numbers)
2. [go-module](https://zhuanlan.zhihu.com/p/599710762)
3. [项目组织](https://zhuanlan.zhihu.com/p/124198314)
4. [项目目录](https://zhuanlan.zhihu.com/p/659823790)
5. [项目目录结构](https://blog.csdn.net/weixin_44798288/article/details/125500769)
6. [开源结构-最佳实践](<[开源界优秀项目的结构](https://www.cnblogs.com/Paul-watermelon/p/11230197.html)>)
7. [编程-命名规范](https://zhuanlan.zhihu.com/p/654606942)
8. [编程-命名规范](https://blog.csdn.net/piaohai/article/details/121389429)
9. [编程-命名规范](https://blog.csdn.net/q1009020096/article/details/108593135)
10. [package-命名规范](https://www.zhaohuabing.com/learning-golang/docs/package/naming/)
11. [package-命名规范](https://cloud.tencent.com/developer/article/2311609)
12. [release](https://mp.weixin.qq.com/s/UanvwYKB3fosRYMP9dW6Wg)
