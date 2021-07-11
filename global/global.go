package global

import "github.com/sirupsen/logrus"

const (
	// Plugins are the path where the plugins sources and executable file are located.
	Plugins = "plugins/"
)

var Log = logrus.New()
