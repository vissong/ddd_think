## message 

`internal/message` 代表的是 message 这一个领域对象。所有与 message 这个领域对象相关的组件，都放在这个包下。

### `message` 包相关文件描述：

`message.go` 的主要工作，定义 message 的实体，定义 message 的 repo 接口，定义 message 的 interaction 接口

`ds` 目录

- data source 目录，主要是实现 message 的 repo 接口，不允许 ds 内的细节，对象等暴露到 message 中。
- data source 是实现，主要依赖 message 中的接口定义，实体类型定义
- data source 的入参与出参一般是实体，所以交给 transaction/interaction 来调用，而不能由实体来调用（禁止循环依赖，外部依赖内部原则）

