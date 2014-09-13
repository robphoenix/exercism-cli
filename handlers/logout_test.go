package handlers

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/codegangsta/cli"
	"github.com/exercism/cli/config"
	"github.com/stretchr/testify/assert"
)

func TestLogoutDeletesConfigFile(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)

	file := filepath.Join(tmpDir, config.File)

	c := config.Config{}
	c.SavePath(file)
	c.Write()

	set := flag.NewFlagSet("global", 0)
	set.String("config", file, "about this option")
	ctx := cli.NewContext(nil, nil, set)

	Logout(ctx)

	_, err = os.Stat(file)
	if err == nil {
		t.Errorf("File exists: %s", file)
	}
}
