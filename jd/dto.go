package jd

/*
request or response parameter
*/
type ReqCreateDto struct {
	ReqBase
	RequestData ReqCreateDataDto `json:"360buy_param_json,omitempty"`
}
type ReqQueryDto struct {
	ReqBase
	RequestData ReqQueryDataDto `json:"360buy_param_json,omitempty"`
}
type ReqCancelDto struct {
	ReqBase
	RequestData ReqCancelDataDto `json:"360buy_param_json,omitempty"`
}

type RespCreateDto struct {
	Response RespCreateResponse `json:"jingdong_ldop_waybill_receive_response"`
}
type RespCreateResponse struct {
	Result RespCreateResult `json:"receiveorderinfo_result"`
}
type RespCreateResult struct {
	ResultCode      int            `json:"resultCode,omitempty"`
	ResultMessage   string         `json:"resultMessage,omitempty"`
	OrderId         string         `json:"orderId,omitempty"`
	DeliveryId      string         `json:"deliveryId,omitempty"`
	PromiseTimeType int            `json:"promiseTimeType,omitempty"`
	PreSortResult   *PreSortResult `json:"preSortResult,omitempty"`
	TransType       int            `json:"transType,omitempty"`
}
type RespQueryDto struct {
	Response RespQueryResponse `json:"jingdong_ldop_receive_trace_get_response"`
}
type RespQueryResponse struct {
	Result RespQueryResult `json:"querytrace_result"`
}
type RespQueryResult struct {
	Code     int        `json:"code,omitempty"`
	Messsage string     `json:"messsage,omitempty"`
	Data     []TraceDTO `json:"data,omitempty"`
}
type RespCancelDto struct {
	Response RespCancelResponse `json:"jingdong_ldop_pickup_cancel_response"`
}
type RespCancelResponse struct {
	Result RespCancelResult `json:"returnType"`
}
type RespCancelResult struct {
	StatusMessage string `json:"statusMessage,omitempty"`
	StatusCode    int    `json:"statusCode,omitempty"`
}

/*
dto
*/
type ReqBase struct {
	Method      string `json:"method,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	AppKey      string `json:"app_key,omitempty"`
	Sign        string `json:"sign,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Format      string `json:"format,omitempty"`
	V           string `json:"v,omitempty"`
}

