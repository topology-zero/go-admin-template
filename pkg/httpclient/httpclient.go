package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type method string

const (
	GET  = method("GET")
	POST = method("POST")
)

type Request struct {
	method string
	url    string
	body   io.Reader
	header map[string]string

	client *http.Client
}

type HTTPOption func(r *Request)

// NewHttpRequest 构建一个 Request 对象
func NewHttpRequest(m method, url string, opt ...HTTPOption) *Request {
	r := &Request{
		method: string(m),
		url:    url,
		client: &http.Client{},
	}
	for _, option := range opt {
		option(r)
	}
	return r
}

// DoHttpRequest 发送HTTP请求
func (r *Request) DoHttpRequest() ([]byte, error) {
	request, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.header {
		request.Header.Add(k, v)
	}

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("status code != 200")
	}

	respStr, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	return respStr, nil
}

func WithTimeout(d time.Duration) HTTPOption {
	return func(r *Request) {
		r.client.Timeout = d
	}
}

func WithJson(body any) HTTPOption {
	return func(r *Request) {
		marshal, _ := json.Marshal(body)
		r.body = bytes.NewBuffer(marshal)
		r.header["Content-Type"] = "application/json;charset=UTF-8"
	}
}
func WithFromData(body map[string]any) HTTPOption {
	return func(r *Request) {
		r.body = strings.NewReader(Map2Str(body))
		r.header["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8"
	}
}

func WithHeader(key, val string) HTTPOption {
	return func(r *Request) {
		r.header[key] = val
	}
}

func Map2Str(m map[string]any) string {
	var strArr []string
	for k, v := range m {
		strArr = append(strArr, k+"="+ConvertString(v))
	}
	return strings.Join(strArr, "&")
}

func ConvertString(value any) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case int64:
		return strconv.Itoa(int(value))
	case json.Number:
		return value.String()
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	default:
		return ""
	}
}
