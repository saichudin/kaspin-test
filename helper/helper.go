package helper

import (
	"encoding/json"
	"time"
	"log"
	"net/http"
	"strings"
    "github.com/itchyny/timefmt-go"
)

func Fetch(method string, url string, body string) (interface{}, error) {
	var err error
	var client = &http.Client{}
	var data interface{}
	log.Println(body)
	// url := "https://dev.nicepay.co.id//nicepay/direct/v2/registration"
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// func GetStrTimeStamp() string {
// 	timeNow := time.Now().Unix()
// 	return strconv.FormatInt(timeNow, 10)
// }

func GetTimeStamp() string {
	str := timefmt.Format(time.Now(), "%Y%m%d%H%M%S")
	return str
}

func TranslateInquiryStatusCc(statusCode string) string {
	switch status := statusCode; status {
	case "0":
		return "Success"
	case "1":
		return "Failed"
	case "2":
		return "Void / Return"
	case "9":
		return "Initialization / Reversal"
	default:
		return ""
	}
}
