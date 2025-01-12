package jd

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pangpanglabs/goutils/httpreq"
	"github.com/relax-space/go-kit/base"
	"github.com/sirupsen/logrus"
)

func ECreate(reqDto ReqECreateDto, custDto ReqCustomerDto) (int, RespECreateDto, error) {
	var respDto RespECreateDto
	reqDto.Method = "jingdong.ldop.waybill.receive"
	reqDto.Timestamp = time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	var err error
	jsonBytes, err := json.Marshal(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	var mapData = make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &mapData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}

	reqDto.Sign, err = SignJD(mapData, custDto.AppSecret)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	mapData["sign"] = reqDto.Sign
	url := custDto.Url + "?" + base.JoinMapObjectEncode(mapData)
	logrus.WithField("url", url).Info("url")
	req := httpreq.New(http.MethodGet, url, nil, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})

	var respInterface interface{}
	statusCode, err := req.Call(&respInterface)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	err = Decode(respInterface, &respDto)
	if err != nil {
		return statusCode, respDto, err
	}
	errorResponse := ErrorResponse{}
	err = Decode(respInterface, &errorResponse)
	if err != nil {
		return statusCode, respDto, err
	}
	respDto.ErrorResponse = errorResponse
	return statusCode, respDto, nil
}

func Create(reqDto ReqCreateDto, custDto ReqCustomerDto) (int, RespCreateDto, error) {
	var respDto RespCreateDto
	reqDto.Method = "jingdong.etms.waybill.send"
	reqDto.Timestamp = time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	var err error
	jsonBytes, err := json.Marshal(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	var mapData = make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &mapData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}

	reqDto.Sign, err = SignJD(mapData, custDto.AppSecret)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	mapData["sign"] = reqDto.Sign
	url := custDto.Url + "?" + base.JoinMapObjectEncode(mapData)
	logrus.WithField("url", url).Info("url")
	req := httpreq.New(http.MethodGet, url, nil, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	var respInterface interface{}
	statusCode, err := req.Call(&respInterface)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	err = Decode(respInterface, &respDto)
	if err != nil {
		return statusCode, respDto, err
	}
	errorResponse := ErrorResponse{}
	err = Decode(respInterface, &errorResponse)
	if err != nil {
		return statusCode, respDto, err
	}
	respDto.ErrorResponse = errorResponse
	return statusCode, respDto, nil
}

func Cancel(reqDto ReqCancelDto, custDto ReqCustomerDto) (int, RespCancelDto, error) {
	var respDto RespCancelDto
	reqDto.Method = "jingdong.ldop.pickup.cancel"
	reqDto.Timestamp = time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")

	var err error
	jsonBytes, err := json.Marshal(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	var mapData = make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &mapData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}

	reqDto.Sign, err = SignJD(mapData, custDto.AppSecret)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	mapData["sign"] = reqDto.Sign

	url := custDto.Url + "?" + base.JoinMapObjectEncode(mapData)
	logrus.WithField("url", url).Info("url")
	req := httpreq.New(http.MethodGet, url, nil, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	var respInterface interface{}
	statusCode, err := req.Call(&respInterface)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	err = Decode(respInterface, &respDto)
	if err != nil {
		return statusCode, respDto, err
	}
	errorResponse := ErrorResponse{}
	err = Decode(respInterface, &errorResponse)
	if err != nil {
		return statusCode, respDto, err
	}
	respDto.ErrorResponse = errorResponse
	return statusCode, respDto, nil
}
func CancelWayBill(reqDto ReqCancelWayBillDto, custDto ReqCustomerDto) (int, RespCancelWayBillDto, error) {
	var respDto RespCancelWayBillDto
	reqDto.Method = "jingdong.ldop.delivery.provider.cancelWayBill"
	reqDto.Timestamp = time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")

	var err error
	jsonBytes, err := json.Marshal(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	var mapData = make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &mapData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}

	reqDto.Sign, err = SignJD(mapData, custDto.AppSecret)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	mapData["sign"] = reqDto.Sign

	url := custDto.Url + "?" + base.JoinMapObjectEncode(mapData)
	logrus.WithField("url", url).Info("url")
	req := httpreq.New(http.MethodGet, url, nil, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	var respInterface interface{}
	statusCode, err := req.Call(&respInterface)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	err = Decode(respInterface, &respDto)
	if err != nil {
		return statusCode, respDto, err
	}
	errorResponse := ErrorResponse{}
	err = Decode(respInterface, &errorResponse)
	if err != nil {
		return statusCode, respDto, err
	}
	respDto.ErrorResponse = errorResponse
	return statusCode, respDto, nil
}

func Query(reqDto ReqQueryDto, custDto ReqCustomerDto) (int, RespQueryDto, error) {
	var respDto RespQueryDto
	reqDto.Method = "jingdong.ldop.receive.trace.get"
	reqDto.Timestamp = time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	var err error
	jsonBytes, err := json.Marshal(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	var mapData = make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &mapData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}

	reqDto.Sign, err = SignJD(mapData, custDto.AppSecret)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	mapData["sign"] = reqDto.Sign

	url := custDto.Url + "?" + base.JoinMapObjectEncode(mapData)
	req := httpreq.New(http.MethodGet, url, nil, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	var respInterface interface{}
	statusCode, err := req.Call(&respInterface)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	err = Decode(respInterface, &respDto)
	if err != nil {
		return statusCode, respDto, err
	}
	errorResponse := ErrorResponse{}
	err = Decode(respInterface, &errorResponse)
	if err != nil {
		return statusCode, respDto, err
	}
	respDto.ErrorResponse = errorResponse

	return statusCode, respDto, nil
}
