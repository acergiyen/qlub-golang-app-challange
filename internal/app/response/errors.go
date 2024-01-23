package response

import "fmt"

func ErrInvalidRequest(err error) []byte {
	return []byte(fmt.Sprintf("Invalid Request - %v", err))
}
