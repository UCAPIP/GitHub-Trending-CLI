package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintTable(repos []Repo) {
	if len(repos) == 0 {
		fmt.Println("No trending repositories found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "STARS\tLANGUAGE\tREPOSITORY\tDESCRIPTION")
	fmt.Fprintln(w, "-----\t--------\t----------\t-----------")

	for _, r := range repos {
		desc := r.Description
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		lang := r.Language
		if lang == "" {
			lang = "—"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", r.Stars, lang, r.FullName, desc)
	}
	w.Flush()
}
