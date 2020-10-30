package repository

import (
	"gitee.com/cristiane/micro-mall-logistics/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

func CreateLogisticsRecord(tx *xorm.Session, model *mysql.LogisticsRecord) error {
	_, err := tx.Table(mysql.TableLogisticsRecord).Insert(model)
	return err
}

func AddLogisticsRecord(model *mysql.LogisticsRecord) error {
	_, err := kelvins.XORM_DBEngine.Table(mysql.TableLogisticsRecord).Insert(model)
	return err
}

func UpdateLogisticsRecord(where interface{}, maps interface{}) error {
	_, err := kelvins.XORM_DBEngine.Table(mysql.TableLogisticsRecord).Where(where).Update(maps)
	return err
}

func GetLogisticsRecordList(selectSql string, where interface{}, orderByAsc, orderByDesc []string, pageSize, pageNum int) ([]mysql.LogisticsRecord, int64, error) {
	var result = make([]mysql.LogisticsRecord, 0)
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
