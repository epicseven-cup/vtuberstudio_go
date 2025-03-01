package internal

type Handler interface {
	Handle() error
}
