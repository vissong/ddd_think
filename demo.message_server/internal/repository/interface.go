package repository

import (
	"github.com/vissong/ddd_think/demo.message_server/internal/entity"
)

// IRepo 存储的接口，声明了存储需要提供的接口
// 当领域中不存在 data_proxy 时候，该接口提供的能力与 data proxy 一致
// 只不过 data proxy 通信是基于 pb 的协议，而 IRepo 的通信，是基于接口的签名
// 接口签名上，只能够设计到领域模型中的实体，值对象，聚合等，不应该将存储层面的对象（DO）传到到外部
type IRepo interface {
	Save(msg *entity.Message) entity.MsgID
	Get(id entity.MsgID) *entity.Message
	List(num int) []*entity.Message
}
