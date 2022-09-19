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
		input.Data = make(map[string]interface{})
		input.IntMap = make(map[string]int)
		input.StringMap["model"] = "Benz"
		input.Data["name"] = "Michael"
		input.Data["child"] = 1
		input.CSRF = "Security check!"
		input.Error = "Error Detected!"
		input.Flash = "This is a message."
		input.Warning = "I warn you!"
		return &input, nil
	}
	return nil, errors.New("no variable found")
}
