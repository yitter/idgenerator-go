package contract

// 版权属于：yitter(yitter@126.com)
// 代码编辑：guoyahao
// 代码修订：yitter
// 开源地址：https://github.com/yitter/idgenerator-go

import "fmt"

// IDGeneratorException .
type IDGeneratorException struct {
	message string
	error   error
}

// IDGeneratorException .
func (e IDGeneratorException) IDGeneratorException(message ...interface{}) {
	fmt.Println(message)
}
