package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-logistics/model/args"
	"gitee.com/cristiane/micro-mall-logistics/model/mysql"
	"gitee.com/cristiane/micro-mall-logistics/pkg/code"
	"gitee.com/cristiane/micro-mall-logistics/pkg/util"
	"gitee.com/cristiane/micro-mall-logistics/pkg/util/email"
	"gitee.com/cristiane/micro-mall-logistics/proto/micro_mall_logistics_proto/logistics_business"
	"gitee.com/cristiane/micro-mall-logistics/repository"
	"gitee.com/cristiane/micro-mall-logistics/vars"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"time"
)

func CreateRecord(ctx context.Context, req *logistics_business.ApplyLogisticsRequest) (result string, retCode int) {
	result = ""
	retCode = code.Success
	// 检查物流记录是否已存在
	retCode = checkLogisticsRecord(ctx, req)
	if retCode != code.Success {
		return
	}
	// 创建物流记录
	logisticsCode, goods, retCode := createLogisticsRecord(ctx, req)
	if retCode != code.Success {
		return
	}

	// 物流通知
	go createLogisticsRecordNotice(ctx, req, goods)

	result = logisticsCode
	return result, retCode
}

func createLogisticsRecord(ctx context.Context, req *logistics_business.ApplyLogisticsRequest) (result, goods string, retCode int) {
	// 处理物流
	retCode = code.Success
	logisticsCode := util.GetUUID()
	goods = json.MarshalToStringNoError(req.Goods)
	orderLogistics := &mysql.OrderLogistics{
		LogisticsCode: logisticsCode,
		OrderCode:     req.OutTradeNo,
		State:         0,
		Courier:       req.Courier,
		FromAddress:   req.Customer.SendAddr,
		ToAddress:     req.Customer.ReceiveAddr,
		Sender:        req.Customer.SendUser,
		Receiver:      req.Customer.ReceiveUser,
		ReceiverPhone: req.Customer.ReceivePhone,
		SendTime:      req.SendTime,
		SenderPhone:   req.Customer.SendPhone,
		TransportKind: fmt.Sprintf("%d", req.CourierType),
		ReceiverKind:  fmt.Sprintf("%d", req.ReceiveType),
		Goods:         goods,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	logisticsRecord := &mysql.LogisticsRecord{
		LogisticsCode: logisticsCode,
		Location:      "亚洲一号仓库",
		State:         int(logistics_business.LogisticsStateType_INIT),
		Description:   "物流单创建成功",
		Flag:          "快递员揽件",
		Operator:      "微商城",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	session := kelvins.XORM_DBEngine.NewSession()
	err := session.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CreateLogisticsRecord Begin err: %v")
		retCode = code.ErrorServer
		return
	}
	err = repository.CreateLogisticsRecord(session, logisticsRecord)
	if err != nil {
		errRoll := session.Rollback()
		if errRoll != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateLogisticsRecord Rollback err: %v", errRoll)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateLogisticsRecord err: %v, model: %v", err, json.MarshalToStringNoError(logisticsRecord))
		retCode = code.ErrorServer
		return
	}
	err = repository.CreateOrderLogistics(session, orderLogistics)
	if err != nil {
		errRoll := session.Rollback()
		if errRoll != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateOrderLogistics Rollback err: %v", errRoll)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateOrderLogistics err: %v, model: %v", err, json.MarshalToStringNoError(orderLogistics))
		retCode = code.ErrorServer
		return
	}
	errCommit := session.Commit()
	if errCommit != nil {
		kelvins.ErrLogger.Errorf(ctx, "CreateOrderLogistics Commit err: %v", errCommit)
		retCode = code.ErrorServer
		return
	}
	return logisticsCode, goods, retCode
}

func createLogisticsRecordNotice(ctx context.Context, req *logistics_business.ApplyLogisticsRequest, goods string) {
	// 触发消息通知
	emailNotice := fmt.Sprintf(args.LogisticsNotice, req.Customer.ReceiveUser, req.OutTradeNo, goods, req.Customer.SendUser, req.Courier)
	for _, receiver := range vars.EmailNoticeSetting.Receivers {
		err := email.SendEmailNotice(ctx, receiver, vars.AppName, emailNotice)
		if err != nil {
			kelvins.ErrLogger.Info(ctx, "createLogisticsRecordNotice SendEmailNotice err, emailNotice: %v", emailNotice)
		}
	}
}

