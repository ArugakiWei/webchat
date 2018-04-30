package common


func CheckFileFormat(name string) bool {
	ext := []string{".exe", ".js"}

	for _, v := range ext {
		if v == name {
			return false
		}
	}
	return true
}
