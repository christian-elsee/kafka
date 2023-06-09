package pkg

import (
	"fmt"
)

// constants ////////////////////////////////////

const PROJECT string = "github.com/christian-elsee/kafka"

// func /////////////////////////////////////////

func Trace(function, pkg string) string {
  return fmt.Sprintf(
    "%s/%s#%s", PROJECT, pkg, function,
  )
}
