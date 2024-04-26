package httpHelper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cuwand/pondasi/errors"
	"github.com/cuwand/pondasi/logger"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ContentTypeApplicationJson           = "application/json"
	ContentTypeApplicationFormUrlEncoded = "application/x-www-form-urlencoded"
)

var ErrRequestTimeout = errors.BadRequest("Request Timeout")

type HttpConfig struct {
	HttpClient *http.Client
	Timeout    *time.Duration
	Host       string
	Port       string
}

func (c HttpConfig) BuildBaseUrl() string {
	if len(c.Port) > 0 {
		return fmt.Sprintf("%s:%s", c.Host, c.Port)
	}

	return c.Host
}

type Converter func(body interface{}, logger logger.Logger, httpResponse *http.Response, err error) error

type HttpRequestPayload struct {
	Method      string
	Url         string
	QueryParams map[string]string
	Body        interface{}
	Result      interface{}
	Logger      logger.Logger
	Client      *http.Client
	TimeoutReq  *time.Duration // default 10s
	Header      *http.Header
	Converter   Converter
}

func HttpRequest(payload HttpRequestPayload) error {
	var bodyByte []byte

	if payload.Body != nil {
		reqBody, err := json.Marshal(payload.Body)

		if err != nil {
			return err
		}

		bodyByte = reqBody
	}

	req, err := http.NewRequest(payload.Method, payload.Url, bytes.NewBuffer(bodyByte))

	if err != nil {
		return err
	}

	if payload.QueryParams != nil {
		q := req.URL.Query()

		for key, val := range payload.QueryParams {
			q.Set(key, val)
		}

		req.URL.RawQuery = q.Encode()
	}

	if payload.Header == nil {
		req.Header.Set("Content-Type", ContentTypeApplicationJson)
	} else {
		req.Header = *payload.Header
	}

	payload.Logger.Info(fmt.Sprintf("URL | %v", payload.Url))
	payload.Logger.Info(fmt.Sprintf("Method | %v", payload.Method))
	payload.Logger.Info(fmt.Sprintf("Header | %v", payload.Header))

	if payload.Body != nil {
		marshaledRequestBody, _ := json.Marshal(payload.Body)
		payload.Logger.Info(fmt.Sprintf("[Request] Body | %v", string(marshaledRequestBody)))
	}

	if payload.QueryParams != nil {
		payload.Logger.Info(fmt.Sprintf("[Request] QueryParam | %v", payload.QueryParams))
	}

	// --- do request

	timeout := 10 * time.Second

	if payload.TimeoutReq != nil {
		timeout = *payload.TimeoutReq
	}

	newClient := http.Client{
		Timeout: timeout,
	}

	if payload.Client != nil {
		newClient = *payload.Client
	}

	resp, err := newClient.Do(req)

	if payload.Converter != nil {
		return payload.Converter(payload.Result, payload.Logger, resp, err)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return ErrRequestTimeout
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	payload.Logger.Info(fmt.Sprintf("[Response] StatusCode | %v", resp.StatusCode))

	if resp.StatusCode != http.StatusOK {
		readResp, _ := ioutil.ReadAll(resp.Body)
		payload.Logger.Error(string(readResp))
		payload.Logger.Error(fmt.Sprintf("%v", resp.StatusCode))

		return errors.BadRequest("request error.")
	}

	readResp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	payload.Logger.Info(fmt.Sprintf("[Response] Body | %v", string(readResp)))

	if err := json.Unmarshal(readResp, &payload.Result); err != nil {
		return err
	}

	return nil
}

type HttpRequestWithResponsePayload struct {
	Method      string
	Url         string
	QueryParams map[string]string
	Body        interface{}
	Logger      logger.Logger
	Client      *http.Client
	TimeoutReq  *time.Duration // default 10s
	Header      *http.Header
	Converter   Converter
}

func HttpRequestWithResponse(payload HttpRequestWithResponsePayload) (*http.Response, error) {
	var bodyByte []byte

	if payload.Body != nil {
		reqBody, err := json.Marshal(payload.Body)

		if err != nil {
			return nil, err
		}

		bodyByte = reqBody
	}

	req, err := http.NewRequest(payload.Method, payload.Url, bytes.NewBuffer(bodyByte))

	if err != nil {
		return nil, err
	}

	if payload.QueryParams != nil {
		q := req.URL.Query()

		for key, val := range payload.QueryParams {
			q.Set(key, val)
		}

		req.URL.RawQuery = q.Encode()
	}

	if payload.Header == nil {
		req.Header.Set("Content-Type", ContentTypeApplicationJson)
	} else {
		req.Header = *payload.Header
	}

	// --- do request

	timeout := 10 * time.Second

	if payload.TimeoutReq != nil {
		timeout = *payload.TimeoutReq
	}

	newClient := http.Client{
		Timeout: timeout,
	}

	if payload.Client != nil {
		newClient = *payload.Client
	}

	resp, err := newClient.Do(req)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return resp, ErrRequestTimeout
	}

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func GenerateBasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type HttpRequestUrlEncodedPayload struct {
	Method     string
	Url        string
	FormData   map[string]string
	Result     interface{}
	Logger     logger.Logger
	Client     *http.Client
	TimeoutReq *time.Duration // default 10s
	Header     *http.Header
	Converter  Converter
}

func HttpRequestFormUrlEncoded(payload HttpRequestUrlEncodedPayload) error {
	data := url.Values{}
	if payload.FormData != nil {
		for key, val := range payload.FormData {
			data.Set(key, val)
		}
	}

	req, err := http.NewRequest(payload.Method, payload.Url, strings.NewReader(data.Encode()))

	if err != nil {
		return err
	}

	if payload.Header == nil {
		req.Header.Set("Content-Type", ContentTypeApplicationFormUrlEncoded)
	} else {
		req.Header = *payload.Header
	}

	payload.Logger.Info(fmt.Sprintf("URL | %v", payload.Url))
	payload.Logger.Info(fmt.Sprintf("Method | %v", payload.Method))
	payload.Logger.Info(fmt.Sprintf("[Request] FormData | %v", payload.FormData))

	// --- do request

	timeout := 10 * time.Second

	if payload.TimeoutReq != nil {
		timeout = *payload.TimeoutReq
	}

	newClient := http.Client{
		Timeout: timeout,
	}

	if payload.Client != nil {
		newClient = *payload.Client
	}

	resp, err := newClient.Do(req)

	if payload.Converter != nil {
		return payload.Converter(payload.Result, payload.Logger, resp, err)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return errors.BadRequest("request timeout.")
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	payload.Logger.Info(fmt.Sprintf("[Response] StatusCode | %v", resp.StatusCode))

	if resp.StatusCode != http.StatusOK {
		readResp, _ := io.ReadAll(resp.Body)
		payload.Logger.Error(string(readResp))
		payload.Logger.Error(fmt.Sprintf("%v", resp.StatusCode))

		return errors.BadRequest("request error.")
	}

	readResp, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	payload.Logger.Info(fmt.Sprintf("[Response] Body | %v", string(readResp)))

	if err := json.Unmarshal(readResp, &payload.Result); err != nil {
		return err
	}

	return nil
}
