package environment

import "os"

// GetEnv looks in the environment variables for the given key,
// falling back to the fallback string if it is not found.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
