// 仓储层，负责聚合对象的CRUD操作

package repo

import (
	"ddd-go-demo/domain/assets/aggregate"
	"ddd-go-demo/domain/assets/entity"
	"ddd-go-demo/infrastructure/dao/assets"
)

var CoinRepo *coinRepository

// coinRepository 金币仓储层
type coinRepository struct {
}

// GetChangeCoinAgg 获取修改金币聚合根
func (r *coinRepository) GetChangeCoinAgg(userId uint64) (*aggregate.ChangeCoinAgg, error) {
	coinRecord, err := assets.CoinDao.GetCoinRecord(userId)
	if err != nil {
		return nil, err
	}

	// 数据库中可能不存在记录，此时需要生成一条新纪录
	if coinRecord == nil {
		coinRecord = entity.NewCoinRecord(userId)
	}
	return aggregate.NewChangeCoinAgg(coinRecord), nil
}

// SaveChangeCoinAgg 保存金币流水聚合根
func (r *coinRepository) SaveChangeCoinAgg(agg *aggregate.ChangeCoinAgg) error {
	// 聚合根中存在两个数据库对象，需开启事务保存
	// todo: 尽量不要在领域层中做开启事务的操作，开启事务的操作放在dao中，领域层更应关注业务逻辑，不应该关注如何开启事务并将数据进行持久化
	return assets.CoinDao.SaveChangeCoinAgg(agg)
}
