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


func SetIdGenerator(options *IdGeneratorOptions) {
	singletonMutex.Lock()
	idGenerator = NewDefaultIdGenerator(options)
	singletonMutex.Unlock()
}

func NextId() uint64 {
	if idGenerator == nil {
		singletonMutex.Lock()
		if idGenerator == nil {
			options := NewIdGeneratorOptions(1)
			idGenerator = NewDefaultIdGenerator(options)
		}
		singletonMutex.Unlock()
	}

	return idGenerator.NewLong()
}
