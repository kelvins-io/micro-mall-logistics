package repository

import (
	"gitee.com/cristiane/micro-mall-logistics/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

func CreateOrderLogistics(tx *xorm.Session, model *mysql.OrderLogistics) error {
	_, err := tx.Table(mysql.TableOrderLogistics).Insert(model)
	return err
}

func UpdateOrderLogistics(where interface{}, maps interface{}) error {
	_, err := kelvins.XORM_DBEngine.Table(mysql.TableLogisticsRecord).Where(where).Update(maps)
	return err
}

func GetOrderLogistics(selectSql string, where interface{}) (*mysql.OrderLogistics, error) {
	var model mysql.OrderLogistics
	_, err := kelvins.XORM_DBEngine.Table(mysql.TableOrderLogistics).Select(selectSql).Where(where).Get(&model)
	return &model, err
}

func GetOrderLogisticsList(selectSql string, where interface{}, orderByAsc, orderByDesc []string, pageSize, pageNum int) ([]mysql.OrderLogistics, int64, error) {
	var result = make([]mysql.OrderLogistics, 0)
	session := kelvins.XORM_DBEngine.Table(mysql.TableLogisticsRecord).
		Select(selectSql).
		Where(where).
		Asc(orderByAsc...).
		Desc(orderByDesc...)

	if pageSize > 0 && pageNum >= 1 {
		session = session.Limit(pageSize, (pageNum-1)*pageSize)
	}
	total, err := session.FindAndCount(&result)
	return result, total, err
}
