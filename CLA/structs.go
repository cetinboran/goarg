package cla

type GoargErrors map[string]*OptionError

type Goarg struct {
	usage    *Usage
	options  []*Option
	mods     map[string]*Goarg
	modeName string
	errors   GoargErrors
}

type Usage struct {
	Title       string
	Description string
	Examples    []string
	Message     string
	mainTitle   string
}

type Option struct {
	PlaceHolder []string
	Active      bool
	Usage       string
	Global      bool
	title       string // this is for usage
}

// Bu value'yi string yerine interface yapabilirsin sonra şuan çok hata çıkar
type Input struct {
	Argument string
	Value    string
	ModeName string
}

type OptionError struct {
	Values []string
}
