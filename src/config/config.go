package config

import "path/filepath"

const (
	TemplatesDir = "templates"
	StaticDir    = "static"
	CssDir       = "static/css"
)

func GetTemplatePath(templateName string) string {
	return filepath.Join(TemplatesDir, templateName)
}
