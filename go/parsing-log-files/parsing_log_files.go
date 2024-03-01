package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(~|\*|=|-)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	re := regexp.MustCompile(`"(.*)(?i)password(.*)"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+[a-zA-Z0-9]*`)
	for i, line := range lines {
		sl := re.FindString(line)
		if sl != "" {
			sl = strings.Replace(sl, "User ", "", -1)
			sl = strings.Replace(sl, " ", "", -1)
			lines[i] = fmt.Sprintf("[USR] %s %s", sl, line)
		}
	}
	return lines
}
