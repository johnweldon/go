package util

import (
	g "code.google.com/p/go-uuid/uuid"
)

type UUID string

func NewUUID() UUID {
	return UUID(g.New())
}
