package limiter

func SetRule(action string, rule []LimitRule) {
	Limiter[action] = rule
}
