package birthday

import (
	"github.com/joyt/godate"
	"regexp"
	"time"
)

func Validate(value string) bool {
	if regexp.MustCompile(`(?mi)[0-9.\\/\sa-z]`).FindAllString(value, -1) == nil {
		return false
	}

	t, _, err := date.ParseAndGetLayout(value)

	if int32(t.Unix()) > int32(time.Now().Unix()) {
		return false
	}

	return err == nil
}
