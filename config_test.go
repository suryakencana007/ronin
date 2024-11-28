package ronin_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/suryakencana007/ronin"
)

type conf struct {
	Key string `conf:"fiber_service"`
}

type second struct {
	Name string `conf:"fiber_name"`
	Port string `conf:"fiber_port"`
}

const toml = `[fiber]
service = "butter"
protocol = "http"
host = ":5000"
stage = "development"
`

const dotenv = `
FIBER_NAME=test
`

func TestConf(t *testing.T) {
	_ = os.Setenv("RISE__FIBER_PORT", "5000")
	tmpDir := t.TempDir()
	// toml
	pathToml := filepath.Join(tmpDir, "/.conf")
	// Write conf.yaml
	if err := os.WriteFile(pathToml, []byte(toml), 0644); err != nil {
		t.Fatalf("failed to write toml file: %v", err)
	}
	// .env
	pathDotenv := filepath.Join(tmpDir, "/.env")
	// Write .env
	if err := os.WriteFile(pathDotenv, []byte(dotenv), 0644); err != nil {
		t.Fatalf("failed to write env file: %v", err)
	}
	cf, err := ronin.Conf[conf](tmpDir, "")
	if err != nil {
		t.Fatalf("%v --- ", err)
	}
	t.Logf("key: %s ---- \n", cf.Key)

	sf, err := ronin.Conf[second](tmpDir, "")
	if err != nil {
		t.Fatalf("%v --- ", err)
	}
	t.Logf("name: %s ---- \n", sf.Name)
	t.Logf("port: %s ---- \n", sf.Port)
}
