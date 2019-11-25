package gogithub

type Gogithub struct {
	Username string
	Password string
}

var defaultObj Gogithub

func InitDefault(username, password string) {
	defaultObj.Username = username
	defaultObj.Password = password
}

