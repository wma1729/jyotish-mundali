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

func StringToFloat32(value string) float32 {
	f64, _ := strconv.ParseFloat(value, 32)
	return float32(f64)
}
