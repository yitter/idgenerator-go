package contract

// @Project: idgenerator-go
// @Author: houseme
// @Description:
// @File: IDGeneratorException
// @Version: 1.0.0
// @Date: 2021/4/2 14:09
// @Package contract
// GitHub: https://github.com/yitter/idgenerator-go

import "fmt"

// IDGeneratorException .
type IDGeneratorException struct {
	message string
	error   error
}

// IDGeneratorException .
func (e IDGeneratorException) IDGeneratorException(message ...interface{}) {
	fmt.Println(message...)
}

// Error .
func (e IDGeneratorException) Error(err error) string {
	e.message = err.Error()
	e.error = err
	return e.message
}
