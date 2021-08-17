package services

import (
	"ddd-go-demo/domain/lottery/entity"
	"ddd-go-demo/domain/lottery/vo"
)

var LotterySvc *lotteryService

// lotteryService 抽奖服务
type lotteryService struct {
}

func (s *lotteryService) Lottery(userId uint64) (*entity.AwardInfo, error) {
	// 根据产品策略选出所有的奖品
	awards := []*entity.AwardInfo{
		{
			1,
			1,
			"30金币",
			30,
		}, {
			2,
			2,
			"华为P40",
			1,
		},
	}

	// 形成奖池对象, 根据产品给出的策略进行抽奖
	pool := vo.NewAwardPool(awards)
	return pool.Lottery(), nil
}

// GetDouDiAward 当奖品库存不充足时, 获取兜底奖品, 兜底奖品一般情况下是无限量的
func (s *lotteryService) GetDouDiAward(userId uint64) (*entity.AwardInfo, error) {
	award := &entity.AwardInfo{
		Id:   1,
		Type: 1,
		Name: "30金币",
		Cnt:  30,
	}

	return award, nil
}
