package server

import (
	"fmt"
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"html/template"
	"io"
	"testing"
)

type PageData struct {
	Title string
}

type pageTemplates interface {
	getTemplates() []string
	getData() PageData
}

type contentPageTemplates struct {
	path  string
	title string
}

func (c *contentPageTemplates) getTemplates() []string {
	return []string{
		getLayoutDirPath("main.gohtml"),
		getPageDirPath(c.path),
	}
}

func (c *contentPageTemplates) getData() PageData {
	return PageData{Title: c.title}
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

func (a *articlePageTemplates) getData() PageData {
	return PageData{Title: a.article.Title}
}

func WriteTemplate(writer io.Writer, templates pageTemplates) (*template.Template, error) {
	tpl := template.Must(template.ParseFiles(templates.getTemplates()...))
	tpl.Execute(writer, templates.getData())

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
