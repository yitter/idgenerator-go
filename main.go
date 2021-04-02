package main

import (
	"fmt"
	"time"
	"github.com/yitter/idgenerator-go/idgen"
)

func main() {
	//fmt.Println("start time:", time.Now())

	// 方法一：直接采用默认方法生成一个Id
	fmt.Println(idgen.NextId())

	// 方法二：自定义参数
	var options = idgen.NewIdGeneratorOptions(1)
	//options.WorkerIdBitLength = 6
	//options.SeqBitLength = 6
	//options.BaseTime = time.Date(2020, 2, 20, 2, 20, 2, 20, time.UTC).UnixNano() / 1e6
	idgen.SetIdGenerator(options)

	var times = 50000
	for {
		var begin = time.Now().UnixNano() / 1e3
		for i := 0; i < times; i++ {
			// fmt.Println(idgen.NextId())
			idgen.NextId()
		}

		var end = time.Now().UnixNano() / 1e3
		fmt.Println(end - begin)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}

	//fmt.Println("end time:", time.Now())
}
