package services

import "ddd-go-demo/domain/assets/entity"

var AssetsSvc *assetsSvc

// assetsSvc 应用层资产service
type assetsSvc struct {
}

// GetCoinHistory 获取用户金币流水信息，在CQE中属于query
func (s *assetsSvc) GetCoinHistory(userId uint64) ([]*entity.CoinHistoryRecord, error) {
	return nil, nil
}

// HandleSendCoinEvent 处理发放金币事件，事件可能来自mq，在CQE中属于event
func (s *assetsSvc) HandleSendCoinEvent(userId uint64, cnt uint32) error {
	return nil
}
