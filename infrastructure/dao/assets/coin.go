package assets

import (
	"ddd-go-demo/domain/assets/aggregate"
	"ddd-go-demo/domain/assets/entity"
)

var CoinDao *coinDao

// coinDao 金币资产dao
type coinDao struct {
}

// GetCoinRecord 从数据库中获取用户金币记录
func (d *coinDao) GetCoinRecord(userId uint64) (*entity.CoinRecord, error) {
	// 此处不提供具体的实现
	return &entity.CoinRecord{
		UserId: userId,
	}, nil
}

// SaveCoinRecordWithTx 开启事务保存金币记录
func (d *coinDao) SaveCoinRecordWithTx(record *entity.CoinRecord) error {
	// 此处不提供具体的实现
	return nil
}

// SaveCoinHistoryWithTx 开启事务保存金币流水
func (d *coinDao) SaveCoinHistoryWithTx(record *entity.CoinHistoryRecord) error {
	// 此处不提供具体的实现
	return nil
}

// SaveChangeCoinAgg 保存更改金币聚合根
func (d *coinDao) SaveChangeCoinAgg(record *aggregate.ChangeCoinAgg) error {
	// 开启事务, 保存金币记录和流水记录, 失败的话需回滚
	err := d.SaveCoinRecordWithTx(record.CoinRecord)
	if err != nil {
		return err
	}

	err = d.SaveCoinHistoryWithTx(record.CoinHistoryRecord)
	if err != nil {
		return err
	}
	// 事务提交
	return nil
}
