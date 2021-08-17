package entity

// AwardInfo 抽奖奖品信息
type AwardInfo struct {
	Id   uint32 // 奖品唯一id
	Type uint32 // 奖品类型，金币，实物奖品
	Name string // 奖品名称
	Cnt  uint32 // 奖励数量
}
