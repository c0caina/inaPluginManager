package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"inaPluginManager/global"
	"inaPluginManager/plugin"
	"inaPluginManager/source"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	global.Log.Formatter = &logrus.TextFormatter{ForceColors: true}
	global.Log.Level = logrus.DebugLevel

	if ok := supportOS("linux", "freebsd", "macos"); !ok {
		global.Log.Fatalln("Your operating system does not support plugins. Currently plugins are only supported on Linux, FreeBSD, and macOS.")
	}

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config, err := readConfig()
	if err != nil {
		global.Log.Fatalln(err)
	}

	if err := os.MkdirAll(global.Plugins, 0750); err != nil {
		global.Log.Fatalf("failed createing directory: %v", err)
	}

	sources := source.New(global.Plugins)
	sources.Build()

	srv := server.New(&config, global.Log)
	srv.CloseOnProgramEnd()
	if err := srv.Start(); err != nil {
		global.Log.Fatalln(err)
	}

	plug := plugin.New(global.Plugins)
	plug.Load(global.Log, srv)

	for {
		if _, err := srv.Accept(); err != nil {
			return
		}
	}
}

// readConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func readConfig() (server.Config, error) {
	c := server.DefaultConfig()
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := ioutil.WriteFile("config.toml", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
}

func supportOS(a ...string) bool {
	for _, OS := range a {
		if runtime.GOOS == OS {
			return true
		}
	}
	return false
}
