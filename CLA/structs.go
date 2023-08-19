package cla

type Goarg struct {
	Title    string
	Usage    string
	Examples []string
	Options  []Option
	Mode     map[string]*Goarg
	Main     bool
}

type Option struct {
	PlaceHolder []string
	Active      bool
	Usage       string
	Error       []string
}

type Input struct {
	Argument string
	Value    string
	Error    []string
}
