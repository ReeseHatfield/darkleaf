package core

type Core struct {
	key string
}

func NewCore() *Core {
	return new(Core)
}

func (c *Core) GetKey() string {
	return c.key
}
