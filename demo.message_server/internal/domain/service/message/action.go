package message

import (
	"github.com/vissong/ddd_think/demo.message_server/internal/message/entity"
)

// QueryFrom 查询消息来源
func (m message) QueryFrom() int {
	return 1
}

func (m message) Get() *entity.Entity {
	return m.entity
}
