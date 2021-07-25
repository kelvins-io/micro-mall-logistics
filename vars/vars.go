package vars

import "gitee.com/kelvins-io/common/queue"

var (
	AppName               = ""
	EmailConfigSetting    *EmailConfigSettingS
	TradeOrderQueueServer *queue.MachineryQueue
	EmailNoticeSetting *EmailNoticeSettingS
)
