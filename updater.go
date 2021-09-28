package cbg_notify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	SaveEndpoint     = "/api/v1/save"
	ContextVarName   = "ctx"
	MessageVarName   = "cnt"
	TimestampVarName = "ts"
)

type StorageResponse struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Messages []string    `json:"msg,omitempty"`
	Redirect string      `json:"redirect,omitempty"`
}

func SendToServerKeeptItem(cnf *RuntimeConfig, item *KeepItem) (ret StorageResponse, err error) {
	return SendToServer(cnf, item.Term, item.Ts)
}

func SendToServer(cnf *RuntimeConfig, msg string, ts time.Time) (ret StorageResponse, err error) {
	ret = StorageResponse{}
	client := http.Client{}

	data := url.Values{}
	data.Set(ContextVarName, cnf.StoreContext)
	data.Set(MessageVarName, msg)
	data.Set(TimestampVarName, strconv.Itoa(int(ts.Unix())))

	req, err := http.NewRequest("POST", cnf.StoreHost+SaveEndpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("RPS-TOKEN", cnf.StoreHash)


	res, err := client.Do(req)
	if err != nil {
		ErrHandle(err, "request send to server error")
		return ret, err
	}
	defer res.Body.Close()

	respBody, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		ErrHandle(errBody, "respond bode reading error")
		return ret, errBody
	}

	jsonErr := json.Unmarshal(respBody, &ret)
	if jsonErr != nil {
		ErrHandle(jsonErr, "json unmarshal error")
		return ret, jsonErr
	}

	return ret, nil
}
