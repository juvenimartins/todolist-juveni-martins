package controller

import "strconv"

//parse to int
func ParseInt(id string) int {
	i, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return i
}
