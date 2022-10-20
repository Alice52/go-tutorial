package dpp

import "time"

// judgeRateByTime 报警速率决策函数
func judgeRateV1(now time.Time) int {
	switch hour := now.Hour(); {
	case hour >= 8 && hour < 20:
		return 10
	case hour >= 20 && hour <= 23:
		return 1
	}
	return -1
}

// judgeRate 报警速率决策函数
func judgeRateV0() int {
	now := time.Now() // 这个内置参数不好mock, 做成入参方便测试
	switch hour := now.Hour(); {
	case hour >= 8 && hour < 20:
		return 10
	case hour >= 20 && hour <= 23:
		return 1
	}
	return -1
}
