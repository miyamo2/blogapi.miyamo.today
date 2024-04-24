package utils

// PtrFromString returns a pointer to a string.
// If v is zero value, returns nil.
func PtrFromString(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
