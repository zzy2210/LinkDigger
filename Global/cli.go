package Global

import "github.com/urfave/cli"

var Finder *cli.App
var Limit int
func init(){
	Finder  = cli.NewApp()
	Finder.Name = "LinkDigger"
	Finder.Version = "1.0.0"
	Finder.Usage = "A tool to find webserver's path"

}
