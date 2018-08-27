package wechat

type MessageType string

var (
	MsgTypeText     MessageType = "text"
	MsgTypeImage    MessageType = "image"
	MsgTypeVoice    MessageType = "voice"
	MsgTypeVideo    MessageType = "video"
	MsgTypeFile     MessageType = "file"
	MsgTypeTextcard MessageType = "textcard"
	MsgTypeNews     MessageType = "news"
	MsgTypeMPNews   MessageType = "mpnews"
)

type Message struct {
	ToUser  string   `json:"touser"`
	ToParty string   `json:"toparty"`
	ToTag   string   `json:"totag"`
	MsgType string   `json:"msgtype"`
	AgentId int      `json:"agentid"`
	Safe    int      `json:"safe"`
	Text    MsgText  `json:"text, omitempty"`
	Image   MsgImage `json:"image, omitempty"`
	//todo... other msg type
}

type MsgText struct {
	Content string `json:"content"`
}

type MsgImage struct {
	MediaId string `json:"media_id"`
}

type MsgRet struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

func (w *WechatWork) SendText(user, party, tag, txt string) (MsgRet, error) {

	msg := Message{
		ToUser:  user,
		ToParty: party,
		ToTag:   tag,
		MsgType: "text",
		Text: MsgText{
			Content: txt,
		},
	}

	return w.Send(msg)
}

//https://work.weixin.qq.com/api/doc#10167
func (w *WechatWork) Send(msg Message) (ret MsgRet, err error) {

	url1, err := w.msgSendAddr()
	if err != nil {
		return
	}

	err = w.Client.CallWithJson(nil, &ret, "POST", url1, msg)
	return
}

func (w *WechatWork) msgSendAddr() (addr string, err error) {
	token, err := w.GetAccessToken()
	if err != nil {
		return
	}

	addr = ApiHost + "/cgi-bin/message/send?access_token=" + token.AccessToken
	return
}
