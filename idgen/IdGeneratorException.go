/*
 * 版权属于：yitter(yitter@126.com)
 * 代码编辑：guoyahao
 * 代码修订：yitter
 * 开源地址：https://github.com/yitter/idgenerator
 */

// Package idgen is a snowflake id generator
package idgen

import "fmt"

type IdGeneratorException struct {
	message string
	error   error
}

func (e IdGeneratorException) IdGeneratorException(message ...interface{}) {
	fmt.Println(message...)
}
