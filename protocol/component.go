package protocol

type OutComponentComponent struct {
	Name        string
	Description string
	Icon        string
	SubGraph    bool
	InPorts     []PortDefinition
	OutPorts    []PortDefinition
}

type OutComponentReady int

type OutComponentSource struct {
	Name     string
	Language string
	Library  string
	Code     string
	Tests    string
}

type InComponentList struct{}

type InComponentGetSource struct {
	Name string
}

type InComponentSource struct {
	Name     string
	Language string
	Code     string
	Library  string
	Tests    string
}
