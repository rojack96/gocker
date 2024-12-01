package gocker

import (
	"gocker/commands"
	"gocker/helpers"
)

type Compose struct {
	Command string
}

// Options

func (c *Compose) FileName(fileName string) *Compose {
	c.Command += " -f" + helpers.AppendString(fileName)

	return c
}

func (c *Compose) EnvFile(fileName string) *Compose {
	c.Command += " --env-file" + helpers.AppendString(fileName)
	return c
}

/* Commands */

func (c *Compose) Build() *commands.Build {
	return &commands.Build{Command: c.Command + " build"}
}

func (c *Compose) Up() *commands.Up {
	return &commands.Up{Command: c.Command + " up"}
}

/*func (c *Compose) Exec() {
	return
}*/