type ReqCreateDataDto struct {
	SalePlat     string `json:"salePlat,omitempty"`
	CustomerCode string `json:"customerCode,omitempty"`
	OrderId      string `json:"orderId,omitempty"`
	ThrOrderId   string `json:"thrOrderId,omitempty"`

	SenderName     string `json:"senderName,omitempty"`
	SenderAddress  string `json:"senderAddress,omitempty"`
	SenderTel      string `json:"senderTel,omitempty"`
	SenderMobile   string `json:"senderMobile,omitempty"`
	SenderPostcode string `json:"senderPostcode,omitempty"`

	ReceiveName    string `json:"receiveName,omitempty"`
	ReceiveAddress string `json:"receiveAddress,omitempty"`
	Province       string `json:"province,omitempty"`
	City           string `json:"city,omitempty"`
	County         string `json:"county,omitempty"`
	Town           string `json:"town,omitempty"`
	ProvinceId     int    `json:"provinceId,omitempty"`
	CityId         int    `json:"cityId,omitempty"`
	CountyId       int    `json:"countyId,omitempty"`
	TownId         int    `json:"townId,omitempty"`
	SiteType       int    `json:"siteType,omitempty"`
	SiteId         int    `json:"siteId,omitempty"`
	SiteName       string `json:"siteName,omitempty"`
	ReceiveTel     string `json:"receiveTel,omitempty"`
	ReceiveMobile  string `json:"receiveMobile,omitempty"`
	Postcode       string `json:"postcode,omitempty"`

	PackageCount         int     `json:"packageCount,omitempty"`
	Weight               float64 `json:"weight,omitempty"`
	VloumLong            float64 `json:"vloumLong,omitempty"`
	VloumWidth           float64 `json:"vloumWidth,omitempty"`
	VloumHeight          float64 `json:"vloumHeight,omitempty"`
	Vloumn               float64 `json:"vloumn,omitempty"`
	Description          string  `json:"description,omitempty"`
	CollectionValue      int     `json:"collectionValue,omitempty"`
	CollectionMoney      float64 `json:"collectionMoney,omitempty"`
	GuaranteeValue       int     `json:"guaranteeValue,omitempty"`
	GuaranteeValueAmount float64 `json:"guaranteeValueAmount,omitempty"`
	SignReturn           int     `json:"signReturn,omitempty"`
	Aging                int     `json:"aging,omitempty"`
	TransType            int     `json:"transType,omitempty"`
	Remark               string  `json:"remark,omitempty"`
	GoodsType            int     `json:"goodsType,omitempty"`
	OrderType            int     `json:"orderType,omitempty"`
	ShopCode             string  `json:"shopCode,omitempty"`
	OrderSendTime        string  `json:"orderSendTime,omitempty"`

	WarehouseCode       string   `json:"warehouseCode,omitempty"`
	AreaProvId          int      `json:"areaProvId,omitempty"`
	AreaCityId          int      `json:"areaCityId,omitempty"`
	ShipmentStartTime   string   `json:"shipmentStartTime,omitempty"`
	ShipmentEndTime     string   `json:"shipmentEndTime,omitempty"`
	IdNumber            string   `json:"idNumber,omitempty"`
	AddedService        string   `json:"addedService,omitempty"`
	ExtendField1        string   `json:"extendField1,omitempty"`
	ExtendField2        string   `json:"extendField2,omitempty"`
	ExtendField3        string   `json:"extendField3,omitempty"`
	ExtendField4        string   `json:"extendField4,omitempty"`
	ExtendField5        string   `json:"extendField5,omitempty"`
	SenderCompany       string   `json:"senderCompany,omitempty"`
	ReceiveCompany      string   `json:"receiveCompany,omitempty"`
	Goods               string   `json:"goods,omitempty"`
	GoodsCount          int      `json:"goodsCount,omitempty"`
	PromiseTimeType     int      `json:"promiseTimeType,omitempty"`
	Freight             float64  `json:"freight,omitempty"`
	PickUpStartTime     string   `json:"pickUpStartTime,omitempty"`
	PickUpEndTime       string   `json:"pickUpEndTime,omitempty"`
	UnpackingInspection string   `json:"unpackingInspection,omitempty"`
	BoxCode             []string `json:"boxCode,omitempty"`
	FileUrl             string   `json:"fileUrl,omitempty"`
}
type ReqQueryDataDto struct {
	CustomerCode string `json:"customerCode,omitempty"`
	WaybillCode  string `json:"waybillCode,omitempty"`
}
type ReqCancelDataDto struct {
	EndReasonName string `json:"endReasonName,omitempty"`
	EndReason     int    `json:"endReason,omitempty"`
	PickupCode    string `json:"pickupCode,omitempty"`
	Source        string `json:"source,omitempty"`
	CustomerCode  string `json:"customerCode,omitempty"`
}
type ReqCustomerDto struct {
	Url       string `json:"url,omitempty"`
	AppSecret string `json:"appSecret,omitempty"`
}

type PreSortResult struct {
	SiteId                 int    `json:"siteId,omitempty"`
	SiteName               string `json:"siteName,omitempty"`
	Road                   string `json:"road,omitempty"`
	SlideNo                string `json:"slideNo,omitempty"`
	SourceSortCenterId     int    `json:"sourceSortCenterId,omitempty"`
	SourceSortCenterName   string `json:"promiseTimeType,omitempty"`
	SourceCrossCode        string `json:"sourceCrossCode,omitempty"`
	SourceTabletrolleyCode string `json:"sourceTabletrolleyCode,omitempty"`
	TargetSortCenterId     int    `json:"targetSortCenterId,omitempty"`
	TargetSortCenterName   string `json:"targetSortCenterName,omitempty"`
	TargetTabletrolleyCode string `json:"targetTabletrolleyCode,omitempty"`
	Aging                  int    `json:"aging,omitempty"`
	AgingName              string `json:"agingName,omitempty"`
	SiteType               int    `json:"siteType,omitempty"`
	IsHideName             int    `json:"isHideName,omitempty"`
	IsHideContractNumbers  int    `json:"isHideContractNumbers,omitempty"`
}
type TraceDTO struct {
	OpeTitle    string `json:"opeTitle,omitempty"`
	OpeRemark   string `json:"opeRemark,omitempty"`
	OpeName     string `json:"opeName,omitempty"`
	OpeTime     string `json:"opeTime,omitempty"`
	WaybillCode string `json:"waybillCode,omitempty"`
	Courier     string `json:"courier,omitempty"`
	CourierTel  string `json:"courierTel,omitempty"`
}
