package joinquant

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	// stock, fund, index, futures, etf, lof, fja, fjb, QDII_fund, lof, fja, fjb, QDII_fund, open_fund, bond_fund, stock_fund, money_market_fund, mixture_fund, options
	CodeStock  = "stock"
	CodeFund   = "fund"
	CodeFuture = "futures"
	CodeEtf    = "etf"
	CodeIndex  = "index"
)

var (
	account  = "<account>"
	password = "<password>"
	ramCache = cache.New(time.Hour, time.Hour*24)
)

func init() {
	account = os.Getenv("JOINQUANT_SDK_ACCOUNT")
	password = os.Getenv("JOINQUANT_SDK_PASSWORD")
}

func request(body map[string]interface{}) string {
	url := "https://dataapi.joinquant.com/apis"
	bodyStr, err := json.Marshal(body)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(bodyStr)))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}

func GetToken() string {
	cacheKey := "ws_joinquant_token"

	if x, found := ramCache.Get(cacheKey); found {
		return x.(string)
	}

	token := request(map[string]interface{}{
		"method": "get_token",
		"mob":    account,
		"pwd":    password,
	})

	ramCache.Set(cacheKey, token, time.Hour)

	return token
}

func GetBillboardList() string {
	body := map[string]interface{}{
		"method": "get_billboard_list",
		"token":  GetToken(),
		//"count":    1000,
		//"end_date": "2020-07-01",
		//"date": "2018-11-29",
	}
	// @todo 无论放什么参数都无效
	return request(body)
}

func GetAllSecurities(code string) string {
	body := map[string]interface{}{
		"method": "get_all_securities",
		"token":  GetToken(),
		"code":   code,
	}
	return request(body)
}

func GetIndexStocks(code string, date string) string {
	body := map[string]interface{}{
		"method": "get_index_stocks",
		"token":  GetToken(),
		"code":   code,
		"date":   date,
	}
	return request(body)
}

func GetIndexWeights(code string, date string) string {
	body := map[string]interface{}{
		"method": "get_index_weights",
		"token":  GetToken(),
		"code":   code,
		"date":   date,
	}
	return request(body)
}

func GetLockedShares(codes []string, date string, endDate string) string {
	body := map[string]interface{}{
		"method": "get_locked_shares",
		"token":  GetToken(),
	}
	if len(codes) > 0 {
		body["code"] = strings.Join(codes, ",")
	}

	body["date"] = date
	body["end_date"] = endDate

	return request(body)
}

func GetFundamentals(codes []string, table string, date string) string {
	body := map[string]interface{}{
		"method": "get_fundamentals",
		"token":  GetToken(),
		"table":  table,
	}
	if len(codes) > 0 {
		body["code"] = strings.Join(codes, ",")
	}

	body["date"] = date
	// balance，income，cash_flow，indicator，valuation，bank_indicator，security_indicator，insurance_indicator
	return request(body)
}
