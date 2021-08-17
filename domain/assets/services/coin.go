// 领域服务层

package services

import "ddd-go-demo/domain/assets/repo"

var CoinSvc *coinService

// coinRepository 金币仓储层
type coinService struct {
}

// AddUserCoin 增加用户金币数量
func (s *coinService) AddUserCoin(userId uint64, cnt int64, remark string) error {
	agg, err := repo.CoinRepo.GetChangeCoinAgg(userId)
	if err != nil {
		return err
	}

	agg.AddCoin(cnt, remark)
	return repo.CoinRepo.SaveChangeCoinAgg(agg)
}
