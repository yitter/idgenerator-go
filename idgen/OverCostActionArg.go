package idgen

// OverCostActionArg .
type OverCostActionArg struct {
	ActionType             int32
	TimeTick               int64
	WorkerId               uint16
	OverCostCountInOneTerm int32
	GenCountInOneTerm      int32
	TermIndex              int32
}

// OverCostActionArg .
func (orca OverCostActionArg) OverCostActionArg(workerId uint16, timeTick int64, actionType int32, overCostCountInOneTerm int32, genCountWhenOverCost int32, index int32) {
	orca.ActionType = actionType
	orca.TimeTick = timeTick
	orca.WorkerId = workerId
	orca.OverCostCountInOneTerm = overCostCountInOneTerm
	orca.GenCountInOneTerm = genCountWhenOverCost
	orca.TermIndex = index
}
