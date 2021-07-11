package plugin

import (
	"embed"
	"github.com/df-mc/dragonfly/server"
	"github.com/sirupsen/logrus"
	"github.com/soypat/rebed"
	"inaPluginManager/global"
	"path/filepath"
	"plugin"
	"strings"
)

type Plugins []Plugin

type Plugin struct {
	Name string
	Path string
}

func New(searchPath string) *Plugins {
	var plugins Plugins

	paths, _ := filepath.Glob(searchPath + "*.so")
	for _, path := range paths {
		name := strings.TrimPrefix(strings.TrimSuffix(path, ".so"), searchPath)
		plugins = append(Plugins{}, Plugin{Name: name, Path: path})
	}
	return &plugins
}

func (plugins Plugins) Load(log *logrus.Logger, srv *server.Server) {
	for _, plug := range plugins {
		p, err := plugin.Open(plug.Path)
		if err != nil {
			global.Log.Errorf("[%s] failed plugin loading: %v", plug.Name, err)
			continue
		}
		symbol, err := p.Lookup("Start")
		if err != nil {
			global.Log.Errorf("[%s] failed plugin loading: %v", plug.Name, err)
			continue
		}
		startFunc, ok := symbol.(func(logger *logrus.Logger, server *server.Server) embed.FS)
		if !ok {
			global.Log.Errorf("[%s] failed plugin loading: %v", plug.Name, err)
			continue
		}
		if err := rebed.Patch(startFunc(log, srv), ""); err != nil {
			global.Log.Errorf("[%s] failed create assets: %v", plug.Name, err)
		}
		global.Log.Infof("[%s] successfully loaded.", plug.Name)
	}
}
