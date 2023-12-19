package models

import "fmt"

type App struct {
	Name string
	Major int
	Minor int
	Build int
	Authors []string
}

func (app *App) GetName() string {
	return app.Name;
}

func (app *App) GetVersion() string {
	return fmt.Sprintf("Version: %d.%d.%d", app.Major, app.Minor, app.Build)
}