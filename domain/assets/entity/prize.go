package entity

import (
	"ddd-go-demo/domain/assets/vo"
	"time"

	json "github.com/json-iterator/go"
)

// PrizeRecord 用户实物奖品信息
type PrizeRecord struct {
	Id          uint64
	UserId      uint64
	PrizeId     uint32 // 奖品id
	Status      uint32 // 礼物状态，1-待发奖，2-已发放，3-已领取
	ReceiveInfo string // 收货信息, 在数据库中存储的是序列化后的字符串信息
	CreateTime  int64
	UpdateTime  int64
}

// SetReceiveInfo 更新收货信息
func (r *PrizeRecord) SetReceiveInfo(info *vo.ReceiveInfo) error {
	// 将收货信息序列化为字符串保存
	receiveInfo, err := json.MarshalToString(info)
	if err != nil {
		return err
	}
	r.ReceiveInfo = receiveInfo
	return nil
}

// GetReceiveInfo 获取收货地址信息
func (r *PrizeRecord) GetReceiveInfo() (*vo.ReceiveInfo, error) {
	if r.ReceiveInfo == "" {
		return nil, nil
	}

	var ret *vo.ReceiveInfo
	err := json.UnmarshalFromString(r.ReceiveInfo, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func NewPrizeRecord(userId uint64, prizeId uint32) *PrizeRecord {
	return &PrizeRecord{
		UserId:     userId,
		PrizeId:    prizeId,
		Status:     1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
}
