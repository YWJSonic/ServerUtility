package messagehandle

import (
	"fmt"
	"time"
)

var (
	// IsAddTimeFlag log title add time flag
	IsAddTimeFlag = true
	// IsPrintTipLog Tip log
	IsPrintTipLog = true
	// IsPrintLog Debug log
	IsPrintLog = true
	// IsPrintErrorLog Error log
	IsPrintErrorLog = true
	// TimeFormat date time print format
	TimeFormat = "2006-01-02 15:04:05Z07:00" //time.stdLongYear + '-' + time.stdZeroMonth + '-' + time.stdZeroDay + ' ' +
)

// log Tag
const (
	Log      = "Log"
	ErrorLog = "Error"
)

// TipLogPirnt ...
func TipLogPirnt(msg string) {
	print(Log, msg)
}

// TipPrintf ...
func TipPrintf(msg string, a ...interface{}) {
	printf(Log, msg, a...)

}

// TipPrintln ...
func TipPrintln(msg string, a ...interface{}) {
	println(Log, msg, a...)

}

// LogPrint ...
func LogPrint(msg string) {

	if !IsPrintLog {
		return
	}

	print(Log, msg)
}

// LogPrintf ...
func LogPrintf(msg string, a ...interface{}) {

	if !IsPrintLog {
		return
	}

	printf(Log, msg, a...)
}

// LogPrintln ...
func LogPrintln(msg string, a ...interface{}) {

	if !IsPrintLog {
		return
	}
	println(Log, msg, a...)
}

// ErrorLogPrint ...
func ErrorLogPrint(msg string) {
	print(ErrorLog, msg)
}

// ErrorLogPrintf ...
func ErrorLogPrintf(msg string, a ...interface{}) {
	printf(ErrorLog, msg, a...)
}

// ErrorLogPrintln ...
func ErrorLogPrintln(msg string, a ...interface{}) {
	println(ErrorLog, msg, a...)
}

func print(logtype, msg string, a ...interface{}) {
	if IsAddTimeFlag {
		fmt.Printf("%s %s %s", time.Now().Format(TimeFormat), logtype, msg)
	} else {
		fmt.Printf("%s %s", logtype, msg)
	}
}
func printf(logtype, msg string, a ...interface{}) {
	if IsAddTimeFlag {
		msg = fmt.Sprintf("%s %s %s", time.Now().Format(TimeFormat), logtype, msg)
	} else {
		msg = fmt.Sprintf("%s %s", logtype, msg)
	}
	fmt.Printf(msg, a...)
}
func println(logtype, msg string, a ...interface{}) {
	var tmp []interface{}
	if IsAddTimeFlag {
		tmp = append(tmp, time.Now().Format(TimeFormat))
	}
	tmp = append(tmp, logtype)
	tmp = append(tmp, msg)
	tmp = append(tmp, a...)
	fmt.Println(tmp...)
}
