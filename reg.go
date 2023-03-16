package main

import (
	"C"
)
import "github.com/yitter/idgenerator-go/regworkerid"

// RegisterOne 注册一个 WorkerId，会先注销所有本机已注册的记录
// address: Redis连接地址，单机模式示例：127.0.0.1:6379，哨兵/集群模式示例：127.0.0.1:26380,127.0.0.1:26381,127.0.0.1:26382
// password: Redis连接密码
// db: Redis指定存储库，示例：1
// sentinelMasterName: Redis 哨兵模式下的服务名称，示例：mymaster，非哨兵模式传入空字符串即可
// maxWorkerId: WorkerId 最大值，示例：63
// minWorkerId: WorkerId 最小值，示例：30
// lifeTimeSeconds: WorkerId缓存时长（秒，3的倍数）
//export RegisterOne
func RegisterOne(address *C.char, password *C.char, db int, sentinelMasterName *C.char, minWorkerId int32, maxWorkerId int32, lifeTimeSeconds int32) int32 {
	return regworkerid.RegisterOne(regworkerid.RegisterConf{
		Address:         C.GoString(address),
		Password:        C.GoString(password),
		DB:              db,
		MasterName:      C.GoString(sentinelMasterName),
		MinWorkerId:     minWorkerId,
		MaxWorkerId:     maxWorkerId,
		LifeTimeSeconds: lifeTimeSeconds,
	})
}

// RegisterMany 注册多个 WorkerId，会先注销所有本机已注册的记录
// address: Redis连接地址，单机模式示例：127.0.0.1:6379，哨兵/集群模式示例：127.0.0.1:26380,127.0.0.1:26381,127.0.0.1:26382
// password: Redis连接密码
// db: Redis指定存储库，示例：1
// sentinelMasterName: Redis 哨兵模式下的服务名称，示例：mymaster，非哨兵模式传入空字符串即可
// maxWorkerId: WorkerId 最大值，示例：63
// minWorkerId: WorkerId 最小值，示例：30
// totalCount: 获取N个WorkerId，示例：5
// lifeTimeSeconds: WorkerId缓存时长（秒，3的倍数）
//export RegisterMany
func RegisterMany(address *C.char, password *C.char, db int, sentinelMasterName *C.char, minWorkerId int32, maxWorkerId int32, totalCount int32, lifeTimeSeconds int32) []int32 {
	return regworkerid.RegisterMany(regworkerid.RegisterConf{
		Address:         C.GoString(address),
		Password:        C.GoString(password),
		DB:              db,
		MasterName:      C.GoString(sentinelMasterName),
		MinWorkerId:     minWorkerId,
		MaxWorkerId:     maxWorkerId,
		TotalCount:      totalCount,
		LifeTimeSeconds: lifeTimeSeconds,
	})
}

// UnRegister 注销本机已注册的 WorkerId
//export UnRegister
func UnRegister() {
	regworkerid.UnRegister()
}

// Validate 检查本地WorkerId是否有效（0-有效，其它-无效）
//export Validate
func Validate(workerId int32) int32 {
	return regworkerid.Validate(workerId)
}
