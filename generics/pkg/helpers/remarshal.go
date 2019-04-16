package helpers

// MarshalStringStringMapToPointer converts a map of map[string]string to *map[string]*string
// because docker is fucking stupid and I hate its guts.
func MarshalStringStringMapToPoint(input *map[string]string) (output *map[string]*string) {
	w := map[string]*string{}
	for k, v := range *input {
		w[k] = &v
	}
	return &w
}

// UnnarshalStringPointerMapToString converts a map of map[string]*string to *map[string]string
// because docker is fucking stupid and I hate its guts.
func UnmarshalStringStringMapToPoint(input *map[string]*string) (output *map[string]string) {
	w := map[string]string{}
	for k, v := range *input {
		w[k] = *v
	}
	return &w
}
