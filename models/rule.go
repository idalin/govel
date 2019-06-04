package models

type SourceRule struct {
	Mode string `json:"mode"`
	Rule string `json:"rule"`
	Js   string `json:"js"`
}
