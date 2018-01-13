package wordpress

type BaseCall interface {
	GetMethord() string
	GetArgs(user string, pwd string) interface{}
}
