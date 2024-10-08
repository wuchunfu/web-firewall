package nft

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shirou/gopsutil/v3/net"
	"server/internal/model"
	"server/internal/service"
	"time"

	"server/api/nft/v1"
)

func (c *ControllerV1) GetPolicy(ctx context.Context, req *v1.GetPolicyReq) (res *v1.GetPolicyRes, err error) {

	if req.Chain > 9 || req.Chain < 1 {
		return nil, errors.New(fmt.Sprintf("错误的请求id: %d，取值范围：1 出站策略 2 入站策略 3 目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 转发策略 8 转发流控 9 ip黑白名单", req.Chain))
	}

	var network []model.Network

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, stat := range interfaces {
		network = append(network, model.Network{
			Index: stat.Index,
			Name:  stat.Name,
		})
		g.Log().Debug(ctx, stat)
	}

	return &v1.GetPolicyRes{
		Data:      *(service.Nft().GetChainList(req.Chain)),
		Network:   network,
		Total:     len(*(service.Nft().GetChainList(req.Chain))),
		Timestamp: time.Now().Unix(),
	}, nil
}
