// todo: ddd中的绝大部分概念已在assets模块中实现, 其余模块将一码带过

package services

var PrizeInventorySvc *prizeInventoryService

// prizeService 奖品库存service
type prizeInventoryService struct {
}

// InventoryEnough 库存是否充足
func (s *prizeInventoryService) InventoryIsEnough(prizeId uint32, cnt uint32) (bool, error) {
	return true, nil
}
