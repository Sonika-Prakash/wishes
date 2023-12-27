package base

import (
	"fmt"

	"github.com/CloudyKit/jet/v6"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// initJet initialises the jet engine
func initJet() *jet.Set {
	return jet.NewSet(
		jet.NewOSFileSystemLoader("./views"),
		jet.InDevelopmentMode(),
	)
}

func (a *Application) generateURL(name string) string {
	return fmt.Sprintf("%s:%s?n=%s", a.server.host, a.server.port, name)
}

func getCapitalizedName(name string) string {
	return cases.Title(language.English, cases.Compact).String(name)
}
