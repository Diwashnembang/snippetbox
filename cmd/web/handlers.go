package main

import (
	"diwashnembang/snippetbox/internal/models"
	"diwashnembang/snippetbox/internal/validator"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}
	data := app.newTemplateDate(r)
	data.Latest = snippets
	app.render(w, http.StatusOK, "home.tmpl.html", data)

}
func (app *application) sinppetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {

			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateDate(r)
	data.Snippet = snippet
	data.Flash = app.sessionMangaer.PopString(r.Context(),"flash")
	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// title := r.PostForm.Get("title")
	// content := r.PostForm.Get("content")
	// expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	var form snippetCreateForm
	err = app.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.infolog.Printf("expires is %v and type is %T", form.Expires, form.Expires)
	form.CheckField(validator.IsStringEmpty(form.Title), "title", "title cannot be empty")
	form.CheckField(validator.MaxChar(form.Title, 100), "title", "title cannot execeed more than 100 chars")
	form.CheckField(validator.IsStringEmpty(form.Content), "content", "content cannot be blank")
	form.CheckField(validator.NotPermitedInt(form.Expires, 365, 7, 1), "expires", "expires should be 7 , 365 or 1")

	if form.HasError() {
		data := app.newTemplateDate(r)
		data.Form = form
		app.render(w, http.StatusNotAcceptable, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionMangaer.Put(r.Context(), "flash", "Sinppet Was sucesfully created")
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateDate(r)
	data.Form = snippetCreateForm{
		Expires: 365,
		// Title:        "",
		// Content:      "",
		// FieldsErrors: make(map[string]string),
	}

	app.render(w, http.StatusOK, "create.tmpl.html", data)
}
