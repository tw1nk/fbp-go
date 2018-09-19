package protocol

type InOutGraphClear struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Library     string `json:"library"`
	Main        bool   `json:"main"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type InOutGraphAddNode struct {
	ID        string                 `json:"id"`
	Component string                 `json:"component"`
	Graph     string                 `json:"graph"`
	Metadata  map[string]interface{} `json:"metadata"` // yuck.
}

type InOutGraphRemoveNode struct {
	ID    string `json:"id"`
	Graph string `json:"graph"`
}

type InOutGraphRenameNode struct {
	Graph string `json:"graph"`
	From  string `json:"from"`
	To    string `json:"to"`
}

type InOutGraphChangeNode struct {
	ID       string                 `json:"id"`
	Graph    string                 `json:"graph"`
	Metadata map[string]interface{} `json:"metadata"` // yuck.
}

type InOutGraphAddEdge struct {
	Source   string                 `json:"src"`
	Target   string                 `json:"tgt"`
	Graph    string                 `json:"graph"`
	Metadata map[string]interface{} `json:"metadata"`
}

type InOutGraphRemoveEdge struct {
	Source string `json:"src"`
	Target string `json:"tgt"`
	Graph  string `json:"graph"`
}

type InOutGraphChangeEdge struct {
	Source   string                 `json:"src"`
	Target   string                 `json:"tgt"`
	Graph    string                 `json:"graph"`
	Metadata map[string]interface{} `json:"metadata"`
}

type InOutGraphAddInitial struct {
	Source   string                 `json:"src"`
	Target   string                 `json:"tgt"`
	Graph    string                 `json:"graph"`
	Metadata map[string]interface{} `json:"metadata"`
}

type InOutGraphRemoveInitial struct {
	Source string `json:"src"`
	Target string `json:"tgt"`
	Graph  string `json:"graph"`
}

type InOutGraphAddInPort struct {
	Public   string                 `json:"public"`
	Node     string                 `json:"node"`
	Port     string                 `json:"port"`
	Metadata map[string]interface{} `json:"metadata"`
	Graph    string                 `json:"graph"`
}

type InOutGraphRemoveInPort struct {
	Public string `json:"public"`
	Graph  string `json:"graph"`
}

type InOutGraphRenameInPort struct {
	Graph string
	From  string `json:"from"`
	To    string `json:"to"`
}

type InOutGraphAddOutPort struct {
	Public   string                 `json:"public"`
	Node     string                 `json:"node"`
	Port     string                 `json:"port"`
	Metadata map[string]interface{} `json:"metadata"`
	Graph    string                 `json:"graph"`
}

type InOutGraphRemoveOutPort struct {
	Public string `json:"public"`
	Graph  string `json:"graph"`
}

type InOutGraphRenameOutPort struct {
	Graph string `json:"graph"`
	From  string `json:"from"`
	To    string `json:"to"`
}

type InGraphAddGroup struct {
	Name     string                 `json:"name"`
	Nodes    []string               `json:"nodes"`
	Metadata map[string]interface{} `json:"metadata"`
	Graph    string                 `json:"graph"`
}

type InOutGraphRemoveGroup struct {
	Name  string `json:"name"`
	Graph string `json:"graph"`
}

type InOutGraphRenameGroup struct {
	Graph string `json:"graph"`
	From  string `json:"from"`
	To    string `json:"to"`
}

type InOutGraphChangeGroup struct {
	Name     string                 `json:"name"`
	Metadata map[string]interface{} `json:"metadata"`
	Graph    string                 `json:"graph"`
}
