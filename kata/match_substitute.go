package kata

import (
	"regexp"
	"strings"
)

// Some Regex abomination
func Change(s string, prog string, version string) string {
	lines := strings.Split(s, "\n")
	var output string

	programRegex := regexp.MustCompile("\\stitle:\\s.+$")
	output += programRegex.ReplaceAllString(lines[0], ": "+prog+" ")

	authorRegex := regexp.MustCompile(`Author:\s(.+$)`)
	oldAuthor := authorRegex.FindStringSubmatch(lines[1])[1]
	output += strings.Replace(lines[1], oldAuthor, "g964 ", -1)

	telRegex := regexp.MustCompile("\\+1-\\d{3}-\\d{3}-\\d{4}")
	versionRegex := regexp.MustCompile(`\s\d+\.\d+$`)

	if !telRegex.MatchString(lines[3]) || !versionRegex.MatchString(lines[5]) {
		return "ERROR: VERSION or PHONE"
	}
	output += telRegex.ReplaceAllString(lines[3], "+1-503-555-0090 ")

	dateRegex := regexp.MustCompile(`\w*\s\w+\s\d+,\s\d{4}$`)
	output += dateRegex.ReplaceAllString(lines[4], "2019-01-01 ")

	if !(versionRegex.FindString(lines[5]) == " 2.0") {
		output += versionRegex.ReplaceAllString(lines[5], " "+version)
	} else {
		output += lines[5]
	}
	return output
}
