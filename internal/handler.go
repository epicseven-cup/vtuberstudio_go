package internal

type Handler interface {
	handle() error
}
