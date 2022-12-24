package delivery

import "fmt"

func ErrorMsg(msg error) string {
	return fmt.Sprintf("%v Error is happened in the process", msg)
}
