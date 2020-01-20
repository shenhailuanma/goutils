package slack

type SlackObj struct {
	Token string
}

var defaultObj SlackObj

func InitDefault(token string) {
	defaultObj = SlackObj{}
	defaultObj.Token = token
}
