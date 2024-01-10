package application

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
)

var pathToTemplates = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	DataMap       map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	User          *models.WalletAccount
}

func (app *Application) Render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/thank-you-activation.gohtml", pathToTemplates),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplates, t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, app.AddDefaultData(td, r)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	// if app.isAuthenticated(r) {
	// 	td.Authenticated = true
	// 	user, ok := app.Session.Get(r.Context(), "user").(data.User)
	// 	if !ok {
	// 		app.ErrorLog.Println("Cant get user from session")
	// 	} else {
	// 		td.User = &user
	// 	}

	// }

	// td.Now = time.Now()

	return td

}

func (app *Application) isAuthenticated(r *http.Request) bool {

	return true
}
