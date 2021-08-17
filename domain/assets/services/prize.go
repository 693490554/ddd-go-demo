package services

import (
	"ddd-go-demo/domain/assets/entity"
	"ddd-go-demo/infrastructure/dao/assets"
)

var PrizeSvc *prizeService

// coinRepository 金币仓储层
type prizeService struct {
}

// AddPrize 增加用户实物奖品
func (s *prizeService) AddPrize(userId uint64, prizeId uint32) error {
	record := entity.NewPrizeRecord(userId, prizeId)
	return assets.PrizeDao.AddPrize(record)
}
