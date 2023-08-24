package cla

type Goarg struct {
	Title       string
	Usage       string
	Examples    []string
	Description string
	Sections     []*Section
	Options     []Option
	Mods        map[string]*Goarg
	ModeName    string
	Main        bool
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
	ModeName string
}

type Section struct {
	Value string
	index int
}
