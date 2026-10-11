// journal_spend sums "Spend" table entries from Obsidian daily journal notes
// (named YYYY-MM-DD.md) over a given week, month, or year, and separately
// totals the subset of entries whose "What" column contains a given
// substring (defaults to "Amazon"). Periods: week, month, year, decade.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const defaultJournalDir = "/Users/ramesh/Documents/obsidian/2-areas/journal/daily"

var (
	fileNameRE  = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})\.md$`)
	headingRE   = regexp.MustCompile(`^#{1,6}\s`)
	spendRE     = regexp.MustCompile(`(?i)^##\s+Spend\s*$`)
	separatorRE = regexp.MustCompile(`^:?-+:?$`)
)

type entry struct {
	date  time.Time
	what  string
	price float64
}

func main() {
	period := flag.String("period", "", "aggregation period: week, month, year, or decade (required)")
	dateStr := flag.String("date", "", "reference date YYYY-MM-DD within the period (default: today)")
	dir := flag.String("dir", defaultJournalDir, "directory containing YYYY-MM-DD.md journal files")
	filter := flag.String("filter", "Amazon", "case-insensitive substring to match against the 'What' column")
	flag.Parse()

	p := strings.ToLower(strings.TrimSpace(*period))
	if p != "week" && p != "month" && p != "year" && p != "decade" {
		fmt.Fprintln(os.Stderr, "error: -period must be one of: week, month, year, decade")
		flag.Usage()
		os.Exit(1)
	}

	ref := time.Now()
	if *dateStr != "" {
		t, err := time.Parse("2006-01-02", *dateStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: invalid -date %q: %v\n", *dateStr, err)
			os.Exit(1)
		}
		ref = t
	}
	// Normalize to a UTC midnight so it compares cleanly against entry dates,
	// which are parsed from filenames with no time-of-day component.
	ref = time.Date(ref.Year(), ref.Month(), ref.Day(), 0, 0, 0, 0, time.UTC)

	entries, err := loadEntries(*dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading journal entries: %v\n", err)
		os.Exit(1)
	}

	start, end := periodBounds(p, ref)
	inPeriod := entries[:0:0]
	for _, e := range entries {
		if !e.date.Before(start) && !e.date.After(end) {
			inPeriod = append(inPeriod, e)
		}
	}
	sort.Slice(inPeriod, func(i, j int) bool { return inPeriod[i].date.Before(inPeriod[j].date) })

	needle := strings.ToLower(strings.TrimSpace(*filter))
	var total, filteredTotal float64
	var filteredEntries []entry
	for _, e := range inPeriod {
		total += e.price
		if needle != "" && strings.Contains(strings.ToLower(e.what), needle) {
			filteredTotal += e.price
			filteredEntries = append(filteredEntries, e)
		}
	}

	fmt.Printf("Period: %s (%s to %s)\n", p, start.Format("2006-01-02"), end.Format("2006-01-02"))
	fmt.Printf("Entries: %d\n", len(inPeriod))
	fmt.Printf("Total spend: %.2f\n", total)
	fmt.Printf("%q spend: %.2f (%d entries)\n", *filter, filteredTotal, len(filteredEntries))

	if len(filteredEntries) > 0 {
		fmt.Println()
		fmt.Printf("Matching %q entries:\n", *filter)
		for _, e := range filteredEntries {
			fmt.Printf("  %s  %-30s %8.2f\n", e.date.Format("2006-01-02"), e.what, e.price)
		}
	}

}

// periodBounds returns the inclusive [start, end] date range (UTC midnight)
// covering ref for the given period: "week" (Monday-Sunday), "month", "year",
// or "decade" (e.g. 2020-2029).
func periodBounds(period string, ref time.Time) (time.Time, time.Time) {
	switch period {
	case "week":
		daysSinceMonday := (int(ref.Weekday()) + 6) % 7 // Monday == 0
		start := ref.AddDate(0, 0, -daysSinceMonday)
		end := start.AddDate(0, 0, 6)
		return start, end
	case "month":
		start := time.Date(ref.Year(), ref.Month(), 1, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 1, -1)
		return start, end
	case "year":
		start := time.Date(ref.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(ref.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
		return start, end
	case "decade":
		first := ref.Year() - ref.Year()%10
		start := time.Date(first, time.January, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(first+9, time.December, 31, 0, 0, 0, 0, time.UTC)
		return start, end
	default:
		panic("unreachable: unknown period " + period)
	}
}

// loadEntries reads every YYYY-MM-DD.md file in dir and parses its "## Spend"
// table, if present.
func loadEntries(dir string) ([]entry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var all []entry
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		m := fileNameRE.FindStringSubmatch(f.Name())
		if m == nil {
			continue
		}
		date, err := time.Parse("2006-01-02", m[1])
		if err != nil {
			continue
		}

		path := filepath.Join(dir, f.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: cannot read %s: %v\n", path, err)
			continue
		}

		rows := parseSpendTable(string(data), f.Name())
		for _, r := range rows {
			all = append(all, entry{date: date, what: r.what, price: r.price})
		}
	}
	return all, nil
}

type row struct {
	what  string
	price float64
}

// parseSpendTable extracts the rows of the markdown table under the "##
// Spend" heading, up to (but not including) the next heading or EOF.
func parseSpendTable(content, fileName string) []row {
	lines := strings.Split(content, "\n")

	start := -1
	for i, line := range lines {
		if spendRE.MatchString(strings.TrimRight(line, "\r")) {
			start = i + 1
			break
		}
	}
	if start == -1 {
		return nil // no Spend section in this file
	}

	end := len(lines)
	for i := start; i < len(lines); i++ {
		if headingRE.MatchString(strings.TrimSpace(lines[i])) {
			end = i
			break
		}
	}

	var rows []row
	headerSeen := false
	for _, raw := range lines[start:end] {
		line := strings.TrimSpace(raw)
		if !strings.HasPrefix(line, "|") {
			continue
		}

		cells := splitTableRow(line)
		if len(cells) < 2 {
			continue
		}
		if isSeparatorRow(cells) {
			continue
		}
		if !headerSeen {
			headerSeen = true // this is the "| What | Price | ... |" header row
			continue
		}

		what := strings.TrimSpace(cells[0])
		priceStr := strings.TrimSpace(cells[1])
		if what == "" || priceStr == "" {
			continue // blank placeholder row
		}

		price, err := parsePrice(priceStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %s: skipping row %q with unparsable price %q\n", fileName, what, priceStr)
			continue
		}
		rows = append(rows, row{what: what, price: price})
	}
	return rows
}

// splitTableRow splits a "| a | b | c |" line into its trimmed cells,
// dropping the empty leading/trailing fields produced by the outer pipes.
func splitTableRow(line string) []string {
	parts := strings.Split(line, "|")
	if len(parts) > 0 && strings.TrimSpace(parts[0]) == "" {
		parts = parts[1:]
	}
	if len(parts) > 0 && strings.TrimSpace(parts[len(parts)-1]) == "" {
		parts = parts[:len(parts)-1]
	}
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}

// isSeparatorRow reports whether every cell looks like a markdown table
// separator cell, e.g. "---" or "---:" or ":---:".
func isSeparatorRow(cells []string) bool {
	for _, c := range cells {
		if c == "" || !separatorRE.MatchString(c) {
			return false
		}
	}
	return true
}

// parsePrice parses a price cell such as "2.99", "$1,234", or "1234" into a
// float64.
func parsePrice(s string) (float64, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "$")
	s = strings.ReplaceAll(s, ",", "")
	return strconv.ParseFloat(s, 64)
}
