package devbox

import (
	"archive/tar"
	"encoding/json"

	"github.com/bketelsen/fleekgen/bling"
	"github.com/bketelsen/fleekgen/devbox/config"
	"github.com/bketelsen/fleekgen/devbox/hook"
)

type Global struct {
	Config *config.Config
	Bling  *bling.Bling
	Hook   string
}

func FromBling(b *bling.Bling) *Global {
	c := &Global{}
	c.Config = config.FromBling(b)
	c.Bling = b
	return c
}

func (g *Global) Files() (map[string][]byte, error) {
	files := make(map[string][]byte)

	dbjson, err := json.Marshal(g.Config)
	if err != nil {
		return files, err
	}
	initsh, err := hook.FromBling(g.Bling)
	if err != nil {
		return files, err
	}
	zshrc, err := hook.Zshrc(g.Bling)
	if err != nil {
		return files, err
	}
	bashrc, err := hook.Bashrc(g.Bling)
	if err != nil {
		return files, err
	}
	files["devbox.json"] = dbjson
	files["init.sh"] = initsh
	files["zsh/.zshrc"] = zshrc
	files["bash/.bashrc"] = bashrc
	return files, nil
}

func (g *Global) Write(files map[string][]byte, w *tar.Writer) error {

	for name, content := range files {
		hdr := &tar.Header{
			Name: name,
			Mode: 0644,
			Size: int64(len(content)),
		}
		if err := w.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := w.Write([]byte(content)); err != nil {
			return err
		}
	}
	return nil
}
