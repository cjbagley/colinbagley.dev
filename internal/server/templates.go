package server

import (
	"fmt"
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"html/template"
	"io"
	"testing"
)

type pageTemplates interface {
	getTemplates() []string
}

type contentPageTemplates struct {
	path string
}

func (c *contentPageTemplates) getTemplates() []string {
	return []string{
		getLayoutDirPath("main.gohtml"),
		getPageDirPath(c.path),
	}
}

type articlePageTemplates struct {
	article data.Article
}

func (a *articlePageTemplates) getTemplates() []string {
	return []string{
		getLayoutDirPath("main.gohtml"),
		getPageDirPath("articles/" + a.article.Template + ".gohtml"),
		getLayoutDirPath("partials/article.gohtml"),
	}
}

func WriteTemplate(writer io.Writer, templates pageTemplates) (*template.Template, error) {
	tpl := template.Must(template.ParseFiles(templates.getTemplates()...))
	tpl.Execute(writer, nil)
	return tpl, nil
}

func getPageDirPath(template string) string {
	if testing.Testing() {
		return fmt.Sprintf("../../web/pages/%s", template)
	}

	return fmt.Sprintf("./web/pages/%s", template)
}

func getLayoutDirPath(template string) string {
	if testing.Testing() {
		return fmt.Sprintf("../../web/layouts/%s", template)
	}

	return fmt.Sprintf("./web/layouts/%s", template)
}
