package commons

import (
	"html/template"
	"os"
)

var Tmpl *template.Template

func getEnv(key, defaultKey string) (val string) {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultKey
}
