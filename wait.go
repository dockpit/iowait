package iowait

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"time"
)

//error that is returned when WaitForRegexp times out
type TimeoutError struct{ error }

func timeoutError(to time.Duration, ls []string) TimeoutError {
	return TimeoutError{fmt.Errorf("timed out waiting for match after %s, scanned lines: %s", to, ls)}
}

//blocks the routine until a line in the given reader matches the given regex, if
//this doesn't happen before the timeout expires an error is returned
func WaitForRegexp(r io.Reader, exp *regexp.Regexp, to time.Duration) error {
	found := make(chan bool)

	//scan in different routine and send on
	//channel when a match was found
	s := bufio.NewScanner(r)
	ls := []string{}
	go func() {
		for s.Scan() {
			ls = append(ls, s.Text())
			if exp.MatchString(s.Text()) {
				found <- true
			}
		}
	}()

	select {
	case <-time.After(to):
		return timeoutError(to, ls)
	case <-found:
	}

	return nil
}
