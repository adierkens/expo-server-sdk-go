package expo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ExpoPushMessage struct {
	To         string                 `json:"to"`
	Data       map[string]interface{} `json:"data"`
	Title      string                 `json:"title"`
	Body       string                 `json:"body"`
	Sound      string                 `json:"sound"` // 'default' | null
	TTL        int32                  `json:"ttl"`
	Expiration int32                  `json:"expiration"`
	Priority   string                 `json:"priority"` // 'default' | 'normal' | 'high'
	Badge      int32                  `json:"badge"`
}

type ApiResult struct {
	Data   ExpoPushResponse `json:"data"`
	Errors []ApiError       `json:"errors"`
}

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

type ExpoPushResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func IsExpoPushToken(token string) bool {
	return strings.HasPrefix(token, "ExponentPushToken[") && strings.HasSuffix(token, "]")
}

const DEFAULT_HOST = "https://exp.host"
const DEFAULT_API_URL = DEFAULT_HOST + "/--/api/v2"

const DEFAULT_PRIORITY = "default"
const HIGH_PRIORITY = "high"
const NORMAL_PRIORITY = "normal"

type Expo struct {
	Host       string
	BaseAPIUrl string
}

func NewExpo(params ...string) *Expo {
	instance := Expo{DEFAULT_HOST, DEFAULT_API_URL}

	if len(params) > 0 {
		instance.Host = params[0]
	}

	if len(params) > 1 {
		instance.BaseAPIUrl = params[1]
	}

	return &instance
}

func NewExpoPushMessage() *ExpoPushMessage {
	message := &ExpoPushMessage{}

	message.Data = make(map[string]interface{})
	message.Sound = "default"
	message.Priority = DEFAULT_PRIORITY
	message.Badge = 1

	return message
}

func (expo *Expo) SendPushNotification(message *ExpoPushMessage) (response *ExpoPushResponse, err error) {
	jsonPayload, err := json.Marshal(message)

	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(jsonPayload)

	req, err := http.NewRequest("POST", expo.BaseAPIUrl+"/push/send", payload)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res := ApiResult{}
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &res)

	if len(res.Errors) > 0 {
		return &res.Data, &res.Errors[0]
	}

	return &res.Data, nil
}
