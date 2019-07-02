package gamelimit

// IsInTotalMoneyWinLimit if int limit return true
func IsInTotalMoneyWinLimit(limit, betMoney, totalWin int64) bool {
	if totalWin > limit {
		return false
	}
	return true
}

// IsInTotalBetRateWinLimit if int limit return true
func IsInTotalBetRateWinLimit(limit, betMoney, totalWin int64) bool {
	if (totalWin / betMoney) > limit {
		return false
	}
	return true
}
