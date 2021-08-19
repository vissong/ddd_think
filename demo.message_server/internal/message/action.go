package message

// QueryFrom 查询消息来源
func (m message) QueryFrom() int {
	return 1
}

func (m message) Get() *Entity {
	return m.entity
}
