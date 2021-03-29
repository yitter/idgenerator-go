// 版权属于：yitter(yitter@126.com)
// 代码编辑：guoyahao
// 代码修订：yitter
// 开源地址：https://gitee.com/yitter/idgenerator

package contract

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
