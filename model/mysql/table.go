package mysql

import "time"

const (
	TableLogisticsRecord = "logistics_record"
	TableOrderLogistics  = "order_logistics"
	TableConfigKv        = "config_kv_store"
)

type LogisticsRecord struct {
	Id            int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	LogisticsCode string    `xorm:"comment('物流单号') index CHAR(40)"`
	Location      string    `xorm:"comment('位置') VARCHAR(255)"`
	State         int       `xorm:"default 0 comment('当前状态') TINYINT"`
	Description   string    `xorm:"comment('描述') TEXT"`
	Flag          string    `xorm:"comment('标记') VARCHAR(255)"`
	Operator      string    `xorm:"comment('操作员') index VARCHAR(512)"`
	CreateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}
type ConfigKvStore struct {
	Id          int       `xorm:"not null pk autoincr comment('主键') INT"`
	ConfigKey   string    `xorm:"not null comment('配置键') unique VARCHAR(255)"`
	ConfigValue string    `xorm:"not null comment('配置值') VARCHAR(255)"`
	Prefix      string    `xorm:"not null comment('配置前缀') VARCHAR(255)"`
	Suffix      string    `xorm:"not null comment('配置后缀') VARCHAR(255)"`
	Status      int       `xorm:"not null default 1 comment('是否启用 1是 0否') TINYINT"`
	IsDelete    int       `xorm:"not null default 0 comment('是否删除 1是 0否') TINYINT"`
	CreateTime  time.Time `xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}
type OrderLogistics struct {
	Id            int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	LogisticsCode string    `xorm:"not null comment('运单号') unique(logistics_code_order_code) CHAR(40)"`
	OrderCode     string    `xorm:"not null comment('订单ID') unique(logistics_code_order_code) index CHAR(40)"`
	State         int       `xorm:"comment('物流状态，0-已下单，1-已取消，2-延迟处理，3-仓库处理中，4-运输中，5-派送中，6-已签收') TINYINT"`
	Courier       string    `xorm:"comment('国内承运人') index VARCHAR(255)"`
	FromAddress   string    `xorm:"comment('发货地址') VARCHAR(255)"`
	ToAddress     string    `xorm:"comment('收获地址') VARCHAR(255)"`
	Sender        string    `xorm:"comment('发货人') VARCHAR(255)"`
	Receiver      string    `xorm:"comment('接收人') VARCHAR(255)"`
	ReceiverPhone string    `xorm:"comment('收货人联系方式') VARCHAR(255)"`
	SenderPhone   string    `xorm:"comment('发送人联系方式') VARCHAR(255)"`
	TransportKind string    `xorm:"comment('运送方式') VARCHAR(255)"`
	ReceiverKind  string    `xorm:"comment('收货方式') VARCHAR(255)"`
	Goods         string    `xorm:"comment('货物') TEXT"`
	SendTime      string    `xorm:"comment('派送时间') VARCHAR(255)"`
	CreateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}
