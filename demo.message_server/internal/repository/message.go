package repository

import (
	"fmt"

	"github.com/vissong/ddd_think/demo.message_server/internal/entity"
)

type message struct {
	client interface{}
}

func New() IRepo {
	return &message{client: nil}
}

func (m message) Save(msg *entity.Message) entity.MsgID {
	fmt.Println("Save", msg)
	return "uuid"
}

func (m message) Get(id entity.MsgID) *entity.Message {
	fmt.Println("Get", id)
	return &entity.Message{
		ID:      id,
		Content: "message content",
	}
}

func (m message) List(num int) []*entity.Message {
	fmt.Println("List", num)
	var result []*entity.Message
	result = append(result, &entity.Message{
		ID:      "uuid",
		Content: "message content",
	})
	return result
}
