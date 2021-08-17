package services

var LotteryRiskSvc *lotteryRiskService

// lotteryRiskService 抽奖风控服务
type lotteryRiskService struct {
}

// UserJoinLotteryIsOk 用户是否可以参与抽奖
func (s *lotteryRiskService) UserJoinLotteryIsOk(userId uint64) error {
	return nil
}
