package utils

import "strings"

func MakeGoStructKey(key string) string {

	// Make character array from the key
	charArray := strings.Split(key, "")

	// Convert the first character to upper case
	charArray[0] = strings.ToUpper(charArray[0])

	// Make letter right after underscore to upper case
	for i := 1; i < len(charArray); i++ {
		if charArray[i] == "_" {
			charArray[i+1] = strings.ToUpper(charArray[i+1])
		}
	}

	newKey := strings.Join(charArray, "")

	// Remove the underscore
	newKey = strings.ReplaceAll(newKey, "_", "")

	// Return the new key
	return newKey
}
