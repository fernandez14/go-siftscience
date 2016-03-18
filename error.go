package gosift

import (
	"fmt"
)

type SiftError struct {
	Status       int
	ErrorMessage string `json:"error_message"`
	Time         int64
	Request      string
}

func (e SiftError) Error() string {
	return fmt.Sprintf("%d - %s", e.Status, e.ErrorMessage)
}