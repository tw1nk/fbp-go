package protocol

type InOutTraceStart struct {
	Graph      string `json:"graph"`
	BufferSize uint64 `json:"buffersize"`
}

type InOutTraceStop struct {
	Graph string `json:"graph"`
}

type InOutTraceClear struct {
	Graph string `json:"graph"`
}

type OutTraceDump struct {
	Graph     string `json:"graph"`
	Type      string `json:"type"`
	FlowTrace string `json:"flowtrace"`
}

type InTraceDump struct {
	Graph string `json:"graph"`
	Type  string `json:"type"`
}
