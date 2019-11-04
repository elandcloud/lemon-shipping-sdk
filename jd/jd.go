package jd

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pangpanglabs/goutils/httpreq"
	"github.com/relax-space/go-kit/base"
)

func Create(reqDto ReqCreateDto, custDto ReqCustomerDto) (int, RespCreateDto, error) {
	var respDto RespCreateDto
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

	req := httpreq.New(http.MethodGet, custDto.Url, base.JoinMapObjectEncode(mapData),
		func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.FormType
			return nil
		})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
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

	req := httpreq.New(http.MethodGet, custDto.Url, reqDto, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
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

	req := httpreq.New(http.MethodGet, custDto.Url, reqDto, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}
