package passwords

import (
	"regexp"
	"testing"
)

func TestCreatePasswords(t *testing.T) {
	t.Run("test create passwords", func(t *testing.T) {
		c := &GeneratePasswords{}

		got := c.Generate()

		match, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", got)

		if !match {
			t.Errorf("This password match the criteria")
		}
	},
	)
}
