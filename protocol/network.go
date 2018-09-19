package protocol

import (
	"encoding/json"
	"time"
)

type OutNetworkStopped struct {
	Time    time.Time `json:"time"`
	Uptime  uint32    `json:"uptime"`
	Graph   string    `json:"graph"`
	Running bool      `json:"running"`
	Started bool      `json:"started"`
	Debug   bool      `json:"debug"`
}

type OutNetworkStarted struct {
	Time    time.Time `json:"time"`
	Graph   string    `json:"graph"`
	Running bool      `json:"running"`
	Started bool      `json:"started"`
	Debug   bool      `json:"debug"`
}

type OutNetworkStatus struct {
	Uptime  uint32 `json:"uptime"`
	Graph   string `json:"graph"`
	Running bool   `json:"running"`
	Started bool   `json:"started"`
	Debug   bool   `json:"debug"`
}

type OutNetworkOutput struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	URL     string `json:"url"`
}

type OutNetworkError struct {
	Message string `json:"message"`
	Stack   string `json:"stack"`
	Graph   string `json:"graph"`
}

type OutNetworkProcessError struct {
	ID    string `json:"id"`
	Error string `json:"error"`
	Graph string `json:"graph"`
}

type OutNetworkIcon struct {
	ID    string `json:"id"`
	Icon  string `json:"icon"`
	Graph string `json:"graph"`
}

type OutNetworkConnect struct {
	ID       string   `json:"id"`
	Source   string   `json:"src"`
	Target   string   `json:"tgt"`
	Graph    string   `json:"graph"`
	Subgraph []string `json:"subgraph"`
}

type OutNetworkBeginGroup struct {
	ID       string   `json:"id"`
	Source   string   `json:"src"`
	Target   string   `json:"tgt"`
	Graph    string   `json:"graph"`
	Subgraph []string `json:"subgraph"`
	Group    string   `json:"group"`
}

type OutNetworkData struct {
	ID       string          `json:"id"`
	Source   string          `json:"src"`
	Target   string          `json:"tgt"`
	Graph    string          `json:"graph"`
	Subgraph []string        `json:"subgraph"`
	Data     json.RawMessage `json:"data"`
	Type     string          `json:"type"`
	Schema   string          `json:"schema"`
}

type OutNetworkEndGroup struct {
	ID       string   `json:"id"`
	Source   string   `json:"src"`
	Target   string   `json:"tgt"`
	Graph    string   `json:"graph"`
	Subgraph []string `json:"subgraph"`
	Group    string   `json:"group"`
}

type OutNetworkDisconnect struct {
	ID       string   `json:"id"`
	Source   string   `json:"src"`
	Target   string   `json:"tgt"`
	Graph    string   `json:"graph"`
	Subgraph []string `json:"subgraph"`
}

type NetworkEdge struct {
	Source Port `json:"src"`
	Target Port `json:"tgt"`
}

type InOutNetworkEdges struct {
	Edges []NetworkEdge `json:"edges"`
}

type InNetworkStart struct {
	Graph string `json:"graph"`
}

type InNetworkGetStatus struct {
	Graph string `json:"graph"`
}

type InNetworkStop struct {
	Graph string `json:"graph"`
}

type InNetworkPersist struct{}

type InNetworkDebug struct {
	Enable bool   `json:"enable"`
	Graph  string `json:"graph"`
}
