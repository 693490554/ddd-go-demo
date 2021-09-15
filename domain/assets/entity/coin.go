package entity

import (
	"time"
)

var MCoinRecord *CoinRecord

// CoinRecord 用户金币记录
// todo 充血模型
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

func (c *CoinRecord) Get(userId uint64) (*CoinRecord, error) {
	return coinRepo.Get(userId)
}

func (c *CoinRecord) Save() error {
	return coinRepo.Save(c)
}

func (c *CoinRecord) SaveWithTx() error {
	return coinRepo.SaveWithTx(c)
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

var coinRepo *coinRepository

// coinRepository 金币仓储层
type coinRepository struct {
}

// Get 获取金币对象
func (c *coinRepository) Get(userId uint64) (*CoinRecord, error) {
	// 实际业务中，此处应该是从db中获取记录
	return NewCoinRecord(userId), nil
}

// Save 保存金币对象
func (c *coinRepository) Save(*CoinRecord) error {
	// 实际业务中，此处应该是将记录保存至db
	return nil
}

// SaveWithTx 开启事务保存
func (c *coinRepository) SaveWithTx(*CoinRecord) error {
	// 实际业务中，此处应该是将记录保存至db
	return nil
}
