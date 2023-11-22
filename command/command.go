package command

// Command is a function that can be invoked by the `-c` flag
type Command func() error
