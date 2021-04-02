package main

import (
	"fmt"
	"time"

	"github.com/yitter/idgenerator-go/contract"
	"github.com/yitter/idgenerator-go/gen"
)

func main() {
	// 方法一：直接采用默认方法生成一个Id
	var yid = gen.IDHelperInit()
	fmt.Println(yid.NextID())

	// 方法二：自定义参数
	var options = contract.NewIDGeneratorOptions(1)
	//options.WorkerIDBitLength = 6
	//options.SeqBitLength = 6
	//options.TopOverCostCount = 2000
	//options.BaseTime = time.Date(2020, 2, 20, 2, 20, 2, 20, time.UTC).UnixNano() / 1e6
	yid.SetIDGenerator(options)

	var times = 50000
	for {
		var begin = time.Now().UnixNano() / 1e6
		for i := 0; i < times; i++ {
			yid.NextID()
		}
		var end = time.Now().UnixNano() / 1e6
		fmt.Println(end - begin)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}

}
