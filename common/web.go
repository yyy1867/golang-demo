package common

type H map[string]interface{}

type Result struct {
	Success bool
	Msg     string
	Data    interface{}
}

type BsmResult struct {
	Success  bool
	Message  string
	Data     interface{}
	Status   string
	Solution string
}
