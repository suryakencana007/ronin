package ronin

import (
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
)

func Conf[C comparable](path, root string) (*C, error) {
	var cfg C
	kof := koanf.New("_")
	// load yaml
	if err := kof.Load(
		file.Provider(
			filepath.Join(path, ".conf"),
		), toml.Parser()); err != nil {
		return nil, err
	}

	// load dotenv
	if err := kof.Load(
		file.Provider(
			filepath.Join(path, ".env"),
		), dotenv.Parser()); err != nil {
		log.Warn().Err(err).Msg(".env file not found! check stage environment.")
	}

	// load env
	if err := kof.Load(
		env.Provider("RISE__", "_",
			func(s string) string {
				return strings.ToLower(strings.TrimPrefix(s, "RISE__"))
			},
		), nil); err != nil {
		return nil, err
	}
	if err := kof.UnmarshalWithConf(root, &cfg,
		koanf.UnmarshalConf{Tag: "conf", FlatPaths: true}); err != nil {
		return nil, err
	}

	return &cfg, nil
}
