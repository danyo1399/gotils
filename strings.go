package gotils

import "unicode"

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(s string) string {
	return ToSnakeCaseWithSeparator(s, '_')
}

func ToSnakeCaseWithSeparator(s string, separator rune) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, separator)
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
