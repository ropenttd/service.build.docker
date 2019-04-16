package protobuf

import (
	"fmt"
	"github.com/ropenttd/tsubasa/generics/api/protobuf"
)

type ApiError struct {
	Message  string
	Code     v1_generics.StatusCode
	HttpCode int32
	Stack    error
}

func (e ApiError) Error() string {
	return fmt.Sprint(e.Stack)
}

// MarshalStatusMessage marshals this ApiError into a StatusResponse.
func (e ApiError) MarshalStatusMessage() (message *v1_generics.StatusResponse) {
	return &v1_generics.StatusResponse{Message: fmt.Sprint(e.Error()), Status: e.Code}
}
