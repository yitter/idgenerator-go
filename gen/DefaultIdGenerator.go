// 版权属于：yitter(yitter@126.com)
// 代码编辑：guoyahao
// 代码修订：yitter
// 开源地址：https://github.com/yitter/idgenerator-go

package gen

import (
	"time"

	"github.com/yitter/idgenerator-go/contract"
	"github.com/yitter/idgenerator-go/core"
)

// DefaultIDGenerator .
type DefaultIDGenerator struct {
	Options              *contract.IDGeneratorOptions
	SnowWorker           contract.ISnowWorker
	IDGeneratorException contract.IDGeneratorException
}

// NewDefaultIDGenerator .
func NewDefaultIDGenerator(options *contract.IDGeneratorOptions) *DefaultIDGenerator {
	if options == nil {
		panic("dig.Options error.")
	}

	minTime := int64(631123200000) // time.Now().AddDate(-30, 0, 0).UnixNano() / 1e6
	if options.BaseTime < minTime || options.BaseTime > time.Now().UnixNano()/1e6 {
		panic("BaseTime error.")
	}

	if options.SeqBitLength+options.WorkerIDBitLength > 22 {
		panic("error：WorkerIDBitLength + SeqBitLength <= 22")
	}

	maxWorkerIDNumber := uint16(1<<options.WorkerIDBitLength) - 1
	if options.WorkerID > maxWorkerIDNumber {
		panic("WorkerID error. (range:[1, " + string(maxWorkerIDNumber) + "]")
	}

	if options.SeqBitLength < 2 || options.SeqBitLength > 21 {
		panic("SeqBitLength error. (range:[2, 21])")
	}

	maxSeqNumber := uint32(1<<options.SeqBitLength) - 1
	if options.MaxSeqNumber > maxSeqNumber {
		panic("MaxSeqNumber error. (range:[1, " + string(maxSeqNumber) + "]")
	}

	if options.MinSeqNumber > maxSeqNumber {
		panic("MinSeqNumber error. (range:[1, " + string(maxSeqNumber) + "]")
	}

	var snowWorker contract.ISnowWorker

	switch options.Method {
	case 1:
		snowWorker = core.NewSnowWorkerM1(options)
	case 2:
		snowWorker = core.NewSnowWorkerM2(options)
	default:
		snowWorker = core.NewSnowWorkerM1(options)
	}

	if options.Method == 1 {
		time.Sleep(time.Duration(500) * time.Microsecond)
	}
	return &DefaultIDGenerator{
		Options:    options,
		SnowWorker: snowWorker,
	}
}

// NewLong .
func (dig DefaultIDGenerator) NewLong() uint64 {
	return dig.SnowWorker.NextID()
}
