package main

import "fmt"

type OverrideError interface {
	Error() string
}

type myErrors struct {
	ErrCode int
	ErrMsg  string
}

func (e *myErrors) Error() string {
	return fmt.Sprintf("%d: %s", e.ErrCode, e.ErrMsg)
}
