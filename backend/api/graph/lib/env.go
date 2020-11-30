package lib

import "os"

// GetDefVal get value for envar key or set default val
func GetDefVal(key string, val string) string {
	if ok := os.Getenv(key); ok != "" {
		val = ok
	}
	return val
}
