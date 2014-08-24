package utilities

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}

func Pop(strings []string) (string, []string) {
	popped, strings := strings[0], strings[1:len(strings)]
	return popped, strings
}
