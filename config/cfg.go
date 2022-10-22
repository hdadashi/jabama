package config

import (
	"errors"

	"github.com/alexedwards/scs/v2"
	render "github.com/hdadashi/jabama/3.render"
)

var PassingSession *scs.SessionManager

var input render.TemplateData

func GlobVar(VarName string) (*render.TemplateData, error) {
	if VarName == "input" {
		input.StringMap = make(map[string]string)
		input.IntMap = make(map[string]int)
		input.StringMap["model"] = "Benz"
		input.CSRF = "Security check!"
		input.Error = "Error Detected!"
		input.Flash = "This is a message."
		input.Warning = "I warn you!"
		return &input, nil
	}
	return nil, errors.New("no variable found")
}
