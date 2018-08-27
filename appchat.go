package wechatwork

//https://work.weixin.qq.com/api/doc#10167
//https://work.weixin.qq.com/api/doc#13288
type AppChatMsg struct {
	ChatId  string      `json:"chatid"`
	MsgType MessageType `json:"msgtype"`
	Safe    int         `json:"safe"`
	Text    MsgText     `json:"text, omitempty"`
	Image   MsgImage    `json:"image, omitempty"`
}

type AppChatRet struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	ChatId  string `json:"chatid"`
}

func (w *WechatWork) AppChatText(msg AppChatMsg) (ret AppChatRet, err error) {
	url1, err := w.appchatSendAddr()
	if err != nil {
		return
	}

	err = w.Client.CallWithJson(nil, &ret, "POST", url1, msg)
	return
}

func (w *WechatWork) appchatSendAddr() (addr string, err error) {
	token, err := w.GetAccessToken()
	if err != nil {
		return
	}

	addr = ApiHost + "/cgi-bin/appchat/send?access_token=" + token.AccessToken
	return
}

func (w *WechatWork) NewChat(name, owner, chatId string, users []string) (ret AppChatRet, err error) {

	param := map[string]interface{}{
		"name":     name,
		"owner":    owner,
		"chatid":   chatId,
		"userlist": users,
	}
	url1, err := w.appchatCreateAddr()
	if err != nil {
		return
	}

	err = w.Client.CallWithJson(nil, &ret, "POST", url1, param)
	return
}

func (w *WechatWork) appchatCreateAddr() (addr string, err error) {
	token, err := w.GetAccessToken()
	if err != nil {
		return
	}

	addr = ApiHost + "/cgi-bin/appchat/create?access_token=" + token.AccessToken
	return
}
