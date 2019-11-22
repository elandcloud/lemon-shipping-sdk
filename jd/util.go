package jd

import (
	"github.com/mitchellh/mapstructure"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

func SignJD(mapData map[string]interface{}, appSecret string) (string, error) {
	strParams := base.JoinMap(mapData, "", "")
	signParam := appSecret + strParams + appSecret

	md5Data, err := sign.GetMD5Hash(signParam, false)
	if err != nil {
		return "", err
	}
	return md5Data, nil
}

// Decode takes a map and uses reflection to convert it into the
// given Go native structure. val must be a pointer to a struct.
func Decode(m interface{}, rawVal interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   rawVal,
		TagName:  "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(m)
}
