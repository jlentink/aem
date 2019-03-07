package main

import (
	"fmt"
	"strings"
)

type TableWriter struct {
	tableRaw string
}

func (w *TableWriter) Write(p []byte) (int, error) {
	w.tableRaw = w.tableRaw + string(p)
	return len(p), nil
}

func (w *TableWriter) getTables() []string {
	tables := make([]string, 0)

	lines := strings.Split(w.tableRaw, "\n")
	table := ""
	for _, line := range lines {
		if line == "" {
			tables = append(tables, fmt.Sprintf("%s", table))
			table = ""
		} else {
			table = fmt.Sprintf("%s%s\n", table, line)
		}
	}

	return tables
}
