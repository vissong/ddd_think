package ds

import (
	"github.com/vissong/ddd_think/demo.message_server/internal/message/entity"
)

type MockDS struct {
}

func (m MockDS) Save(msg *entity.Entity) entity.MsgID {
	panic("implement me")
}

func (m MockDS) Get(id entity.MsgID) *entity.Entity {
	panic("implement me")
}

func (m MockDS) List(num int) []*entity.Entity {
	panic("implement me")
}
