package main

import (
	_ "github.com/AndreVuillemot160/gclone/backend/all" // import all backends
	_ "github.com/AndreVuillemot160/gclone/backend/drive"
	_ "github.com/AndreVuillemot160/gclone/cmd/all" // import all commands
	"github.com/rclone/rclone/cmd"
	"github.com/rclone/rclone/fs"
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

func main() {
	fs.Version = fs.Version + "-mod-V20210901"
	cmd.Main()
}
