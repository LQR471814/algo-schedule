package db

import _ "embed"

//go:embed schema.sql
var Schema string

type Size = int64

const (
	SIZE_SMALL = iota
	SIZE_MEDIUM
)

type Challenge = int64

const (
	CHALLENGE_EASY Challenge = iota
	CHALLENGE_MEDIUM
	CHALLENGE_HARD
)
