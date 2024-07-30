package server

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"

	"github.com/cjbagley/colinbagley.dev/internal/data"
)

type PageData struct {
	Title    string
	Articles []data.Article
}

type pageTemplates interface {
	getTemplates() []string
	getData() PageData
}

type contentPageTemplates struct {
	path     string
	title    string
	articles []data.Article
}

func (c *contentPageTemplates) getTemplates() []string {
	return []string{
		getLayoutDirPath("main.gohtml"),
		getPageDirPath(c.path),
	}
}

func (c *contentPageTemplates) getData() PageData {
	return PageData{Title: c.title, Articles: c.articles}
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

func WriteHttpResponse(w http.ResponseWriter, templates pageTemplates) error {
	tpl := template.Must(template.ParseFiles(templates.getTemplates()...))
	tpl.Execute(w, templates.getData())
	w.Header().Set("Content-Type", "text/html")

	return nil
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
