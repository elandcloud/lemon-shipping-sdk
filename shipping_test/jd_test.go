package shipping_test

import (
	"encoding/json"
	"fmt"
	"lemon-shipping-sdk/jd"
	"net/http"
	"reflect"
	"testing"

	"github.com/pangpanglabs/goutils/test"
	"github.com/relax-space/go-kit/base"
)

func Test_JD_Create(t *testing.T) {
	expDto := jd.RespCreateDto{
		Response: jd.RespCreateResponse{
			Result: jd.RespCreateResult{
				ResultCode:      0,
				ResultMessage:   "",
				OrderId:         "",
				DeliveryId:      "",
				PromiseTimeType: 1,
				PreSortResult:   &jd.PreSortResult{},
				TransType:       1,
			},
		},
	}
	custDto := jd.ReqCustomerDto{
		Url:       "",
		AppSecret: "",
	}
	reqDto := jd.ReqCreateDto{
		ReqBase: jd.ReqBase{
			Method:      "",
			AccessToken: "",
			AppKey:      "",
			Sign:        "",
			Timestamp:   "",
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqCreateDataDto{
			SalePlat:       "",
			CustomerCode:   "",
			OrderId:        "",
			SenderName:     "",
			SenderAddress:  "",
			ReceiveName:    "",
			ReceiveAddress: "",
			PackageCount:   0,
			Weight:         1.0,
			Vloumn:         1.0,
		},
	}

	statusCode, respDto, err := jd.Create(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)

}
func Test_JD_Cancel(t *testing.T) {
	expDto := jd.RespCancelDto{
		Response: jd.RespCancelResponse{
			Result: jd.RespCancelResult{
				StatusMessage: "",
				StatusCode:    0,
			},
		},
	}
	custDto := jd.ReqCustomerDto{
		Url:       "",
		AppSecret: "",
	}
	reqDto := jd.ReqCancelDto{
		ReqBase: jd.ReqBase{
			Method:      "",
			AccessToken: "",
			AppKey:      "",
			Sign:        "",
			Timestamp:   "",
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqCancelDataDto{
			EndReasonName: "",
			EndReason:     19,
			PickupCode:    "",
			Source:        "",
			CustomerCode:  "",
		},
	}

	statusCode, respDto, err := jd.Cancel(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
func Test_JD_Query(t *testing.T) {
	expDto := jd.RespQueryDto{
		Response: jd.RespQueryResponse{
			Result: jd.RespQueryResult{
				Code:     0,
				Messsage: "",
				Data:     []jd.TraceDTO{},
			},
		},
	}
	custDto := jd.ReqCustomerDto{
		Url:       "",
		AppSecret: "",
	}
	reqDto := jd.ReqQueryDto{
		ReqBase: jd.ReqBase{
			Method:      "",
			AccessToken: "",
			AppKey:      "",
			Sign:        "",
			Timestamp:   "",
			Format:      "json",
			V:           "2.0",
		},
		RequestData: jd.ReqQueryDataDto{
			CustomerCode: "",
			WaybillCode:  "",
		},
	}
	statusCode, respDto, err := jd.Query(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}

func Test_JD_Sign(t *testing.T) {
	map2 := getMapFromJson()
	// strAnd := base.JoinMap(map2, "", "")
	fmt.Println("=======", base.JoinMapObjectEncode(map2))

	sign, _ := jd.SignJD(map2, "110")

	fmt.Println("=======", sign, base.JoinMapObjectEncode(map2))
	// fmt.Println(" map的长度：", len(map2))
	// //1.定义一个slice
	// s1 := make([]string, 0, len(map2))
	// map3 := make(map[string]interface{})
	// //2.遍历map获取key-->s1中
	// for key := range map2 {
	// 	s1 = append(s1, key)
	// }
	// //3.给s1进行排序
	// //sort.Ints(s1) //使用sort包下的方法直接排序，不用自己写冒泡了。
	// sort.Strings(s1)
	// //4. 遍历s1，map
	// for _, k := range s1 { // 先下标，再数值
	// 	map3[k] = map2[k]
	// }
	// b, berror := json.Marshal(map3)
	// if berror != nil {
	// 	fmt.Print("berror:", berror)
	// }
	// fmt.Println(string(b))
}

func getMapFromJson() map[string]interface{} {
	var jsonBody = []byte(`{"snifeni":"王者农药","undinfi":["name","消消乐"],"akehi":"绝地求生","yunin":"传奇霸业","zhondfi":"连连看"}`)
	var map4 = make(map[string]interface{})
	err := json.Unmarshal(jsonBody, &map4)
	if err != nil {
		println("unJsonerr:", err)
	} else {
		// for k, v := range map4 {
		// 	print("map4 k and Value :", k, v)
		// }
	}

	reqCreateDto := jd.ReqCreateDto{
		ReqBase: jd.ReqBase{Method: "reciver"},
		RequestData: jd.ReqCreateDataDto{
			CustomerCode: "123456",
			OrderId:      "weee",
			City:         "sksksk",
		},
	}
	jsonBytes, err := json.Marshal(reqCreateDto)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("string(jsonBytes)======", string(jsonBytes))
	var map5 = make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &map5)
	if err != nil {
		println("unJsonerr:", err)
	}

	map_1 := StructToMapDemo(reqCreateDto)
	fmt.Println("map_1===========", map_1)
	fmt.Println("map_1===========", base.JoinMap(map_1, "", ""))

	return map5
}

func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
