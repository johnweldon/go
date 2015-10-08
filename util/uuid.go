package util

import (
	g "github.com/pborman/uuid"
)

type UUID string

func NewUUID() UUID {
	return UUID(g.New())
}
