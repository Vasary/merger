package patronymic

import "regexp"

func Validate(value string) bool  {
	return regexp.MustCompile(`(?mi)[^a-zа-я]`).FindAllString(value, -1) == nil && value != ""
}