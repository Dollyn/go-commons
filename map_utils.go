package utils

func Values(m map[string]interface{}) []interface {} {
	result := make([]interface{}, len(m))
	i := 0
	for _, v := range m {
		result[i] = v
		i++
	}
	return result
}