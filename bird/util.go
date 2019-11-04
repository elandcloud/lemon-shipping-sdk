package bird

import (
	"encoding/base64"
	"net/url"

	"github.com/relax-space/go-kit/sign"
)

func SignBird(text string) (string, error) {
	md5Data, err := sign.GetMD5Hash(text, true)
	if err != nil {
		return "", err
	}
	da := base64.URLEncoding.EncodeToString([]byte(md5Data))
	encodeData := url.QueryEscape(da)
	return encodeData, nil
}
