package notetxt

import (
        "regexp"
        "strings"
)

func TitleToFilename (title string) string {
        title_clearer := regexp.MustCompile("[^a-zA-Z0-9\\s\\.\\-_]+")
        whitespace_clearer := regexp.MustCompile("\\s+")

        out := title_clearer.ReplaceAllString(title, "")
        out = strings.ToLower(out)
        out = whitespace_clearer.ReplaceAllString(out, " ")
        out = strings.Replace(out, " ", "-", -1)
        return out
}

type Note struct {
        name string
        filename string
        categories []string
}

