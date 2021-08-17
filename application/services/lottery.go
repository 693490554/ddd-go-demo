// todo: 应用层服务，负责领域服务的编排，实现完整的业务流程，不应该包含业务逻辑

package services

import (
	assets "ddd-go-demo/domain/assets/services"
	inventory "ddd-go-demo/domain/inventory/services"
	lottery "ddd-go-demo/domain/lottery/services"
	risk "ddd-go-demo/domain/risk/services"
	"errors"
)

var LotterySvc *lotteryService

// lotteryService 应用层抽奖service
type lotteryService struct {
}

// UserLottery 用户抽奖，在CQE中属于command, CQE在项目中的README中有详细介绍
func (s *lotteryService) UserLottery(userId uint64) (string, error) {
	// 过风控
	err := risk.LotteryRiskSvc.UserJoinLotteryIsOk(userId)
	if err != nil {
		return "", err
	}

	// 抽奖-根据产品的策略抽出奖品
	awardInfo, err := lottery.LotterySvc.Lottery(userId)
	if err != nil {
		return "", err
	}

	// 判断奖品库存是否充足
	enough, err := inventory.PrizeInventorySvc.InventoryIsEnough(awardInfo.Id, awardInfo.Cnt)
	if err != nil {
		return "", err
	}
	// 库存不充足获取兜底奖励
	if !enough {
		awardInfo, err = lottery.LotterySvc.GetDouDiAward(userId)
		if err != nil {
			return "", err
		}
	}

	// 调用资产领域发放奖品
	// todo: 奖品类型存在于抽奖上下文中，资产上下文不应关注奖品类型，两个不同上下文交互有两种方式实现
	// 1. 外部判断不同type调用不同的发奖方法
	// 2. 资产上下文中实现防腐层，在防腐层中增加对type的判断，不同的type调用不同发资产方法
	if awardInfo.Type == 1 {
		err = assets.CoinSvc.AddUserCoin(userId, int64(awardInfo.Cnt), "抽奖发放")
		if err != nil {
			return "", err
		}
	} else if awardInfo.Type == 2 {
		err = assets.PrizeSvc.AddPrize(userId, awardInfo.Id)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("奖品类型不支持")
	}

	// 实际业务场景中，发奖成功后lottery领域可能需要产生领域事件(domain event)，消费方可能是风控系统等等
	// GenSendAwardEvent

	// 发奖成功返回奖品信息
	return awardInfo.Name, nil
}
