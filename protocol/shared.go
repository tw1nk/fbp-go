package protocol

import "encoding/json"

type InputMessage interface {
	Parse() error
}

type OutputMessage interface {
	Marshall() error
}

type InMessage struct {
	Protocol  string          `json:"protocol"`
	Command   string          `json:"command"`
	Secret    string          `json:"secret"`
	RequestID string          `json:"requestId"`
	Payload   json.RawMessage `json:"payload"`
}

type OutMessage struct {
	Protocol   string      `json:"protocol"`
	Command    string      `json:"command"`
	ResponseTo string      `json:"responseTo"`
	Payload    interface{} `json:"payload"`
}

type PortDefinition struct {
	ID          string        `json:"id"`
	Type        string        `json:"type"`
	Schema      string        `json:"schema"`
	Required    bool          `json:"required"`
	Addressable bool          `json:"addressable"`
	Description string        `json:"description"`
	Values      []interface{} `json:"values"`
	Default     interface{}   `json:"default"`
}

type Port struct {
	Node  string      `json:"node"`
	Port  string      `json:"port"`
	Index interface{} `json:"index"` // string or number :(
}
