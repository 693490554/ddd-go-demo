package aggregate

import (
	"ddd-go-demo/domain/assets/entity"
)

var MChangeCoinAgg *ChangeCoinAgg

// todo 充血模型： 包含领域逻辑和持久化相关方法，对领域对象的CRUD还是通过repo来进行操作的。
// ChangeCoinAgg 用户金币变更聚合根
type ChangeCoinAgg struct {
	CoinRecord        *entity.CoinRecord
	CoinHistoryRecord *entity.CoinHistoryRecord
}

// AddCoin 增加用户金币数量, 增加和减少分为两个方法，语义明确
func (c *ChangeCoinAgg) AddCoin(cnt int64, remark string) {
	// todo 生成一条流水记录并且增加用户金币数量，此处体现聚合根的一个主要功能: 封装业务逻辑
	c.CoinHistoryRecord = entity.NewCoinHistoryRecord(c.CoinRecord.UserId, cnt, remark)
	c.CoinRecord.AddCoin(cnt)
}

// ReduceCoin 减少用户金币数量
func (c *ChangeCoinAgg) ReduceCoin(cnt int64, remark string) {
	c.CoinHistoryRecord = entity.NewCoinHistoryRecord(c.CoinRecord.UserId, cnt, remark)
	c.CoinRecord.ReduceCoin(cnt)
}

// Get 获取更改金币聚合根
func (c *ChangeCoinAgg) Get(userId uint64) (*ChangeCoinAgg, error) {
	return coinAggRepo.GetChangeCoinAgg(userId)
}

// Save 保存更改金币聚合根
func (c *ChangeCoinAgg) Save() error {
	return coinAggRepo.SaveChangeCoinAgg(c)
}

func NewChangeCoinAgg(coinRecord *entity.CoinRecord) *ChangeCoinAgg {
	return &ChangeCoinAgg{
		CoinRecord: coinRecord,
	}
}

var coinAggRepo *coinAggRepository

// coinAggRepository 金币仓储层
type coinAggRepository struct {
}

// GetChangeCoinAgg 获取修改金币聚合根
func (r *coinAggRepository) GetChangeCoinAgg(userId uint64) (*ChangeCoinAgg, error) {
	// 实际业务场景中，需从数据库中获取用户金币记录
	coinRecord := entity.NewCoinRecord(userId)
	// 数据库中可能不存在记录，此时需要生成一条新纪录
	return NewChangeCoinAgg(coinRecord), nil
}

// SaveChangeCoinAgg 保存金币流水聚合根
func (r *coinAggRepository) SaveChangeCoinAgg(agg *ChangeCoinAgg) error {
	// todo 聚合根中存在两个数据库对象，实际业务场景中，需开启事务保存
	// todo 此处体现了聚合根的一个主要功能：在聚合根内保证领域对象数据的一致性
	err := entity.MCoinRecord.SaveWithTx()
	if err != nil {
		return err
	}

	err = entity.MPrizeRecord.SaveWithTx()
	if err != nil {
		return err
	}

	// 事务提交，结束
	return nil
}
