package ds

import (
	"fmt"

	"github.com/vissong/ddd_think/demo.message_server/internal/message/entity"
)

type messageDS struct {
	client interface{} // may be mysql,redis,rpc
}

type DO struct {
	entity.Entity
}

func New() entity.Repository {
	return &messageDS{client: nil}
}

func (m messageDS) Save(msg *entity.Entity) entity.MsgID {
	fmt.Println("Save", msg)
	return "uuid"
}

func (m messageDS) Get(id entity.MsgID) *entity.Entity {
	fmt.Println("Get", id)
	return &entity.Entity{
		ID:      id,
		Content: "messageDS content",
	}
}

func (m messageDS) List(num int) []*entity.Entity {
	fmt.Println("List", num)
	var result []*entity.Entity
	result = append(result, &entity.Entity{
		ID:      "uuid",
		Content: "messageDS content",
	})
	return result
}
