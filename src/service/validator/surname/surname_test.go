package surname

import (
	"testing"
)

var dataProvider = []struct {
	value  string
	result bool
}{
	{
		value:  "Smith",
		result: true,
	},
	{
		value:  "Surname-",
		result: true,
	},
	{
		value:  "Иванов",
		result: true,
	},
	{
		value:  "Абдурахман-ибн-хаттаб",
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
		value:  "",
		result: false,
	},
}

func TestSurnameValidator(t *testing.T) {
	t.Run("Validate surname", func(t *testing.T) {
		for _, entity := range dataProvider {
			if Validate(entity.value) != entity.result {
				t.Error("Error in surname validation")
			}
		}
	})
}
