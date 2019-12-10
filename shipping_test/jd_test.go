package shipping_test

import (
	"lemon-shipping-sdk/jd"
	"net/http"
	"testing"

	"github.com/pangpanglabs/goutils/test"
)

func Test_JD_ECreate(t *testing.T) {
	expDto := jd.RespECreateDto{
		Response: jd.RespECreateResponse{
			Result: jd.RespECreateResult{
				ResultCode:      104,
				ResultMessage:   "重复运单",
				OrderId:         "",
				DeliveryId:      "JDV000348594219",
				PromiseTimeType: 1,
				TransType:       1,
			},
		},
	}
	custDto := jd.ReqCustomerDto{
		Url:       jdDto.Url,
		AppSecret: jdDto.AppSecret,
	}
	reqDto := jd.ReqECreateDto{
		ReqBase: jd.ReqBase{
			AccessToken: jdDto.AccessToken,
			AppKey:      jdDto.AppKey,
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqECreateDataDto{
			SalePlat:       "0030001",
			CustomerCode:   jdDto.CustomerCode,
			OrderId:        "1",
			SenderName:     "张三",
			SenderAddress:  "北京市朝阳区酒仙桥路",
			SenderTel:      "13051488887",
			ReceiveName:    "李四",
			ReceiveAddress: "上海市宣武区向阳路",
			ReceiveTel:     "13051488888",
			PackageCount:   1,
			Weight:         1.0,
			Vloumn:         1.0,
			Remark:         "仅供测试使用，无需上门揽收",
		},
	}

	statusCode, respDto, err := jd.ECreate(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto.Response.Result.ResultCode, respDto.Response.Result.ResultCode)
	test.Equals(t, expDto.Response.Result.ResultMessage, respDto.Response.Result.ResultMessage)
	test.Equals(t, expDto.Response.Result.DeliveryId, respDto.Response.Result.DeliveryId)
	test.Equals(t, expDto.Response.Result.PromiseTimeType, respDto.Response.Result.PromiseTimeType)
	test.Equals(t, expDto.Response.Result.TransType, respDto.Response.Result.TransType)

}
func Test_JD_Create(t *testing.T) {
	expDto := jd.RespCreateDto{
		Response: jd.RespCreateResponse{
			Result: jd.RespCreateResult{
				Code:       0,
				Message:    "",
				OrderId:    "",
				DeliveryId: "",
			},
		},
		ErrorResponse: jd.ErrorResponse{
			Response: jd.Response{
				Code:   "61",
				ZhDesc: "deliveryId-参数不能为空，请填写参数值!（解决方案参考：http://open.jd.com/home/home#/doc/common?listId=533）",
				EnDesc: "parameter null value is not valid ,please  refer to  the help document and confirm(Solution reference:http://open.jd.com/home/home#/doc/common?listId=533)",
			},
		},
	}

	custDto := jd.ReqCustomerDto{
		Url:       jdDto.Url,
		AppSecret: jdDto.AppSecret,
	}
	reqDto := jd.ReqCreateDto{
		ReqBase: jd.ReqBase{
			AccessToken: jdDto.AccessToken,
			AppKey:      jdDto.AppKey,
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqCreateDataDto{
			DeliveryId:     "",
			SalePlat:       "0030001",
			CustomerCode:   jdDto.CustomerCode,
			OrderId:        "1",
			SenderName:     "张三",
			SenderAddress:  "北京市朝阳区酒仙桥路",
			SenderTel:      "13051488887",
			ReceiveName:    "李四",
			ReceiveAddress: "上海市宣武区向阳路",
			ReceiveTel:     "13051488888",
			PackageCount:   1,
			Weight:         1.0,
			Vloumn:         1.0,
			Remark:         "仅供测试使用，无需上门揽收",
		},
	}

	statusCode, respDto, err := jd.Create(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto.ErrorResponse.Response.Code, respDto.ErrorResponse.Response.Code)
	test.Equals(t, expDto.ErrorResponse.Response.ZhDesc, respDto.ErrorResponse.Response.ZhDesc)

}
func Test_JD_Cancel(t *testing.T) {
	expDto := jd.RespCancelDto{
		Response: jd.RespCancelResponse{
			Result: jd.RespCancelResult{
				StatusMessage: "",
				StatusCode:    0,
			},
		},
		ErrorResponse: jd.ErrorResponse{
			Response: jd.Response{
				Code:   "88",
				ZhDesc: "该appKey=D4841B62CA4A82D6244023D7FEDB6435无权调用method=jingdong.ldop.pickup.cancel（解决方案参考：http://open.jd.com/home/home#/doc/common?listId=533）",
				EnDesc: "The appKey has no right to call this interface. Please go to dev.jd.com for interface application(Solution reference:http://open.jd.com/home/home#/doc/common?listId=533)",
			},
		},
	}
	custDto := jd.ReqCustomerDto{
		Url:       jdDto.Url,
		AppSecret: jdDto.AppSecret,
	}
	reqDto := jd.ReqCancelDto{
		ReqBase: jd.ReqBase{
			AccessToken: jdDto.AccessToken,
			AppKey:      jdDto.AppKey,
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqCancelDataDto{
			EndReasonName: "客户取消服务单，终止取件",
			EndReason:     19,
			PickupCode:    "",
			Source:        "",
			CustomerCode:  jdDto.CustomerCode,
		},
	}

	statusCode, respDto, err := jd.Cancel(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto.ErrorResponse.Response.Code, respDto.ErrorResponse.Response.Code)
	test.Equals(t, expDto.ErrorResponse.Response.ZhDesc, respDto.ErrorResponse.Response.ZhDesc)
}
func Test_JD_CancelWayBill(t *testing.T) {
	expDto := jd.RespCancelWayBillDto{
		Response: jd.RespCancelWayBillResponse{
			Result: jd.RespCancelResult{
				StatusMessage: "已取消订单不能再次取消",
				StatusCode:    1,
			},
		},
		ErrorResponse: jd.ErrorResponse{
			Response: jd.Response{
				Code:   "",
				ZhDesc: "",
				EnDesc: "",
			},
		},
	}
	custDto := jd.ReqCustomerDto{
		Url:       jdDto.Url,
		AppSecret: jdDto.AppSecret,
	}
	reqDto := jd.ReqCancelWayBillDto{
		ReqBase: jd.ReqBase{
			AccessToken: jdDto.AccessToken,
			AppKey:      jdDto.AppKey,
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqCancelWayBillDataDto{
			CancelReason: "客户取消服务单，终止取件",
			OperatorName: "19",
			WaybillCode:  "JDV000340894610",
			Source:       "",
			CustomerCode: jdDto.CustomerCode,
			UserPin:      "",
		},
	}

	statusCode, respDto, err := jd.CancelWayBill(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto.Response.Result.StatusCode, respDto.Response.Result.StatusCode)
	test.Equals(t, expDto.Response.Result.StatusMessage, respDto.Response.Result.StatusMessage)
}
func Test_JD_Query(t *testing.T) {
	expDto := jd.RespQueryDto{
		Response: jd.RespQueryResponse{
			Result: jd.RespQueryResult{
				Code:     100,
				Messsage: "成功",
				Data: []jd.TraceDTO{
					jd.TraceDTO{
						OpeTitle:    "下单取消",
						OpeRemark:   "因用户发起取消原因，运单取消",
						OpeName:     "京东快递",
						WaybillCode: "JDV000340894610",
					},
				},
			},
		},
	}
	custDto := jd.ReqCustomerDto{
		Url:       jdDto.Url,
		AppSecret: jdDto.AppSecret,
	}
	reqDto := jd.ReqQueryDto{
		ReqBase: jd.ReqBase{
			AccessToken: jdDto.AccessToken,
			AppKey:      jdDto.AppKey,
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqQueryDataDto{
			CustomerCode: jdDto.CustomerCode,
			WaybillCode:  "JDV000340894610",
		},
	}
	statusCode, respDto, err := jd.Query(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto.Response.Result.Code, respDto.Response.Result.Code)
	test.Equals(t, expDto.Response.Result.Messsage, respDto.Response.Result.Messsage)
}
