package engine

// State 是业务无关的不透明上下文
type State interface {
	Value() any
}
