package service

import (
	"go-my-life/internal/domain/entity/health"
	"go-my-life/internal/domain/repository/healthdb"
)

type PhysicalQueryParam struct {
	UserId    int64 `json:"userId"`
	PageSize  int   `json:"pageSize"`
	PageIndex int   `json:"pageIndex"`
}

func QueryPhysicals(param *PhysicalQueryParam) ([]health.Physical, int64) {
	physicals, total := healthdb.QueryPhysicals(param.UserId, param.PageSize, param.PageIndex)
	if len(physicals) == 0 {
		return nil, 0
	}
	var res []health.Physical
	for _, physical := range physicals {
		one := health.Physical{
			Id:         physical.Id,
			UserId:     physical.UserId,
			Height:     physical.Height,
			Weight:     physical.Weight,
			Desc:       physical.Desc,
			RecordTime: physical.RecordTime,
		}
		res = append(res, one)
	}

	return res, total
}
