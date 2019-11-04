package best

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

func toBase64(dto interface{}) (data string, err error) {
	da, err := json.Marshal(dto)
	if err != nil {
		return
	}
	encodeData := url.QueryEscape(string(da))
	data = base64.StdEncoding.EncodeToString([]byte(encodeData))
	return
}
