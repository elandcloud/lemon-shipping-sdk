package jd

import (
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
