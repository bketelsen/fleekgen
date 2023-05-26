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
		InitHook []string            `json:"init_hook,omitempty"`
		Scripts  map[string][]string `json:"scripts,omitempty"`
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
	c.Shell.InitHook = []string{". ${DEVBOX_GLOBAL_ROOT}/init.sh"}
	c.Env = make(map[string]string)
	c.Shell.Scripts = scripts
	c.Env["DEVBOX_GLOBAL_PREFIX"] = "$HOME/.local/share/devbox/global/default/.devbox/nix/profile/default"
	c.Env["DEVBOX_GLOBAL_ROOT"] = "$HOME/.local/share/devbox/global/current"
	return c
}
