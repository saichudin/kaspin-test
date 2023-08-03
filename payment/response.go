package payment

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RegistrationResponse struct {
	TrxId       string `json:"transaction_id"`
	ReferenceNo string `json:"reference_no"`
	Amount      string `json:"amount"`
	TrxDate     string `json:"transaction_date"`
	TrxTime     string `json:"transaction_time"`
}

type RegistrationResponseNicepay struct {
	ResultCd    string
	ResultMsg   string
	TXid        string
	ReferenceNo string
	PayMethod   string
	Amt         string
	TransDt     string
	TransTm     string
	Description string
}

type InquiryResponse struct {
	TrxId       string `json:"transaction_id"`
	ReferenceNo string `json:"reference_no"`
	Amount      string `json:"amount"`
	Status      string `json:"status"`
}

type InquiryResponseNicepay struct {
	TXid        string
	IMid        string
	Currency    string
	Amt         string
	InstmntMon  string
	InstmntType string
	ReferenceNo string
	GoodsNm     string
	PayMethod   string
	BillingNm   string
	ReqDt       string
	ReqTm       string
	Status      string
	ResultCd    string
	ResultMsg   string
}
