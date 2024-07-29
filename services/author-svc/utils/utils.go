package utils

func GetPage(offset int, limit int) int {
	return (offset - 1) * limit
}
