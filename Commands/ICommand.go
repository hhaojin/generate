package Commands

type ICommand interface {
	Init()
	Run()
}
