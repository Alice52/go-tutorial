## usage

1. 如果使用 github.com/xxx@xx`(look like URLs)` 在 go work 的 module_a 下执行 go mod tidy 会报错(尝试在公网上查找 module_b)
2. 只有非类似链接的本地依赖描述才会使用本地依赖 || 或者使用 replace 进行本地替换

## reference

1. [sample](https://github.com/bozaro/go-work-play/tree/go-mod-tidy)
2. [issue](https://github.com/golang/go/issues/50750)
