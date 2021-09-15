// 领域服务层

package services

import (
	"ddd-go-demo/domain/assets/aggregate"
)

var CoinSvc *coinService

// coinRepository 金币仓储层
type coinService struct {
}

// AddUserCoin 增加用户金币数量
func (s *coinService) AddUserCoin(userId uint64, cnt int64, remark string) error {
	agg, err := aggregate.MChangeCoinAgg.Get(userId)
	if err != nil {
		return err
	}
	agg.AddCoin(cnt, remark)
	return agg.Save()
}
