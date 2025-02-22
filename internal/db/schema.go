package db

import _ "embed"

//go:embed schema.sql
var Schema string

type Size = int64

const (
	SIZE_SMALL = iota
	SIZE_MEDIUM
)
