package name

import (
	"testing"
)

var dataProvider = []struct {
	value  string
	result bool
}{
	{
		value:  "John",
		result: true,
	},
	{
		value:  "John-",
		result: false,
	},
	{
		value:  "Иван",
		result: true,
	},
	{
		value:  "Tammy",
		result: true,
	},
	{
		value:  "Brown ",
		result: false,
	},
	{
		value:  "Hello Kitty",
		result: false,
	},
	{
		value:  "!@#$%^&*()1_ ",
		result: false,
	},
	{
		value:  "\n",
		result: false,
	},
	{
		value:  "\r",
		result: false,
	},
	{
		value:  "One two three",
		result: false,
	},
	{
		value:  " whitespace before",
		result: false,
	},
	{
		value:  "",
		result: false,
	},
}

func TestNameValidator(t *testing.T) {
	t.Run("Validate name", func(t *testing.T) {
		for _, entity := range dataProvider {
			if Validate(entity.value) != entity.result {
				t.Error("Error in name validation")
			}
		}
	})
}
