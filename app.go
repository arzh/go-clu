package clu

import "errors"

type Command interface {
	Init(ArgConf)
	Cmd(Args)
}

type Cmds map[string]Command

type AppConf interface {
	AddFlag(string, string, string)
	AddVar(string, string, string, string)
	AddCmd(string, Command)
}

type App struct {
	Args
	cmds Cmds
}

func newApp() *App {
	app := new(App)

	app.values = make(map[string]*string)
	app.flags = make(map[string]*bool)
	app.help = make(map[string]*string)
	app.loosies = make([]string, 0)
	app.cmds = make(Cmds)

	return app

}

func (app *App) AddCmd(name string, fp Command) {
	app.cmds[name] = fp
}

func (app *App) Run() error {
	if len(app.loosies) > 0 {
		if fp := app.cmds[app.loosies[0]]; fp != nil {
			// app.loosies = app.loosies[1:]
			// fp(app.Args)
			// nApp := 
			// return nil
			a := newArgs()
			fp.Init(a)
			fp.Cmd(*a)
		}
	}

	return errors.New("Command unknown")
}