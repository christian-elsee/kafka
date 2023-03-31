package pkg

import (
	"testing"
	"regexp"
)

func TestTrace(t *testing.T) {
	t.Logf("Enter")
	defer t.Logf("Exit")

	want := regexp.MustCompilePOSIX(`pkg/test_trace.go#TestTrace`)
	t.Logf("What we want :: is=%s", "pkg/test_trace.go#TestTrace")

	have := Trace("TestTrace", "pkg/test_trace.go")
	t.Logf("What we have :: is=%s", have)

	if !want.MatchString(have) {
		t.Fatalf(
			"What we want does not match what we have :: want=%s have=%s", 
			want.String(), 
			have,
		)
	}

}
