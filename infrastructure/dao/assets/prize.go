package assets

import "ddd-go-demo/domain/assets/entity"

var PrizeDao *prizeDao

// prizeDao 实物资产dao
type prizeDao struct {
}

// GetPrizes 获取实物奖品记录
func (d *prizeDao) GetPrizes(userId uint64) ([]*entity.PrizeRecord, error) {
	return nil, nil
}

// AddPrize 增加实物资产
func (d *prizeDao) AddPrize(record *entity.PrizeRecord) error {
	return nil
}
