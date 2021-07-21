package main

import "fmt"

type Logger struct {
	TimeoutSeconds int
	LogInfo        map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{
		LogInfo: make(map[string]int),
	}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	lastTimestamp, logEntryExists := this.LogInfo[message]

	// log is new, OR log is not new but the timestamp threshold is passed
	if !logEntryExists {
		this.LogInfo[message] = timestamp
		return true
	} else if timestamp-lastTimestamp >= 10 {
		this.LogInfo[message] = timestamp
		return true
	}

	// log is not new and the threshold has not been reached
	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

func main() {
	type Log struct {
		Timestamp int
		Message   string
	}

	cases := []struct {
		LogEntry []Log
		Expected []bool
	}{
		// case 0
		{
			LogEntry: []Log{
				Log{1, "foo"},
				Log{2, "bar"},
				Log{3, "foo"},
				Log{8, "bar"},
				Log{10, "foo"},
				Log{11, "foo"},
			},

			Expected: []bool{true, true, false, false, false, true},
		},
	}

	for _, c := range cases {
		logger := Constructor()

		for i := 0; i < len(c.LogEntry); i++ {
			validMessage := logger.ShouldPrintMessage(c.LogEntry[i].Timestamp, c.LogEntry[i].Message)
			expectedOutcome := c.Expected[i]

			if validMessage != expectedOutcome {
				panic(fmt.Errorf("not equal, log #%d", i))
			}
		}
	}
}
