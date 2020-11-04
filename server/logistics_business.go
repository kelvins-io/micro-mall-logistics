package server

import (
	"context"
	"gitee.com/cristiane/micro-mall-logistics/pkg/code"
	"gitee.com/cristiane/micro-mall-logistics/proto/micro_mall_logistics_proto/logistics_business"
	"gitee.com/cristiane/micro-mall-logistics/service"
	"gitee.com/kelvins-io/common/errcode"
	"strconv"
)

type LogisticsServer struct {
}

func NewLogisticsServer() *LogisticsServer {
	return &LogisticsServer{}
}

func (l *LogisticsServer) ApplyLogistics(ctx context.Context, req *logistics_business.ApplyLogisticsRequest) (*logistics_business.ApplyLogisticsResponse, error) {
	// 参数检查，违禁物品检查

	// 创建物流订单
	result := &logistics_business.ApplyLogisticsResponse{
		Common: &logistics_business.CommonResponse{
			Code: logistics_business.RetCode_SUCCESS,
			Msg:  "",
		},
		LogisticsCode: "",
	}
	id, retCode := service.CreateRecord(ctx, req)
	if retCode != code.Success {
		switch retCode {
		case code.LogisticsCodeExist:
			result.Common.Code = logistics_business.RetCode_LOGISTICS_CODE_EXIST
			result.Common.Msg = errcode.GetErrMsg(retCode)
		case code.LogisticsCodeNotExist:
			result.Common.Code = logistics_business.RetCode_LOGISTICS_CODE_NOT_EXIST
			result.Common.Msg = errcode.GetErrMsg(retCode)
		default:
			result.Common.Code = logistics_business.RetCode_ERROR
			result.Common.Msg = errcode.GetErrMsg(retCode)
		}
		return result, nil
	}
	result.LogisticsCode = id
	return result, nil
}

func (l *LogisticsServer) QueryRecord(ctx context.Context, req *logistics_business.QueryRecordRequest) (*logistics_business.QueryRecordResponse, error) {
	result := &logistics_business.QueryRecordResponse{
		Common: &logistics_business.CommonResponse{
			Code: logistics_business.RetCode_SUCCESS,
			Msg:  "",
		},
	}
	record, retCode := service.QueryRecord(ctx, req)
	if retCode != code.Success {
		result.Common.Code = logistics_business.RetCode_ERROR
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	result.Courier = record.Courier
	result.CourierType = record.CourierType
	result.ReceiveType = record.ReceiveType
	result.Customer = &logistics_business.CustomerInfo{
		SendUser:     record.Customer.SendUser,
		SendAddr:     record.Customer.SendAddr,
		SendPhone:    record.Customer.SendPhone,
		SendTime:     record.Customer.SendTime,
		ReceiveUser:  record.Customer.ReceiveUser,
		ReceiveAddr:  record.Customer.ReceiveAddr,
		ReceivePhone: record.Customer.ReceivePhone,
	}
	result.Goods = record.Goods
	result.StateList = make([]*logistics_business.LogisticsState, 0)
	for i := 0; i < len(record.StateList); i++ {
		row := record.StateList[i]
		state, _ := strconv.Atoi(row.State)
		result.StateList[i] = &logistics_business.LogisticsState{
			Id:            row.Id,
			LogisticsCode: row.LogisticsCode,
			State:         logistics_business.LogisticsStateType(state),
			Description:   row.Description,
			Flag:          row.Flag,
			Operator:      row.Operator,
			CreateTime:    row.CreateTime,
		}
	}
	return result, nil
}

func (l *LogisticsServer) UpdateState(ctx context.Context, req *logistics_business.UpdateStateRequest) (*logistics_business.UpdateStateResponse, error) {
	result := &logistics_business.UpdateStateResponse{Common: &logistics_business.CommonResponse{
		Code: logistics_business.RetCode_SUCCESS,
		Msg:  "",
	}}
	retCode := service.UpdateState(ctx, req)
	if retCode != code.Success {
		if retCode == code.LogisticsCodeNotExist {
			result.Common.Code = logistics_business.RetCode_LOGISTICS_CODE_NOT_EXIST
			result.Common.Msg = errcode.GetErrMsg(code.LogisticsCodeNotExist)
		} else {
			result.Common.Code = logistics_business.RetCode_ERROR
			result.Common.Msg = errcode.GetErrMsg(code.ErrorServer)
		}
		return result, nil
	}
	return result, nil
}
