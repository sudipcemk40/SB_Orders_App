package main

import "strings"

func GetPDFName(no string) string {

	// Format 1:
	// SB-25/2015 -> SB25-2015.pdf

	if strings.Contains(no, "/") {

		parts := strings.Split(no, "/")

		if len(parts) != 2 {
			return ""
		}

		circularPart := strings.ReplaceAll(parts[0], "-", "")
		year := parts[1]

		return circularPart + "-" + year + ".pdf"
	}

	// Format 2:
	// SB25-2015 -> SB25-2015.pdf

	if strings.HasPrefix(no, "SB") && strings.Contains(no, "-") {
		return no + ".pdf"
	}

	return ""
}
