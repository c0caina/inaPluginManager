package source

import (
	"inaPluginManager/global"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Sources []Source

type Source struct {
	Name string
	Path string
}

func New(searchPath string) *Sources {
	var sources Sources

	paths, _ := filepath.Glob(searchPath + "*/main.go")
	for _, path := range paths {
		p := strings.TrimSuffix(path, "/main.go")
		n := strings.TrimPrefix(p, searchPath)
		sources = append(Sources{}, Source{Name: n, Path: p})
	}
	return &sources
}

func (sources Sources) Build() {
	for _, source := range sources {
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "../")
		cmd.Dir = source.Path
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			global.Log.Errorf("[%s] unsuccessfully built.", source.Name)
			continue
		}
		global.Log.Infof("[%s] successfully built.", source.Name)
	}
}
