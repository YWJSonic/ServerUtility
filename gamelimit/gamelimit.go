package gamelimit

import "fmt"

// IsInTotalMoneyWinLimit if int limit return true
// 0 is no limit
func IsInTotalMoneyWinLimit(limit, betMoney, totalWin int64) bool {
	if limit == 0 {
		return true
	}
	if totalWin > limit {
		return false
	}
	return true
}

// IsInTotalBetRateWinLimit if int limit return true
// 0 is no limit
func IsInTotalBetRateWinLimit(limit, betMoney, totalWin int64) bool {
	if limit == 0 {
		return true
	}
	if (totalWin / betMoney) > limit {
		return false
	}
	return true
}

// IsPlayerDayWinInLimit ...
func IsPlayerDayWinInLimit(limit, playerTotalWin int64) bool {

	if playerTotalWin >= limit {
		panic(fmt.Sprintf("PlayerDayWin:%d", playerTotalWin))
	}

	return true
}
