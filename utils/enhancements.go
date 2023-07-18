package utils

func Contains(s []HttpMethod, e HttpMethod) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
