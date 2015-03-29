// Copyright 2015 JoongSeob Vito Kim. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/dorajistyle/whiplash/game"
)

func main() {
	app := cli.NewApp()
        app.Flags = []cli.Flag {
        cli.StringFlag{
    Name: "lang,l",
    Value: "english",
    Usage: "language for the game. not implemented yet.",
  },
   cli.StringFlag{
    Name: "bpm",
    Value: "400.00",
    Usage: "BPM. The game's tempo.",
  },
   cli.StringFlag{
    Name: "text",
    Value: "hello world.",
    Usage: "Text that you want to type",
  },
   cli.StringFlag{
    Name: "skipIntro",
    Value: "false",
    Usage: "Use 'true' when you want to skip intro",
  },
   cli.StringFlag{
    Name: "skipChecking",
    Value: "false",
    Usage: "Use 'true' when you want to skip checking that you were rushing or dragging.",
  },

   cli.StringFlag{
    Name: "level",
    Value: "normal",
    Usage: "easy|normal|hard|guru",
  },

}
    app.Name = "The tempo game. Whiplash."
    app.Usage = "A parody of the movie Whiplash"
    app.Version = "0.1.0"
    app.Author = "https://github.com/dorajistyle(JoongSeob Vito Kim)"
    app.Action = func(c *cli.Context) {
        game.Whiplash(c.Float64("bpm"),
                      c.String("text"),
                      c.Bool("skipIntro"),
                      c.Bool("skipChecking"), 
                      c.String("level"))
    }
    app.Run(os.Args)
}
