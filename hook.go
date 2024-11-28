package ronin

import (
	"go.uber.org/fx"
)

type LifeCycle interface {
	// Version of the service.
	Version() string
	// Name of the service.
	Name() string
	// Yoroi is a type of traditional Japanese armor worn by ronin.
	Yoroi() fx.Option
}
