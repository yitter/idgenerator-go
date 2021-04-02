package gen

// 版权属于：yitter(yitter@126.com)
// 代码编辑：guoyahao
// 代码修订：yitter
// 开源地址：https://github.com/yitter/idgenerator-go

import (
	"sync"

	"github.com/yitter/idgenerator-go/contract"
)

var ins *YitIDHelper
var once sync.Once

// YitIDHelper .
type YitIDHelper struct {
	idGenInstance interface {
		NewLong() uint64
	}
}

// GetIns .
func GetIns() *YitIDHelper {
	once.Do(func() {
		ins = &YitIDHelper{}
	})
	return ins
}

// GetIDGenInstance .
func (yih *YitIDHelper) GetIDGenInstance() interface{} {
	return yih.idGenInstance
}

// SetIDGenerator .
func (yih *YitIDHelper) SetIDGenerator(options *contract.IDGeneratorOptions) {
	yih.idGenInstance = NewDefaultIDGenerator(options)
}

// NextID .
func (yih *YitIDHelper) NextID() uint64 {
	once.Do(func() {
		if yih.idGenInstance == nil {
			options := contract.NewIDGeneratorOptions(1)
			yih.idGenInstance = NewDefaultIDGenerator(options)
		}
	})

	return yih.idGenInstance.NewLong()
}

// IDHelperInit
func IDHelperInit() YitIDHelper {
	return YitIDHelper{}
}
