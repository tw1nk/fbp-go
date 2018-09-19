package protocol

import "encoding/json"

// InGetRuntime Request the information about the runtime. When receiving this message
//      the runtime should response with a runtime message.
//      If the runtime is currently running a graph and it is
//      able to speak the full Runtime protocol, it should follow up
//      with a ports message.
type InGetRuntime struct {
}

// InPacket Runtimes that can be used as remote subgraphs (i.e. ones that have
//      reported supporting the protocol:runtime capability) need to be able to
//      receive and transmit information packets at their exposed ports.
//      These packets can be send from the client to the runtimes input ports, or from
//      runtimes output ports to the client.
type InPacket struct {
	Port    string          `json:"port"`
	Event   string          `json:"event"`
	Type    string          `json:"type"`
	Schema  string          `json:"schema"`
	Payload json.RawMessage `json:"payload"`
	Graph   string          `json:"graph"`
}

// OutError Error response to a command on runtime protocol
type OutError struct {
	Message string `json:"message"`
}

// OutPorts Message sent by the runtime to signal its available ports. The runtime
//      is responsible for sending the up-to-date list of available ports back to
//      client whenever it changes.
type OutPorts struct {
	Graph    string           `json:"graph"`
	InPorts  []PortDefinition `json:"inPorts"`
	OutPorts []PortDefinition `json:"outPorts"`
}

type OutRuntime struct {
	ID                string   `json:"id"`
	Label             string   `json:"label"`
	Version           string   `json:"version"`
	AllCapabilities   []string `json:"allCapabilities"`
	Capabilities      []string `json:"capabilities"`
	Graph             string   `json:"graph"`
	Type              string   `json:"type"`
	Namespace         string   `json:"namespace"`
	Repository        string   `json:"repository"`
	RepositoryVersion string   `json:"repositoryVersion"`
}

type OutPacketSent struct {
	Port    string          `json:"port"`
	Event   string          `json:"event"`
	Type    string          `json:"type"`
	Schema  string          `json:"schema"`
	Payload json.RawMessage `json:"payload"`
	Graph   string          `json:"graph"`
}

type OutPacket struct {
	Port    string          `json:"port"`
	Event   string          `json:"event"`
	Type    string          `json:"type"`
	Schema  string          `json:"schema"`
	Payload json.RawMessage `json:"payload"`
	Graph   string          `json:"graph"`
}
