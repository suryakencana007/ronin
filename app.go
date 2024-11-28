package ronin

import (
	"go.uber.org/fx"
)

// Option is server type return func.
type Option = func(a *App) error

// Yoroi will assign to armours field app.
func Yoroi(armours ...fx.Option) Option {
	return func(a *App) error {
		a.armor = fx.Options(armours...)
		return nil
	}
}

// SetName will assign to name field app.
func SetName(name string) Option {
	return func(a *App) error {
		a.name = name
		return nil
	}
}

// SetVersion will assign to version field app.
func SetVersion(version string) Option {
	return func(a *App) error {
		a.version = version
		return nil
	}
}

// App is application.
type App struct {
	name    string
	version string
	armor   fx.Option
}

// Version return version app.
func (a *App) Version() string {
	return a.version
}

// Name return name app.
func (a *App) Name() string {
	return a.name
}

// Yoroi return plugins app.s
func (a *App) Yoroi() fx.Option {
	return a.armor
}

// Ryu is Ronin martial art style.
func Ryu(opts ...Option) LifeCycle {
	a := &App{}
	for _, o := range opts {
		if err := o(a); err != nil {
			panic(err)
		}
	}
	return a
}
