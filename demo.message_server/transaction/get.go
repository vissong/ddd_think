package transaction

import (
	"github.com/vissong/ddd_think/demo.message_server/internal/message"
	"github.com/vissong/ddd_think/demo.message_server/internal/message/ds"
)

// Save 保存消息
func Save(content string) message.MsgID {
	msg := message.New("test msg")
	return ds.New().Save(msg)
}
