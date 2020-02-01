package utils

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const StringsDoubleSpace = "  "
const StringsSingleSpace = " "

func CompressWhitespace(input string) string {
	mostlyRemoved := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				input,
				"\r", "",
			),
			"\t", StringsSingleSpace,
		),
		"\n", StringsSingleSpace,
	)
	for strings.Contains(mostlyRemoved, StringsDoubleSpace) {
		mostlyRemoved = strings.ReplaceAll(mostlyRemoved, StringsDoubleSpace, StringsSingleSpace)
	}
	return strings.Trim(mostlyRemoved, StringsSingleSpace)
}

func FormatMigrationName(input string) string {
	return strings.ReplaceAll(
		strings.ToLower(
			CompressWhitespace(
				input,
			),
		), " ", "_",
	)
}

func Hash(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}
