package birthday

func Validate(value string) bool {
	return len(value) >= 8 && len(value) <= 10
}
