// +build man

package main

import (
	"os"
	"strings"

	"github.com/kyoh86/gobase/app"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var manFlags struct{}

var manCommand = &cobra.Command{
	Use:    "man",
	Short:  "Generate manual",
	Hidden: true,
	Args:   cobra.ExactArgs(0),
	RunE: func(*cobra.Command, []string) error {
		header := &doc.GenManHeader{
			Title:   strings.ToUpper(app.Name),
			Section: "1",
		}
		if err := doc.GenManTree(facadeCommand, header, "."); err != nil {
			return err
		}
		if err := os.MkdirAll("./usage", 0755); err != nil {
			return err
		}
		facadeCommand.DisableAutoGenTag = true
		if err := doc.GenMarkdownTree(facadeCommand, "./usage"); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	facadeCommand.AddCommand(manCommand)
}
