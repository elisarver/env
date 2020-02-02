// Package env enhances os.Getenv and os.LookupEnv
// with fallback and must functionality.
package env

import (
	"os"
)

// Get is just a wrap on os.Getenv()
func Get(key string) string {
	return os.Getenv(key)
}

// Get is just a wrap on os.LookupEnv()
func Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}

// FallbackLookup is similar to Lookup, but has a fallback value.
// It returns the found/fallback value, and a boolean term stating
// whether the value was found in the environment.
// It will return false when it returns the fallback, in case you need
// to warn users.
func FallbackLookup(key, fallback string) (string, bool) {
	value, found := Lookup(key)
	if !found {
		value = fallback
	}
	return value, found
}

// MustLookup is similar to os.LookupEnv, but panics with a LookupError
// on failure. This is meant to handle startup configuration, especially
// in places where it is unwise to store a sensitive value in code.
func MustLookup(key string) string {
	value, found := Lookup(key)
	if found {
		return value
	}
	panic(LookupError{
		Key: key,
	})
}

// LookupError represents a missing value in the environment. It is
// thrown as a panic from MustLookup.
type LookupError struct {
	// Key is the name of the requested key
	Key string
}

// Error fulfills the error interface.
func (e LookupError) Error() string {
	return `env: MustLookup("` + e.Key + `"): Key not found`
}
