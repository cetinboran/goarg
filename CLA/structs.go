package cla

type Goarg struct {
	Title    string
	Usage    string
	Examples []string
	Options  []Option
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
