package transaction

import (
	"fmt"

	"github.com/vissong/ddd_think/demo.message_server/internal/message"
	"github.com/vissong/ddd_think/demo.message_server/internal/message/ds"
	"github.com/vissong/ddd_think/demo.message_server/internal/message/entity"
)

// Save 保存消息
func Save(content string) entity.MsgID {

	oldMsg := message.LoadByID("abc")
	fmt.Println(oldMsg.Get())

	msg := message.New("test msg")
	return ds.New().Save(msg.Get())
}
