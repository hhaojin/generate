package Commands

var Commands map[string]ICommand

func init() {
	Commands = make(map[string]ICommand)
}

type CommandSet struct{}

func NewCommandSet() *CommandSet {
	return &CommandSet{}
}

func (this *CommandSet) Register(name string, command ICommand) {
	Commands[name] = command
}

func (this *CommandSet) Each() {
	for _, v := range Commands {
		v.Init()
	}
}

func (this *CommandSet) Run() {
	for _, v := range Commands {
		v.Run()
	}
}
