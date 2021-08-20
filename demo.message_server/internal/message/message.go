package message

import (
	"github.com/google/uuid"
	"github.com/vissong/ddd_think/demo.message_server/internal/message/ds"
	"github.com/vissong/ddd_think/demo.message_server/internal/message/entity"
)

type message struct {
	entity *entity.Entity
}

type Aggregation interface {
	Get() *entity.Entity
	QueryFrom() int
}

// New 创建一个新的聚合
func New(content string) Aggregation {
	return &message{entity: &entity.Entity{
		ID:      entity.MsgID(uuid.New().String()),
		Content: content,
	}}
}

// LoadByID 从ds中加载聚合
func LoadByID(id entity.MsgID) Aggregation {
	entity := ds.New().Get(id)
	return &message{entity: entity}
}
