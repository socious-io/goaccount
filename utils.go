package goaccount

import "fmt"

func endpoint(path string, args ...any) string {
	return fmt.Sprintf("%s/%s", config.Host, fmt.Sprintf(path, args...))
}

func bearer(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}