func checkLogisticsRecord(ctx context.Context, req *logistics_business.ApplyLogisticsRequest) int {
	// 检查订单号是否已申请物流
	where := map[string]interface{}{
		"order_code": req.OutTradeNo,
	}
	orderLogisticsDB, err := repository.GetOrderLogistics("order_code", where)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetOrderLogistics err: %v, where: %v", err, json.MarshalToStringNoError(where))
		return code.ErrorServer
	}
	if orderLogisticsDB.OrderCode != "" {
		return code.LogisticsCodeExist
	}
	return code.Success
}

func QueryRecord(ctx context.Context, req *logistics_business.QueryRecordRequest) (*args.LogisticsRecord, int) {
	result := &args.LogisticsRecord{
		Courier:     "",
		CourierType: "",
		ReceiveType: "",
		Customer:    args.CustomerInfo{},
		Goods:       "",
		StateList:   make([]args.LogisticsState, 0),
	}
	retCode := code.Success
	where := map[string]interface{}{
		"logistics_code": req.LogisticsCode,
	}
	recordList, total, err := repository.GetLogisticsRecordList("*", where, nil, nil, 0, 0)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetLogisticsRecordList err: %v, where: %v", err, json.MarshalToStringNoError(where))
		retCode = code.ErrorServer
		return result, retCode
	}
	result.StateList = make([]args.LogisticsState, total)
	for i := 0; i < len(recordList); i++ {
		result.StateList[i] = args.LogisticsState{
			Id:            recordList[i].Id,
			LogisticsCode: recordList[i].LogisticsCode,
			Location:      recordList[i].Location,
			State:         args.LogisticsStateType[recordList[i].State],
			Description:   recordList[i].Description,
			Flag:          recordList[i].Flag,
			Operator:      recordList[i].Operator,
			CreateTime:    util.ParseTimeOfStr(recordList[i].CreateTime.Unix()),
		}
	}
	orderLogisticsList, total, err := repository.GetOrderLogisticsList("*", where, nil, nil, 0, 0)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetOrderLogisticsList err: %v, where: %v", err, json.MarshalToStringNoError(where))
		retCode = code.ErrorServer
		return result, retCode
	}
	if len(orderLogisticsList) > 0 {
		row := orderLogisticsList[0]
		result.Courier = row.Courier
		result.ReceiveType = row.ReceiverKind
		result.Customer = args.CustomerInfo{
			SendUser:     row.Sender,
			SendAddr:     row.FromAddress,
			SendPhone:    row.SenderPhone,
			SendTime:     row.SendTime,
			ReceiveUser:  row.Receiver,
			ReceiveAddr:  row.ToAddress,
			ReceivePhone: row.ReceiverPhone,
		}
		result.Goods = row.Goods
	}
	return result, retCode
}

func UpdateState(ctx context.Context, req *logistics_business.UpdateStateRequest) int {
	retCode := code.Success
	where := " logistics_code = '" + req.LogisticsCode + "'"
	recordDB, err := repository.GetOrderLogistics("*", where)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetOrderLogistics err: %v, where: %v", err, json.MarshalToStringNoError(where))
		retCode = code.ErrorServer
		return retCode
	}
	if recordDB.LogisticsCode != req.LogisticsCode {
		retCode = code.LogisticsCodeNotExist
		return retCode
	}
	record := &mysql.LogisticsRecord{
		LogisticsCode: req.LogisticsCode,
		Location:      req.State.Location,
		State:         int(req.State.State),
		Description:   req.State.Description,
		Flag:          req.State.Flag,
		Operator:      req.State.Operator,
	}
	err = repository.AddLogisticsRecord(record)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "AddLogisticsRecord err: %v, record: %v", err, json.MarshalToStringNoError(record))
		retCode = code.ErrorServer
		return retCode
	}
	return retCode
}
