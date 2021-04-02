package idgen

import (
	"sync"
)

//var ins *YitIdHelper
//var once sync.Once

var singletonMutex sync.Mutex
var idGenerator *DefaultIdGenerator

//// YitIDHelper .
//type YitIDHelper struct {
//	idGenInstance interface {
//		NewLong() uint64
//	}
//}
//
//// GetIns .
//func GetIns() *YitIDHelper {
//	once.Do(func() {
//		ins = &YitIDHelper{}
//	})
//	return ins
//}
//
//// GetIDGenInstance .
//func (yih *YitIDHelper) GetIDGenInstance() interface{} {
//	return yih.idGenInstance
//}
//
//// SetIDGenerator .
//func (yih *YitIDHelper) SetIDGenerator(options *IdGeneratorOptions) {
//	yih.idGenInstance = NewDefaultIdGenerator(options)
//}
//
//// NextId .
//func (yih *YitIDHelper) NextId() uint64 {
//	once.Do(func() {
//		if yih.idGenInstance == nil {
//			options := NewIdGeneratorOptions(1)
//			yih.idGenInstance = NewDefaultIdGenerator(options)
//		}
//	})
//
//	return yih.idGenInstance.NewLong()
//}
//
//// IDHelperInit .
//func IDHelperInit() YitIDHelper {
//	return YitIDHelper{}
//}

// SetIdGenerator .
func SetIdGenerator(options *IdGeneratorOptions) {
	singletonMutex.Lock()
	idGenerator = NewDefaultIdGenerator(options)
	singletonMutex.Unlock()
}

// NextId .
func NextId() uint64 {
	if idGenerator == nil {
		singletonMutex.Lock()
		defer singletonMutex.Unlock()
		options := NewIdGeneratorOptions(1)
		idGenerator = NewDefaultIdGenerator(options)
	}

	return idGenerator.NewLong()
}
