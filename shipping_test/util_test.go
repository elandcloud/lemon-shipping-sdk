package shipping_test

import "os"

var bestDto = struct {
	Url string
	Pid string
	Key string
}{
	Url: "http://183.129.172.49/o2o/api/process",
	Pid: os.Getenv("BEST_PID_DEV"),
	Key: os.Getenv("BEST_KEY_DEV"),
}

var birdDto = struct {
	Url         string
	EBusinessId string
	Key         string
}{
	Url:         "http://sandboxapi.kdniao.com:8080/kdniaosandbox/gateway/exterfaceInvoke.json",
	EBusinessId: os.Getenv("BIRD_EBusinessId"),
	Key:         os.Getenv("BIRD_APIKEY"),
}

var sfDto = struct {
	Url       string
	Linkcode  string
	CheckWord string
}{
	Url:       "http://bsp-ois.sit.sf-express.com:9080/bsp-ois/sfexpressService",
	Linkcode:  os.Getenv("SF_Linkcode"),
	CheckWord: os.Getenv("SF_CHECKWORD"),
}

var springDto = struct {
	Url  string
	Code string
	Pwd  string
}{
	Url:  "http://sync.spring56.com/api/interface.php",
	Code: os.Getenv("SPRING_CODE"),
	Pwd:  os.Getenv("SPRING_PWD"),
}

var jdDto = struct {
	Url          string
	AppKey       string
	AppSecret    string
	AccessToken  string
	CustomerCode string
}{
	Url:          "https://api.jd.com/routerjson",
	AppKey:       os.Getenv("JD_APPKEY"),
	AppSecret:    os.Getenv("JD_APPSECRET"),
	AccessToken:  os.Getenv("JD_ACCESSTOKEN"),
	CustomerCode: os.Getenv("JD_CUSTOMERCODE"),
}
