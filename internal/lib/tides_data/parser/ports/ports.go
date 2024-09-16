package tides_data_ports_parser

import (
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type Port struct {
	ExternalId int
	Name       string
}

type Link struct {
	Class string
	Id    string
	Href  string
	Text  string
}

func parseLink(n *html.Node, link *Link) {
	for _, a := range n.Attr {
		if a.Key == "class" {
			link.Class = a.Val
		}
		if a.Key == "id" {
			link.Id = a.Val
		}
		if a.Key == "href" {
			link.Href = a.Val
		}
	}
	var extractText func(n *html.Node)
	extractText = func(node *html.Node) {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.TextNode {
				link.Text = child.Data
				break
			}
			extractText(n)
		}
	}
	extractText(n)
}

func parse_port_link(n *html.Node, result *[]Port) {
	link := Link{}
	parseLink(n, &link)
	idInt, _ := strconv.Atoi(strings.TrimPrefix(link.Href, "/"))
	*result = append(*result, Port{
		ExternalId: idInt,
		Name:       link.Text,
	})
}

func is_port_link(n *html.Node) bool {
	result := false
	if n.Type == html.ElementNode && n.Data == "a" {
		linkData := struct {
			Class string
			Id    string
		}{}
		for _, a := range n.Attr {
			if a.Key == "class" {
				linkData.Class = a.Val
			}
			if a.Key == "id" {
				linkData.Id = a.Val
			}
		}
		portIdRegexp := regexp.MustCompile(`^Port\d+_\d+$`)

		if portIdRegexp.MatchString(linkData.Id) && strings.Contains(linkData.Class, "Port") {
			result = true
		}
	}
	return result
}

func parse_node(n *html.Node, result *[]Port) {
	// TODO : Here we should verify that the date of the page is the one requested
	if is_port_link(n) {
		parse_port_link(n, result)
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse_node(c, result)
	}
}

func ExtractPortsFromHtml(document *html.Node, ports *[]Port) {
	parse_node(document, ports)
}
