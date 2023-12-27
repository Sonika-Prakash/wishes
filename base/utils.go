package base

import (
	"fmt"
	"strings"

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
	if strings.Contains(name, " ") {
		splitName := strings.Split(name, " ")
		name = strings.Join(splitName, "%20")
	}
	return fmt.Sprintf("%s?n=%s", a.server.renderHost, name)
}

func (a *Application) generateCopyMsg(link string) string {
	return fmt.Sprintf("Hey have you seen this?\nI have a surprise message for you!\nClick on the below link.\n\n%s\n", link)
}

func getCapitalizedName(name string) string {
	return cases.Title(language.English, cases.Compact).String(name)
}
