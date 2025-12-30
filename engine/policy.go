package engine

// LoopPolicy 控制收敛策略
type LoopPolicy struct {
	MaxIteration int
	StopOnFatal  bool
}
