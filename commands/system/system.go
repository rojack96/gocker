package system

type System struct {
	command string
}

func New(cmd string) *System {
	return &System{command: cmd}
}
