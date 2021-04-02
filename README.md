#  ❄️ idenerator-go

## 介绍
项目更多介绍参照：https://github.com/yitter/idgenerator

## Go环境

1.SDK，go1.14

2.启用 Go-Modules

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

3. 安装方式
```
    go get -u -v github.com/yitter/idgenerator-go
```
或 go.mod 中添加引用
```
require github.com/yitter/idgenerator-go v1.2.0
```

## Go代码示例
```

// 定义参数
var options = idgen.NewIdGeneratorOptions(1)
options.WorkerId = 1
options.WorkerIdBitLength = 6
options.SeqBitLength = 6
// ...
idgen.SetIdGenerator(options)

// 调用方法生成Id
var id = idgen.NextId()

```

## 代码贡献者(按时间顺序)
guoyahao | amuluowin | houseme
