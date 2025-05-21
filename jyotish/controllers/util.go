package controllers

import (
	"strconv"
	"strings"
)

func SplitPath(path string) []string {
	components := strings.Split(path, "/")
	if components[len(components)-1] == "" {
		return components[1 : len(components)-1]
	} else {
		return components[1:]
	}
}

func StringToFloat(value string) float64 {
	f, _ := strconv.ParseFloat(value, 32)
	return f
}
