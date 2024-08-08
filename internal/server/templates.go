package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/cjbagley/colinbagley.dev/internal/data"
)

type page struct {
	Title     string
	Published string
	Updated   string
	Articles  []data.Article
}

// GetTextDate takes a date in the format 2006-1-2 and
// formats it to a full text string, for example
// 8th August 2024
func GetTextDate(date string) string {
	if date == "" {
		return ""
	}
	t, err := time.Parse("2006-1-2", date)
	if err != nil {
		return ""
	}

	var suffix string
	switch t.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	default:
		suffix = "th"
	}

	return t.Format(fmt.Sprintf("2%s January 2006", suffix))
}

type pageTemplates interface {
	getTemplates() []string
	getData() page
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

func (c *contentPageTemplates) getData() page {
	return page{Title: c.title, Articles: c.articles}
}

type articlePageTemplates struct {
	article data.Article
}

func (a *articlePageTemplates) getTemplates() []string {
	return []string{
		getLayoutDirPath("main.gohtml"),
		getPageDirPath("articles/" + a.article.Template + ".gohtml"),
		getLayoutDirPath("partials/article.gohtml"),
		getLayoutDirPath("partials/published.gohtml"),
		getLayoutDirPath("partials/updated.gohtml"),
	}
}

func (a *articlePageTemplates) getData() page {
	return page{Title: a.article.Title, Published: a.article.Published, Updated: a.article.Updated}
}

// WriteHTTPResponse writes the given templates to the HTTP response writer
func WriteHTTPResponse(w http.ResponseWriter, templates pageTemplates) {
	funcs := template.FuncMap{
		"textDate": GetTextDate,
	}

	tpl, err := template.New("main.gohtml").Funcs(funcs).ParseFiles(templates.getTemplates()...)
	if err != nil {
		LogError(err)
		serveErrorPage(w)
		return
	}

	err = tpl.Execute(w, templates.getData())
	if err != nil {
		LogError(err)
		serveErrorPage(w)
		return
	}

	w.Header().Set("Content-Type", "text/html")
}

func serveErrorPage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	f, err := os.ReadFile(getPageDirPath("500.html"))
	if err != nil {
		LogError(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	if _, err = w.Write(f); err != nil {
		LogError(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
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
