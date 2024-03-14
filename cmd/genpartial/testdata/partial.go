// !DO NOT EDIT! Generated by genpartial
//
//go:build !generate
//+gotohelm:ignore=true
package main

type (
	PartialMapGeneric[T any] map[string]T
)

type PartialExampleStruct struct {
	A1 PartialMapGeneric[int]                  `json:",omitempty"`
	A2 PartialMapGeneric[PartialNestedStruct]  `json:",omitempty"`
	A3 PartialMapGeneric[*PartialNestedStruct] `json:",omitempty"`
	A4 PartialMapGeneric[*PartialNestedStruct] `json:",omitempty"`
	A5 PartialMapGeneric[IntAlias]             `json:",omitempty"`

	B1 *int `json:",omitempty"`
	B2 *int `json:",omitempty"`

	C1 struct {
		Any any  `json:",omitempty"`
		Int *int `json:",omitempty"`
	} `json:",omitempty"`
	C2 *struct{} `json:",omitempty"`

	D1 *PartialNestedStruct `json:",omitempty"`
	D2 *PartialNestedStruct `json:",omitempty"`

	E1 []any  `json:",omitempty"`
	E2 []int  `json:",omitempty"`
	E3 []*int `json:",omitempty"`

	F1 []*int   `json:"L,omitempty"`
	F2 *string  `yaml:"M"json:",omitempty"`
	F3 IntAlias `json:",omitempty"`
}

type PartialNestedStruct struct {
	Map map[string]string `json:",omitempty"`
}