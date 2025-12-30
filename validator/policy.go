package validator

type Mode string

const (
	// 串行执行，遇到 fatal 可中断
	ModeSequential Mode = "sequential"

	// 并行执行，收集所有结果
	ModeParallel Mode = "parallel"
)

type Policy struct {
	Mode           Mode
	StopOnFatal    bool
	MaxFeedbackNum int // 0 = unlimited
}

// 默认策略
func DefaultPolicy() Policy {
	return Policy{
		Mode:        ModeSequential,
		StopOnFatal: true,
	}
}
