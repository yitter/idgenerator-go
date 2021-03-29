package contract

// 版权属于：yitter(yitter@126.com)
// 代码编辑：guoyahao
// 代码修订：yitter
// 开源地址：https://github.com/yitter/idgenerator-go

// OverCostActionArg .
type OverCostActionArg struct {
	ActionType             int32
	TimeTick               int64
	WorkerID               uint16
	OverCostCountInOneTerm int32
	GenCountInOneTerm      int32
	TermIndex              int32
}

// OverCostActionArg .
func (orca OverCostActionArg) OverCostActionArg(workerID uint16, timeTick int64, actionType int32, overCostCountInOneTerm int32, genCountWhenOverCost int32, index int32) {
	orca.ActionType = actionType
	orca.TimeTick = timeTick
	orca.WorkerID = workerID
	orca.OverCostCountInOneTerm = overCostCountInOneTerm
	orca.GenCountInOneTerm = genCountWhenOverCost
	orca.TermIndex = index
}
