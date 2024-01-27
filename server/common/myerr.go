package common

import "fmt"

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d, err: %s", e.Code, e.Msg)
}
