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

func (m messageDS) Save(msg *message.Entity) message.MsgID {
	fmt.Println("Save", msg)
	return "uuid"
}

func (m messageDS) Get(id message.MsgID) *message.Entity {
	fmt.Println("Get", id)
	return &message.Entity{
		ID:      id,
		Content: "messageDS content",
	}
}

func (m messageDS) List(num int) []*message.Entity {
	fmt.Println("List", num)
	var result []*message.Entity
	result = append(result, &message.Entity{
		ID:      "uuid",
		Content: "messageDS content",
	})
	return result
}
