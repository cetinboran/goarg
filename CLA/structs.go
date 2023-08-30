package cla

type Goarg struct {
	Usage    *Usage
	Options  []Option
	Mods     map[string]*Goarg
	ModeName string
	Errors   map[string]*OptionError
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

// Bu value'yi string yerine interface yapabilirsin sonra şuan çok hata çıkar
type Input struct {
	Argument string
	Value    string
	ModeName string
}

type OptionError struct {
	Values []string
}
