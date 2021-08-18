package port

type port struct {
}

func New() IPort {
	return port{}
}

func (p port) Verify() error {
	return nil
}
