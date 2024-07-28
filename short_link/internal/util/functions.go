package util

import "regexp"

func IsValidURLPath(path string) bool {
	// Definimos una expresión regular que solo permite números, letras, guiones y guiones bajos.
	var validPathPattern = `^[a-zA-Z0-9\-_]*$`

	matched, err := regexp.MatchString(validPathPattern, path)
	if err != nil {
		// Si hay un error con la expresión regular, devolvemos false
		return false
	}

	return matched
}
