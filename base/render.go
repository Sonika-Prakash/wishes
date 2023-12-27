package base

import (
	"fmt"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

func (a *Application) render(w http.ResponseWriter, r *http.Request, viewName string, vars jet.VarMap) error {

	template, err := a.view.GetTemplate(fmt.Sprintf("%s.html", viewName))
	if err != nil {
		return err
	}
	if err := template.Execute(w, vars, nil); err != nil {
		return err
	}

	return nil
}
