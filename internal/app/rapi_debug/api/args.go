package api

import (
	"github.com/maldan/go-rapi/rapi_core"
	"time"
)

type ArgsEmpty struct {
}
type ArgsId struct {
	Id string `json:"id"`
}
type ArgsAB struct {
	A int `json:"a"`
	B int `json:"b"`
}
type ArgsSuper struct {
	Sas  string         `json:"sas"`
	T    time.Time      `json:"t"`
	File rapi_core.File `json:"file"`
}
