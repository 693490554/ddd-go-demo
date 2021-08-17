package entity

import "time"

// CoinRecord 用户金币记录
type CoinRecord struct {
	Id         uint64
	UserId     uint64
	Cnt        uint64 // 用户金币数量>=0
	CreateTime int64
	UpdateTime int64
}

func (c *CoinRecord) AddCoin(cnt int64) {
	c.UpdateTime = time.Now().Unix()
	c.Cnt += uint64(cnt)
}

func (c *CoinRecord) ReduceCoin(cnt int64) {
	c.UpdateTime = time.Now().Unix()
	c.Cnt -= uint64(cnt)
}

func NewCoinRecord(userId uint64) *CoinRecord {
	return &CoinRecord{
		UserId:     userId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
}

// CoinHistoryRecord 金币流水记录, 记录金币获取及消耗途径
type CoinHistoryRecord struct {
	Id         uint64
	UserId     uint64
	Cnt        int64
	Remark     string
	CreateTime int64
}

func NewCoinHistoryRecord(userId uint64, cnt int64, remark string) *CoinHistoryRecord {
	return &CoinHistoryRecord{
		UserId:     userId,
		Cnt:        cnt,
		Remark:     remark,
		CreateTime: time.Now().Unix(),
	}
}
