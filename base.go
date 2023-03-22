package bprompt


type Base interface{
	Len() (int)
	Ask() (*Base, error)
}
