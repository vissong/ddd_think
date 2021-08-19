## message 

`internal/message` 代表的是 message 这一个领域对象。所有与 message 这个领域对象相关的组件，都放在这个包下。

### `message` 包相关文件描述：

`message.go` 的主要工作，定义 message 的聚合，定义 message 的 repo 接口，定义 message 的 interaction 接口

`ds` 目录
- data source 目录，主要是实现 message 的 repo 接口，不允许 ds 内的细节，对象等暴露到 message 中。
- data source 是实现，主要依赖 message 中的接口定义，实体类型定义
- data source 的入参与出参一般是实体，所以交给 transaction/interaction 来调用，而不能由实体来调用（禁止循环依赖，外部依赖内部原则）

`entity.go` 实现 message 的实体与方法，被聚合持有，并通过 interaction/ds 被 transaction 获取。

`action.go` 实现 message 的 inter action，用于对 message 这一聚合的操作

### REPO 接口原则

在 `message` 是聚合，持有 `entity` 的前提下，repo 接口的定义，根据实际情况可以是对实体的操作，也可以是对聚合的操作。
但是需要注意的是，需要注意避免由于接口设计的不合理，导致 ds 内部执行保存的时候，无法完成工作，如果出现这类问题，首先需要考虑的是修改接口的定义。
而不是打破架构的边界，从 ds 层中直接调用某个 rpc 请求去获取所需要的数据。

