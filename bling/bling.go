package bling

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

type Shell string

var (
	Bash Shell = "bash"
	Zsh  Shell = "zsh"
	Fish Shell = "fish"
)

type Bling struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Packages    []string `yaml:"packages"`
	Programs    []string `yaml:"programs"`
	PackageMap  map[string]*Package
	ProgramMap  map[string]*Program
}

func (b *Bling) IncludesPackage(name string) bool {
	for _, pkg := range b.Packages {
		if pkg == name {
			return true
		}
	}
	return false
}
func (b *Bling) IncludesProgram(name string) bool {
	for _, pkg := range b.Programs {
		if pkg == name {
			return true
		}
	}
	return false
}
func (b *Bling) InitBash() string {
	sb := strings.Builder{}
	for name, prog := range b.ProgramMap {
		if prog.Init != nil {
			sb.WriteString("# " + name + "\n")
			sb.WriteString(prog.Init.Bash)
		}
	}
	return sb.String()
}
func (b *Bling) InitZsh() string {
	sb := strings.Builder{}
	for name, prog := range b.ProgramMap {
		if prog.Init != nil {
			sb.WriteString("# " + name + "\n")
			sb.WriteString(prog.Init.Zsh)
		}
	}
	return sb.String()
}

func (b *Bling) Aliases() string {
	sb := strings.Builder{}
	for name, prog := range b.ProgramMap {
		if len(prog.Aliases) > 0 {
			sb.WriteString("# " + name + "\n")
			for _, alias := range prog.Aliases {
				sb.WriteString("# " + alias.Description + "\n")
				sb.WriteString("alias " + alias.Key + "='" + alias.Value + "'\n")
			}
		}
	}
	return sb.String()
}

func ProgramInLevels(name string) []string {
	included := make([]string, 0, 4)
	none, err := NoBling()
	if err != nil {
		return []string{}
	}
	low, err := LowBling()
	if err != nil {
		return []string{}
	}
	dflt, err := DefaultBling()
	if err != nil {
		return []string{}
	}
	high, err := HighBling()
	if err != nil {
		return []string{}
	}
	if none.IncludesProgram(name) {
		included = append(included, "None")
	}
	if low.IncludesProgram(name) {
		included = append(included, "Low")
	}
	if dflt.IncludesProgram(name) {
		included = append(included, "Default")
	}
	if high.IncludesProgram(name) {
		included = append(included, "High")
	}
	return included

}

var (
	//go:embed none.yml
	none []byte
	//go:embed low.yml
	low []byte
	//go:embed default.yml
	dflt []byte
	//go:embed high.yml
	high []byte
)

func loadBling(bytes []byte) (*Bling, error) {

	var b Bling

	err := yaml.Unmarshal(bytes, &b)
	if err != nil {
		return &b, err
	}
	progs, err := LoadPrograms()
	if err != nil {
		return &b, err
	}
	pkgs, err := LoadPackages()
	if err != nil {
		return &b, err
	}
	packMap := make(map[string]*Package, len(b.Packages))
	progMap := make(map[string]*Program, len(b.Programs))

	for _, pkg := range pkgs {
		packMap[pkg.Name] = pkg
	}
	for _, prog := range progs {
		progMap[prog.Name] = prog
	}
	b.PackageMap = lo.PickByKeys(packMap, b.Packages)
	b.ProgramMap = lo.PickByKeys(progMap, b.Programs)

	return &b, nil
}

func NoBling() (*Bling, error) {

	return loadBling(none)
}
func LowBling() (*Bling, error) {
	return loadBling(low)
}
func DefaultBling() (*Bling, error) {
	return loadBling(dflt)
}
func HighBling() (*Bling, error) {
	return loadBling(high)
}

func FromString(bling string) (*Bling, error) {

	switch bling {
	case "none":
		return NoBling()
	case "low":
		return LowBling()
	case "default":
		return DefaultBling()
	case "high":
		return HighBling()
	default:
		return nil, fmt.Errorf("unknown bling: %s", bling)
	}
}
