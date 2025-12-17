// Package main implements the GitHub plugin for Relicta.
package main

import (
	"github.com/relicta-tech/relicta-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&GitHubPlugin{})
}
