package main

import (
	"C"
	"fmt"
	"time"

	"github.com/yitter/idgenerator-go/idgen"
)

func SetOptions(workerId uint16) {
	var options = idgen.NewIdGeneratorOptions(workerId)
	idgen.SetIdGenerator(options)
}

func NextId() int64 {
	return idgen.NextId()
}

func main() {
	const testGenId = true // 测试设置

	if testGenId {
		// 自定义参数
		var options = idgen.NewIdGeneratorOptions(1)
		options.WorkerIdBitLength = 6
		options.SeqBitLength = 10
		options.BaseTime = time.Date(2020, 2, 20, 2, 20, 2, 20, time.UTC).UnixNano() / 1e6
		idgen.SetIdGenerator(options)

		var genCount = 500000
		for j := 0; j < 10; j++ {
			var begin = time.Now().UnixNano() / 1e6
			for i := 0; i < genCount; i++ {
				// 生成ID
				idgen.NextId()
				// fmt.Println(id)
			}
			var end = time.Now().UnixNano() / 1e6

			fmt.Println("耗时：", (end - begin), "ms")
			time.Sleep(time.Duration(1000) * time.Millisecond)
		}
	} else {
		// ip := "localhost"
		ipAddr := C.CString("localhost:6379")
		passChar := C.CString("")
		sentinelMasterName := C.CString("")

		workerIdList := RegisterMany(ipAddr, passChar, 4, sentinelMasterName, 3, 10, 5, 15)
		for _, value := range workerIdList {
			fmt.Println("注册的WorkerId:", value)
		}

		id := RegisterOne(ipAddr, passChar, 4, sentinelMasterName, 3, 10, 15)
		fmt.Println("注册的WorkerId:", id)

		// C.free(unsafe.Pointer(ipChar))
		// C.free(unsafe.Pointer(passChar))

		// var workerId = regworkerid.RegisterOne(ip, 6379, "", 4)
		// fmt.Println("注册的WorkerId:", workerId)

		fmt.Println("end")
		time.Sleep(time.Duration(300) * time.Second)

	}
}

// To Build a dll/so：

// windows:
// go build -o ./target/yitidgengo.dll -buildmode=c-shared main.go reg.go

// linux init: go install -buildmode=shared -linkshared std
// go build -o ./target/yitidgengo.so -buildmode=c-shared main.go reg.go

// https://books.studygolang.com/advanced-go-programming-book/ch2-cgo/ch2-09-static-shared-lib.html
