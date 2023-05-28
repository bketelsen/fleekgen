package hook

import (
	"bytes"
	"embed"

	"text/template"

	"github.com/bketelsen/fleekgen/bling"
)

var (
	tmplNewBuf = bytes.NewBuffer(make([]byte, 0, 4096))
	zshBuf     = bytes.NewBuffer(make([]byte, 0, 4096))
	bashBuf    = bytes.NewBuffer(make([]byte, 0, 4096))
)

func FromBling(bling *bling.Bling) ([]byte, error) {
	var err error
	tmplNewBuf.Reset()
	// TODO: cache template parsing
	tmpl, err := template.ParseFS(tmplFS, "tmpl/init.sh.tmpl")
	if err != nil {
		return []byte{}, err
	}
	err = tmpl.Execute(tmplNewBuf, bling)
	if err != nil {
		return []byte{}, err
	}
	return tmplNewBuf.Bytes(), nil
}

func Zshrc(bling *bling.Bling) ([]byte, error) {
	var err error
	zshBuf.Reset()
	// TODO: cache template parsing
	tmpl, err := template.ParseFS(tmplFS, "tmpl/zsh/zshrc.tmpl")
	if err != nil {
		return []byte{}, err
	}
	err = tmpl.Execute(zshBuf, bling)
	if err != nil {
		return []byte{}, err
	}
	return zshBuf.Bytes(), nil
}
func Bashrc(bling *bling.Bling) ([]byte, error) {
	var err error
	bashBuf.Reset()
	// TODO: cache template parsing
	tmpl, err := template.ParseFS(tmplFS, "tmpl/bash/bashrc.tmpl")
	if err != nil {
		return []byte{}, err
	}
	err = tmpl.Execute(bashBuf, bling)
	if err != nil {
		return []byte{}, err
	}
	return bashBuf.Bytes(), nil
}

//go:embed tmpl/* tmpl/bash/* tmpl/zsh/*
var tmplFS embed.FS
