package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl"
	"github.com/garlsecurity/securepassctl/spctl/cmd"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

var (
	// OptionDebug contains the --debug flag
	OptionDebug bool
	// Version of spctl
	Version string
	// GNUHelpStyle indicates whether enable GNU-like help screen
	GNUHelpStyle string
)

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
}

func loadConfiguration(userConfigFile string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if userConfigFile != "" {
		if info, err := os.Stat(userConfigFile); err != nil {
			return err
			//log.Fatalf("error: %v", err)
		} else if !info.Mode().IsRegular() {
			//			log.Fatalf("error: %q is not a regular file", userConfigFile)
			return fmt.Errorf("%q is not a regular file", userConfigFile)
		}
		service.LoadConfiguration([]string{userConfigFile})
		return nil
	}
	SystemConfigFiles := []string{"/etc/securepass.conf",
		"/usr/local/etc/securepass.conf",
		filepath.Join(cwd, "securepass.conf")}
	service.LoadConfiguration(SystemConfigFiles)
	return nil
}

func main() {
	if b, e := strconv.ParseBool(GNUHelpStyle); e != nil || b {
		cli.AppHelpTemplate = `Usage: {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Flags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
{{.Usage}}
  {{if .Flags}}
  {{range .Flags}}{{.}}
  {{end}}{{end}}{{if .Commands}}
Commands:
    {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
    {{end}}{{end}}

spctl home page: <https://github.com/garlsecurity/securepassctl>
SecurePass online help: <http://www.secure-pass.net/integration-guides-examples/>
Report bugs to <https://github.com/garlsecurity/securepassctl/issues>
`
		cli.CommandHelpTemplate = `Usage: {{.HelpName}}{{if .Flags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
{{.Usage}}
  {{if .Flags}}
  {{range .Flags}}{{.}}
  {{end}}{{end}}{{if .Subcommands}}
Commands:
    {{range .Subcommands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
    {{end}}{{end}}
`
		cli.SubcommandHelpTemplate = `Usage: {{.HelpName}}{{if .Flags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
{{.Usage}}
  {{if .Flags}}
  {{range .Flags}}{{.}}
  {{end}}{{end}}{{if .Commands}}
Commands:
    {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
    {{end}}{{end}}
`
	}
	args := handleCompatMode(os.Args)
	a := cli.NewApp()
	a.Name = "spctl"
	a.Usage = "Manage distributed identities."
	a.Author = "Alessio Treglia"
	a.Email = "alessio@debian.org"
	a.Copyright = "Copyright © 2016 Alessio Treglia <alessio@debian.org>"
	a.Version = Version
	a.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, D",
			Usage:       "enable debug output",
			Destination: &OptionDebug,
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "configuration file",
		},
	}
	a.Commands = cmd.Commands
	a.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			securepassctl.DebugLogger.SetOutput(os.Stderr)
		}
		err := loadConfiguration(c.GlobalString("config"))
		return err
	}

	a.Run(args)
}
