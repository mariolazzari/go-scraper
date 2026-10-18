package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	// sort key
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// sort pages
	sorted := make(map[string]PageData, len(pages))
	for _, key := range keys {
		sorted[key] = pages[key]
	}

	// marshal json
	sortedJson, err := json.MarshalIndent(&sorted, "", "  ")
	if err != nil {
		return err
	}

	// write file
	err = os.WriteFile(filename, []byte(sortedJson), 0666)
	if err != nil {
		return err
	}

	return nil
}
