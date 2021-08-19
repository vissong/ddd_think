package ds

import (
	"fmt"

	"github.com/vissong/ddd_think/demo.message_server/internal/message"
)

type messageDS struct {
	client interface{} // may be mysql,redis,rpc
}

func New() message.IRepo {
	return &messageDS{client: nil}
}

func (m messageDS) Save(msg *message.Message) message.MsgID {
	fmt.Println("Save", msg)
	return "uuid"
}

func (m messageDS) Get(id message.MsgID) *message.Message {
	fmt.Println("Get", id)
	return &message.Message{
		ID:      id,
		Content: "messageDS content",
	}
}

func (m messageDS) List(num int) []*message.Message {
	fmt.Println("List", num)
	var result []*message.Message
	result = append(result, &message.Message{
		ID:      "uuid",
		Content: "messageDS content",
	})
	return result
}
