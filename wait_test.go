package iowait_test

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/dockpit/iowait"
	"github.com/stretchr/testify/assert"
)

func TestWaitForRegexpTimeout(t *testing.T) {
	err := iowait.WaitForRegexp(strings.NewReader(""), regexp.MustCompile(`.*serving on.*`), time.Millisecond*10)

	assert.IsType(t, iowait.TimeoutError{}, err)
	assert.Contains(t, err.Error(), "10ms")
	assert.Error(t, err)
}

func TestWaitForRegexp(t *testing.T) {
	err := iowait.WaitForRegexp(strings.NewReader("goji is serving on localhost"), regexp.MustCompile(`.*serving on.*`), time.Millisecond*10)
	assert.NoError(t, err)
}
