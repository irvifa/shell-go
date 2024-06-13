package shell

type Command interface {
	Execute(args []string) bool
}
