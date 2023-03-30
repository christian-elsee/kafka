package cmd

import (
  "testing"

  "github.com/matryer/is"
)

func Test(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")

  is := is.New(t)
  t.Logf("What we want :: is=%s", "NoErr")

  err := Execute()
  t.Logf("What we have :: is=%s", err)

  is.NoErr(err)
}
