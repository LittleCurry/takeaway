package model



type Order struct {
	OrderId    string `json:"order_id"`
	Openid     string `json:"openid"`
	UserId     string `json:"user_id"`
	Status     int    `json:"status"`
	TotalFee   string `json:"total_fee"`
	CreateTime string `json:"create_time"`
}