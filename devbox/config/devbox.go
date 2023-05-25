package config

import "github.com/bketelsen/fleekgen/bling"

// Config defines a devbox environment as JSON.
type Config struct {
	// Packages is the slice of Nix packages that devbox makes available in
	// its environment. Deliberately do not omitempty.
	Packages []string `json:"packages"`

	// Env allows specifying env variables
	Env map[string]string `json:"env,omitempty"`
	// Shell configures the devbox shell environment.
	Shell struct {
		// InitHook contains commands that will run at shell startup.
		InitHook []string `json:"init_hook,omitempty"`
	} `json:"shell,omitempty"`

	// Nixpkgs specifies the repository to pull packages from
	Nixpkgs NixpkgsConfig `json:"nixpkgs,omitempty"`
}

type NixpkgsConfig struct {
	Commit string `json:"commit,omitempty"`
}

func FromBling(b *bling.Bling) *Config {
	c := &Config{}
	c.Packages = b.Packages
	c.Packages = append(c.Packages, b.Programs...)
	c.Shell.InitHook = []string{". ~/.local/share/devbox/global/default/init.sh"}
	return c
}
