package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetHostByURL(url string) string {
	if url == "" {
		return ""
	}
	if strings.HasPrefix(url, "/") {
		return ""
	}
	urlStrArray := strings.Split(url, "//")
	if len(urlStrArray) != 2 {
		return ""
	}
	schema := urlStrArray[0]
	host := strings.Split(urlStrArray[1], "/")[0]
	return fmt.Sprintf("%s//%s", schema, host)
}

func IsExist(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}
