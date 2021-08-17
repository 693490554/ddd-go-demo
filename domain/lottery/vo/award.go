package vo

import "ddd-go-demo/domain/lottery/entity"

// AwardPool 奖池对象
type AwardPool struct {
	awardPool []*entity.AwardInfo
}

// Lottery 根据产品策略抽出某个奖品
func (p *AwardPool) Lottery() *entity.AwardInfo {
	return p.awardPool[0]
}

func NewAwardPool(pool []*entity.AwardInfo) *AwardPool {
	return &AwardPool{
		awardPool: pool,
	}
}
