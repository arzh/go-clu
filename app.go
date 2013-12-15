package clu

import "errors"

type CmdFunc func(Args)

type AppConf interface {
	AddFlag(string, string, string)
	AddVar(string, string, string, string)
	AddCmd(string, CmdFunc)
}

type App struct {
	Args
	cmds map[string]CmdFunc
}

func newApp() *App {
	app := new(App)

	app.values = make(map[string]*string)
	app.flags = make(map[string]*bool)
	app.help = make(map[string]*string)
	app.loosies = make([]string, 0)
	app.cmds = make(map[string]CmdFunc)

	return app

}

func (app *App) AddCmd(name string, fp CmdFunc) {
	app.cmds[name] = fp
}

func (app *App) Run() error {
	if len(app.loosies) > 0 {
		if fp := app.cmds[app.loosies[0]]; fp != nil {
			app.loosies = app.loosies[1:]
			fp(app.Args)
			return nil
		}
	}

	return errors.New("Command unknown")
}