package cla

type Goarg struct {
	Usage    *Usage
	Options  []Option
	Mods     map[string]*Goarg
	ModeName string
}

type Usage struct {
	Title       string
	Description string
	Examples    []string
	Message     string
}

type Option struct {
	PlaceHolder []string
	Active      bool
	Usage       string
	Global      bool
}

type Input struct {
	Argument string
	ValueMap map[string]string
	ModeName string
}
