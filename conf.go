package ronin

import (
	"go.uber.org/fx"
)

type Environment int

const (
	Development Environment = iota + 1
	Production
)

// String - give the type a String function
func (e Environment) String() string {
	return [...]string{"development", "production"}[e-1]
}

// EnumIndex -  give the type a EnumIndex function
func (e Environment) Index() int {
	return int(e)
}

func ParseEnv(str string) Environment {
	switch str {
	case Development.String():
		return Development
	case Production.String():
		return Production
	default:
		return Development
	}
}

type Configuration struct {
	Stage string `conf:"cfg_stage"`
}

func (c Configuration) GetStage() Environment {
	return ParseEnv(c.Stage)
}

type ResultConf struct {
	fx.Out

	*Configuration
}

func Config() (ResultConf, error) {
	conf, err := Conf[Configuration]("./", "")
	if err != nil {
		return ResultConf{}, err
	}
	return ResultConf{
		Configuration: conf,
	}, nil
}

var ModConf = fx.Module(
	"Configuration Module",
	fx.Provide(Config),
)
