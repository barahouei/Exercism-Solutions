package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	pattern := `(^\[TRC])|(^\[DBG])|(^\[INF])|(^\[WRN])|(^\[ERR])|(^\[FTL])`

	r := regexp.MustCompile(pattern)

	return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	pattern := `<(=*-*~*\**)*>`

	r := regexp.MustCompile(pattern)

	return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `".*(?i)password.*"`

	r := regexp.MustCompile(pattern)

	counter := 0

	for _, v := range lines {
		if r.MatchString(v) {
			counter++
		}
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	pattern := `(end-of-line)+\d*`

	r := regexp.MustCompile(pattern)

	return r.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `(User\s+)((?i)[a-z]*\d*)`

	r := regexp.MustCompile(pattern)

	newLines := []string{}

	for _, v := range lines {
		s := r.FindStringSubmatch(v)

		if s != nil {
			user := s[2]
			user = "[USR] " + user + " " + v

			newLines = append(newLines, user)
		} else {
			newLines = append(newLines, v)
		}
	}

	return newLines
}
