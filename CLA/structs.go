package cla

type Goarg struct {
	Title       string
	Usage       string
	Examples    []string
	Description string
	Options     []Option
	Mods        map[string]*Goarg
	ModeName    string
	Main        bool
}

type Option struct {
	PlaceHolder []string
	Active      bool
	Usage       string
	Global      bool
}

type Input struct {
	Argument string
	Value    interface{} // Bu string ti ordan sıkıntı çıkabilir.
	ModeName string
}
