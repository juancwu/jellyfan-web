package util

import "strings"

func JoinAcceptMimeTypes(accept []string) string {
	return strings.Join(accept, ",")
}
