package args

var (
	LogisticsStateType map[int]string
)

func init() {
	LogisticsStateType = map[int]string{
		0: "仓库准备中",
		1: "已打包",
		2: "交付物流",
		3: "运输中",
	}
}

type LogisticsState struct {
	Id            int64  `json:"id"`
	LogisticsCode string `json:"logistics_code"`
	State         string `json:"state"`
	Description   string `json:"description"`
	Flag          string `json:"flag"`
	Operator      string `json:"operator"`
	CreateTime    string `json:"create_time"`
	Location      string `json:"location"`
}

type LogisticsRecord struct {
	Courier     string           `json:"courier"`
	CourierType string           `json:"courier_type"`
	ReceiveType string           `json:"receive_type"`
	Customer    CustomerInfo     `json:"customer"`
	Goods       string           `json:"goods"`
	StateList   []LogisticsState `json:"state_list"`
}

type CustomerInfo struct {
	SendUser     string `json:"send_user"`
	SendAddr     string `json:"send_addr"`
	SendPhone    string `json:"send_phone"`
	SendTime     string `json:"send_time"`
	ReceiveUser  string `json:"receive_user"`
	ReceiveAddr  string `json:"receive_addr"`
	ReceivePhone string `json:"receive_phone"`
}
