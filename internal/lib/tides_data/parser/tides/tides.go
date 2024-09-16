package tides_data_tides_parser

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Tide struct {
	DaysOffset int
	High       bool
	Time       time.Duration
	Level      int
	Coef       *int
}

func extract_table_rows(node *html.Node, table_rows *[]*html.Node) {
	if node.Type == html.ElementNode && node.Data == "tr" {
		*table_rows = append(*table_rows, node)
		return
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		extract_table_rows(child, table_rows)
	}
}

func extract_text_nodes_contents(node *html.Node) (text_contents []struct {
	content string
	bold    bool
}) {
	var extract_text_contents func(node *html.Node)
	extract_text_contents = func(node *html.Node) {
		if node.Type == html.TextNode {
			text_content := struct {
				content string
				bold    bool
			}{
				content: strings.TrimSpace(node.Data),
				bold:    node.Parent.Data == "b",
			}

			text_contents = append(text_contents, text_content)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			extract_text_contents(child)
		}
	}
	extract_text_contents(node)
	return
}

func update_tide_with_text_node_content(tide *Tide, text_node_content string, column string) {
	switch column {
	case "Heure":
		timeRegext := regexp.MustCompile(`^(?P<hours>\d{2})h(?P<minutes>\d{2})$`)
		timeMatches := timeRegext.FindStringSubmatch(text_node_content)
		if len(timeMatches) == 0 {
			log.Fatalf("Invalid value for column Heure: %s", text_node_content)
		}
		hoursString := timeMatches[1]
		hoursInt, _ := strconv.Atoi(hoursString)
		minutesString := timeMatches[2]
		minutesInt, _ := strconv.Atoi(minutesString)

		tide.Time = (time.Duration(hoursInt) * time.Hour) + (time.Duration(minutesInt) * time.Minute)
	case "Hauteur":
		levelRegext := regexp.MustCompile(`(\d{1,2}),(\d{1,2})m`)
		levelMatches := levelRegext.FindStringSubmatch(text_node_content)
		if len(levelMatches) == 0 {
			log.Fatalf("Invalid value for column Level: %s", text_node_content)
		}

		levelMeters := levelMatches[1]
		levelCentimeters := levelMatches[2]
		tide.Level, _ = strconv.Atoi(levelMeters + levelCentimeters)
	case "Coeff":
		if text_node_content != "" {
			coef, _ := strconv.Atoi(text_node_content)
			tide.Coef = &coef
		}
	default:
		log.Fatalf("Unreconized header : %s", column)
	}
}

func parse_row(row *html.Node, headers *[]string) (tides []Tide) {
	tides = []Tide{}
	var day_offset *int
	for _, attribute := range row.Attr {
		if attribute.Key == "id" {
			sanitized_id, _ := strconv.Atoi(strings.TrimPrefix(attribute.Val, "MareeJours_"))
			day_offset = &sanitized_id
		}
	}

	column_index := 0
	for child := row.FirstChild; child != nil; child = child.NextSibling {
		columns := *headers
		column := columns[column_index]
		if child.Data == "td" {
			text_nodes := extract_text_nodes_contents(child)

			for index, text_node := range text_nodes {

				if len(tides) < (index + 1) {
					new_tide := Tide{DaysOffset: *day_offset, High: text_node.bold}
					update_tide_with_text_node_content(&new_tide, text_node.content, column)
					tides = append(tides, new_tide)
					continue
				}

				existing_tide := tides[index]
				update_tide_with_text_node_content(&existing_tide, text_node.content, column)
				tides[index] = existing_tide
			}
		}
		column_index++
	}

	return tides
}

func parse_header(row *html.Node) (headers []string) {
	for table_header := row.FirstChild; table_header != nil; table_header = table_header.NextSibling {
		text_content := table_header.FirstChild.Data
		headers = append(headers, strings.TrimSuffix(strings.TrimSpace(text_content), "."))
	}

	return
}

func parse_table_rows(table_row *html.Node, headers *[]string, result *[]Tide) {
	tide_rows := parse_row(table_row, headers)
	*result = append(*result, tide_rows...)
}

func extract_headers(table_rows []*html.Node) []string {
	for _, table_row := range table_rows {
		for _, attribute := range table_row.Attr {
			if attribute.Key == "class" && strings.Contains(attribute.Val, "MJE") {
				return parse_header(table_row)
			}
		}
	}

	return []string{}
}

func parse_tide_table(n *html.Node, result *[]Tide) {
	var table_rows []*html.Node

	extract_table_rows(n, &table_rows)

	headers := extract_headers(table_rows)

	for _, table_row := range table_rows {
		parse_table_rows(table_row, &headers, result)
	}
}

func is_tide_table(n *html.Node) bool {
	result := false
	if n.Type == html.ElementNode && n.Data == "table" {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "MareeJours" {
				result = true
				break
			}
		}
	}
	return result
}

func parse_node(n *html.Node, result *[]Tide) {
	// TODO : Here we should verify that the date of the page is the one requested
	if is_tide_table(n) {
		parse_tide_table(n, result)
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse_node(c, result)
	}
}

func ExtractTidesFromHtml(document *html.Node, tides *[]Tide) {
	parse_node(document, tides)
}
