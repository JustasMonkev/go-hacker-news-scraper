package helper

import "log"

// Check Generic error handler
// used to return a value or a err
func Check[T any](value T, err error) T {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return value
}

// CheckErr used only to return a err
func CheckErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
