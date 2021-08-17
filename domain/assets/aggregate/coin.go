package aggregate

import (
	"ddd-go-demo/domain/assets/entity"
)

// ChangeCoinAgg 用户金币变更聚合根
type ChangeCoinAgg struct {
	CoinRecord        *entity.CoinRecord
	CoinHistoryRecord *entity.CoinHistoryRecord
}

// AddCoin 增加用户金币数量, 增加和减少分为两个方法，语义明确
func (c *ChangeCoinAgg) AddCoin(cnt int64, remark string) {
	c.CoinHistoryRecord = entity.NewCoinHistoryRecord(c.CoinRecord.UserId, cnt, remark)
	c.CoinRecord.AddCoin(cnt)
}

// ReduceCoin 减少用户金币数量
func (c *ChangeCoinAgg) ReduceCoin(cnt int64, remark string) {
	c.CoinHistoryRecord = entity.NewCoinHistoryRecord(c.CoinRecord.UserId, cnt, remark)
	c.CoinRecord.ReduceCoin(cnt)
}

func NewChangeCoinAgg(coinRecord *entity.CoinRecord) *ChangeCoinAgg {
	return &ChangeCoinAgg{
		CoinRecord: coinRecord,
	}
}
