package types

// Package data
type Package struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
	Depend  []int  `json:"Depend"`
}
