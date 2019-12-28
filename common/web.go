package common

type H map[string]interface{}

type Result struct {
	Success bool
	Msg     string
	Data    interface{}
}
