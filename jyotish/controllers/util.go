package controllers

import "strings"

func SplitPath(path string) []string {
	components := strings.Split(path, "/")
	if components[len(components)-1] == "" {
		return components[1 : len(components)-1]
	} else {
		return components[1:]
	}
}
