package idgen

// @Project: idgenerator-go
// @Author: houseme
// @Description:
// @File: IdGeneratorException
// @Version: 1.0.0
// @Date: 2021/4/2 14:09
// GitHub: https://github.com/yitter/idgenerator-go

import "fmt"

// IdGeneratorException .
type IdGeneratorException struct {
	message string
	error   error
}

// IdGeneratorException .
func (e IdGeneratorException) IdGeneratorException(message ...interface{}) {
	fmt.Println(message...)
}

// Error .
func (e IdGeneratorException) Error(err error) string {
	e.message = err.Error()
	e.error = err
	return e.message
}
