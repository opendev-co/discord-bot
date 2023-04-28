package discord

import (
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ValidGithubLink(matches []string) bool {
	return len(matches) > 1
}

func GithubLinkMatch(link string) []string {
	reg := regexp.MustCompile(`http(?:s|)://github\.com/(.*?/.*?)/blob/(.*?/.*?)#L([0-9]+)(?:C([0-9]+))?(?:-L([0-9]+))?(?:C([0-9]+))?`)
	match := reg.FindStringSubmatch(link)

	return match
}

func FormatGithubLines(match []string) (string, bool) {
	url := "https://raw.githubusercontent.com/" + match[1] + "/" + match[2]

	start, err := strconv.Atoi(match[3])
	if err != nil {
		return "", false
	}

	end, err := strconv.Atoi(match[5])
	if err != nil {
		end = start + 4
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", false
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	lines := strings.Split(string(body), "\n")

	if len(lines) < 2 {
		return "", false
	}
	if start > end {
		end = start + 4
	}
	if start > len(lines) {
		start = 0
	}
	if end > len(lines) {
		end = start + 4
	}

	lines = lines[start : end+1]

	extension := filepath.Ext(match[2])[1:]
	characters := 0

	message := "```" + extension + "\n"

	for i, line := range lines {
		characters += len(line)

		if characters > 1000 {
			message += "\n\nLimit of characters reached, try to select in 2 parts the lines"
			break
		}

		diff := len(strconv.Itoa(end)) - len(strconv.Itoa(start+i))
		space := strings.Repeat(" ", diff)

		message += "\n[>" + strconv.Itoa(start+i) + "] " + space + line
	}

	message += "\n```"
	return message, true
}
