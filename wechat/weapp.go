package wechat

import (
	"encoding/json"
	"fmt"

	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/util"
)

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError

	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

//Code2Session 登录凭证校验。
func Code2Session(jsCode string) (result ResCode2Session, err error) {
	weappKey := common.GetWeappKey()
	urlStr := fmt.Sprintf(code2SessionURL, weappKey["appid"], weappKey["appsecret"], jsCode)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
