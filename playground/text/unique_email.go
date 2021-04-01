package text

import "strings"

// https://leetcode.com/problems/unique-email-addresses/
func numUniqueEmails(emails []string) int {
	table := make(map[string]bool)

	for _, email := range emails {
		atSignIndex := strings.Index(email, "@")
		localName := strings.Replace(email[:atSignIndex], ".", "", -1)
		plusSignIndex := strings.Index(localName, "+")
		if plusSignIndex > -1 {
			localName = localName[:plusSignIndex]
		}
		domainName := email[atSignIndex:]
		table[localName+domainName] = true
	}

	return len(table)
}
