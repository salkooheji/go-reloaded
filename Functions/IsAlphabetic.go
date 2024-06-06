package reloaded


func IsAlphabetic(s string) bool {
	// This function remains unchanged, checking for alphabetical characters
	for _, r := range s {
		if !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z') {
			return false
		}
	}
	return true
}