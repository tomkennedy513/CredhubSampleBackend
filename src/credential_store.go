package src

var valuesMap = map[string][]byte{}

func SetValue(path string, value []byte) {
	valuesMap[path] = value
}

func GetValue(path string) ([]byte, bool) {
	value, exists := valuesMap[path]
	return value, exists
}

func DeleteValue(path string) bool {
	var exists bool

	if _, exists = valuesMap[path]; exists {
		delete(valuesMap, path)
	}

	return exists
}

