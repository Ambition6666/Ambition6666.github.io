package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WxLogin 微信用户授权
func (m User) WxLogin(jscode string) (session WxSession, err error) {
	client := &http.Client{}

	//生成要访问的url
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", "wxa8a7dd7461e5d24d", "5e5274c59d49bb451946026c20fd51f9", jscode)

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, e := client.Do(reqest)
	if e != nil {
		panic(e)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	jsonStr := string(body)
	//解析json
	if err := json.Unmarshal(body, &session); err != nil {
		session.SessionKey = jsonStr
		return session, err
	}

	return session, err
}
