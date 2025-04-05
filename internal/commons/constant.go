package commons

import (
	"html/template"
)

var Tmpl *template.Template

var (
	CountCluster   = 0
	ConfigReloadCh = make(chan struct{})
)
