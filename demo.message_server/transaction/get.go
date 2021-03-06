package transaction

import (
	"github.com/vissong/ddd_think/demo.message_server/internal/entity"
	"github.com/vissong/ddd_think/demo.message_server/internal/repository"
)

// Save 保存消息
func Save(content string) entity.MsgID {
	msg := entity.New("test msg")
	return repository.New().Save(msg)
}
