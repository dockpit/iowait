package iowait

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"time"
)

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
		return fmt.Errorf("timed out waiting for match after %s, scanned lines: %s", to, ls)
	case <-found:
	}

	return nil
}
