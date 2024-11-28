package ronin

import (
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func Run[T LifeCycle](l T) error {
	result, err := Config()
	if err != nil {
		return err
	}
	InitLog(result.GetStage().String())
	log.Info().Str("Stage:", result.GetStage().String()).Msg("Environment stage")
	log.Info().Msg("initializer configuration")
	log.Info().Str("version", l.Version()).Msg("life cycle")
	// base modules
	fx.New(append([]fx.Option{
		fx.Provide(Config),
		fx.WithLogger(Logger),
	}, l.Yoroi())...).Run()
	return nil
}
