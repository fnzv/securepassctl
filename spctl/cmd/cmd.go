package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/cmd/app"
	"github.com/garlsecurity/securepassctl/spctl/cmd/config"
	"github.com/garlsecurity/securepassctl/spctl/cmd/groupmember"
	"github.com/garlsecurity/securepassctl/spctl/cmd/logs"
	"github.com/garlsecurity/securepassctl/spctl/cmd/ping"
	"github.com/garlsecurity/securepassctl/spctl/cmd/radius"
	"github.com/garlsecurity/securepassctl/spctl/cmd/user"
)

// Commands contains the app's commands
var Commands = []cli.Command{
	ping.Command,
	app.Command,
	config.Command,
	groupmember.Command,
	logs.Command,
	radius.Command,
	user.Command,
}
