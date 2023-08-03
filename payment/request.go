package payment

import (
	"crypto/sha256"
	"encoding/hex"
)

const IMid = "IONPAYTEST"
const ApiKey = "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A=="

type RegistrationRequest struct {
	Timestamp      string `json:"timeStamp"`
	IMid           string `json:"iMid"`
	PayMethod      string `json:"payMethod"`
	Currency       string `json:"currency"`
	Amt            string `json:"amt"`
	ReferenceNo    string `json:"referenceNo"`
	GoodsNm        string `json:"goodsNm"`
	DbProcessUrl   string `json:"dbProcessUrl"`
	Description    string `json:"description"`
	MerchantToken  string `json:"merchantToken"`
	InstmntType    string `json:"instmntType"`
	InstmntMon     string `json:"instmntMon"`
	RecurrOpt      string `json:"recurrOpt"`
	VacctValidDt   string `json:"vacctValidDt"`
	VacctValidTm   string `json:"vacctValidTm"`
	BillingNm      string `json:"billingNm"`
	BillingPhone   string `json:"billingPhone"`
	BillingEmail   string `json:"billingEmail"`
	BillingAddr    string `json:"billingAddr"`
	BillingCity    string `json:"billingCity"`
	BillingState   string `json:"billingState"`
	BillingPostCd  string `json:"billingPostCd"`
	BillingCountry string `json:"billingCountry"`
}

type RegistrationForm struct {
	Item     string `json:"item"`
	Amount   string `json:"amount"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
	Postal   string `json:"postal"`
	Country  string `json:"country"`
}

type InquiryRequest struct {
	TimeStamp     string `json:"timeStamp"`
	MerchantToken string `json:"merchantToken"`
	ReferenceNo   string `json:"referenceNo"`
	TXid          string `json:"tXid"`
	Amt           string `json:"amt"`
	IMid          string `json:"iMid"`
}

type InquiryForm struct {
	ReferenceNo   string `json:"reference_no"`
	TransactionId string `json:"transaction_id"`
	Amount        string `json:"amount"`
}

func GenerateTokenString(tokenString string) string {
	h := sha256.New()
	h.Write([]byte(tokenString))
	token := hex.EncodeToString(h.Sum(nil))

	return token
}
