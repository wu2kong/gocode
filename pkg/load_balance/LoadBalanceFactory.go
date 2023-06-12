package load_balance

type LbType int

const (
	LbRandom           LbType = iota
	LbRoundRobin              // 轮询
	LbWeightRoundRobin        // 加权轮询
	LbConsistentHash          // 一致性哈希
)

func LoadBalanceFactory(lbType LbType) LoadBalance {
	switch lbType {
	case LbRandom:
		return &RandomBalance{}
	case LbConsistentHash:
		return NewConsistentHashBalance(10, nil)
	case LbRoundRobin:
		return &RoundRobinBalance{}
	case LbWeightRoundRobin:
		return &WeightRoundRobinBalance{}
	default:
		return &RandomBalance{}
	}
}
