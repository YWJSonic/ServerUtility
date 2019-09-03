package gamelimit

import (
	"os"

	"gitlab.com/ServerUtility/messagehandle"
)

// ServerDayPayLimit All player total can win score
var ServerDayPayLimit int64 = 15000000

// ServerDayPayDefault refresh default value
var ServerDayPayDefault int64

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

// IsServerDayPayInLimit if over limit shutdown server
func IsServerDayPayInLimit(allPlayerWinScore int64) bool {
	if allPlayerWinScore > ServerDayPayLimit {
		messagehandle.LogPrintf("AllPlayerDayWin:%d", allPlayerWinScore)
		os.Exit(0)
	}

	return true
}
