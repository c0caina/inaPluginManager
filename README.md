# inaPluginManager
This add-in allows you to load plugins in the .so format.

- [x] Compiling the plugins from the source code.
- [x] Loading plugins in the .so format.
- [x] Unload assets in plugins/assets. 
## Usage
In fact, the Start function is called in PluginManager, which provides a link to *Logger and *Server, and also gives your assets.

[inaPluginTemplate](https://github.com/c0caina/inaPluginTemplate "Plugin template"):
```golang
package main

import (
	"embed"
	"github.com/df-mc/dragonfly/server"
	"github.com/sirupsen/logrus"
)

//go:embed plugins
var assets embed.FS

func Start(logger *logrus.Logger, server *server.Server) embed.FS {
	logger.Info("I running")
	
	return assets
}
```
***Please note that the versions of plugin packages must be identical to the versions of inaPluginManager.***
## Contributing
I am an extremely inexperienced programmer, but I tried to make this bootloader. I would be glad if you point out my non-obvious mistakes.
## Contact
I can be found on the [Dragonfly](https://github.com/df-mc/dragonfly "github Dragonfly") server.

[![Chat on Discord](https://img.shields.io/badge/Chat-On%20Discord-738BD7.svg?style=for-the-badge)](https://discord.gg/U4kFWHhTNR)
