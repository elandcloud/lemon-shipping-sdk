package best

import (
	"encoding/xml"
	"net/http"

	"github.com/pangpanglabs/goutils/httpreq"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

/*
1.base64 param
2.request api
*/
func Create(reqDto ReqCreateDto, custDto ReqCustomerDto) (int, RespBase, error) {
	var respDto RespBase
	reqDto.ServiceType = "TMS_CREATE_ORDER"
	bizData, err := xml.Marshal(reqDto.BizData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(bizData) + reqDto.PartnerKey
	reqDto.Sign, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = reqDto.ServiceType
	reqMap["partnerID"] = reqDto.PartnerID
	reqMap["bizData"] = string(bizData)
	reqMap["sign"] = reqDto.Sign

	data := base.JoinMapString(reqMap)

	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.XmlType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	return statusCode, respDto, nil
}

func Cancel(reqDto ReqCancelDto, custDto ReqCustomerDto) (int, RespBase, error) {
	var respDto RespBase
	reqDto.ServiceType = "TMS_CANCEL_ORDER"
	bizData, err := xml.Marshal(reqDto.BizData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(bizData) + reqDto.PartnerKey
	reqDto.Sign, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = reqDto.ServiceType
	reqMap["partnerID"] = reqDto.PartnerID
	reqMap["bizData"] = string(bizData)
	reqMap["sign"] = reqDto.Sign

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.XmlType
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
	reqDto.ServiceType = "TMS_TRACE_QUERY"
	bizData, err := xml.Marshal(reqDto.BizData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(bizData) + reqDto.PartnerKey
	reqDto.Sign, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = reqDto.ServiceType
	reqMap["partnerID"] = reqDto.PartnerID
	reqMap["bizData"] = string(bizData)
	reqMap["sign"] = reqDto.Sign

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.XmlType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}
