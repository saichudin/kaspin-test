package payment

import (
	"encoding/json"
	"saichudin/kaspin/helper"
)

func RegistrationService(form *RegistrationForm) (RegistrationResponse, string) {
	// prepare payload data
	timeNowStr := helper.GetTimeStamp()
	refNo := "MerchantRefNo" + timeNowStr
	amount := form.Amount
	tokenString := timeNowStr + IMid + refNo + amount + ApiKey
	token := GenerateTokenString(tokenString)

	var payload RegistrationRequest
	payload.Timestamp = timeNowStr
	payload.IMid = IMid
	payload.PayMethod = "01"
	payload.Currency = "IDR"
	payload.Amt = amount
	payload.ReferenceNo = refNo
	payload.GoodsNm = form.Item
	payload.DbProcessUrl = "https://merchant.com/api/dbProcessUrl/Notif"
	payload.Description = ""
	payload.MerchantToken = string(token)
	payload.InstmntType = "2"
	payload.InstmntMon = "1"
	payload.RecurrOpt = "0"
	payload.VacctValidDt = "20240927"
	payload.VacctValidTm = "121517"
	payload.BillingNm = form.Name
	payload.BillingPhone = form.Phone
	payload.BillingEmail = form.Email
	payload.BillingAddr = form.Address
	payload.BillingCity = form.City
	payload.BillingState = form.Province
	payload.BillingPostCd = form.Postal
	payload.BillingCountry = form.Country

	url := "https://dev.nicepay.co.id/nicepay/direct/v2/registration"
	body, _ := json.Marshal(payload)

	var fetch interface{}
	fetch, _ = helper.Fetch("POST", url, string(body))

	jsonData, _ := json.Marshal(fetch)
	var structData RegistrationResponseNicepay
	json.Unmarshal(jsonData, &structData)

	var data RegistrationResponse
	data.TrxId = structData.TXid
	data.ReferenceNo = structData.ReferenceNo
	data.Amount = structData.Amt
	data.TrxDate = structData.TransDt
	data.TrxTime = structData.TransTm

	return data, structData.ResultMsg
}

func InquiryService(form *InquiryForm) (InquiryResponse, string) {
	// prepare payload data
	timeNowStr := helper.GetTimeStamp()
	refNo := form.ReferenceNo
	amount := form.Amount
	tokenString := timeNowStr + IMid + refNo + amount + ApiKey
	token := GenerateTokenString(tokenString)

	var payload InquiryRequest
	payload.TimeStamp = timeNowStr
	payload.MerchantToken = token
	payload.ReferenceNo = refNo
	payload.TXid = form.TransactionId
	payload.Amt = amount
	payload.IMid = IMid

	url := "https://dev.nicepay.co.id/nicepay/direct/v2/inquiry"
	body, _ := json.Marshal(payload)

	var fetch interface{}
	fetch, _ = helper.Fetch("POST", url, string(body))

	jsonData, _ := json.Marshal(fetch)
	var structData InquiryResponseNicepay
	json.Unmarshal(jsonData, &structData)

	var data InquiryResponse
	data.TrxId = structData.TXid
	data.ReferenceNo = structData.ReferenceNo
	data.Amount = structData.Amt
	data.Status = helper.TranslateInquiryStatusCc(structData.Status)

	return data, structData.ResultMsg
}