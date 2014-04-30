package html

import (
	"io"
	"strings"

	"code.google.com/p/go.net/html"
)

type Predicate func(*html.Node) bool

func ElementWithClass(el, class string) []Predicate {
	return []Predicate{Element(el), Class(class)}
}

func Element(name string) Predicate {
	name_lower := strings.ToLower(name)
	return func(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == name_lower }
}

func Class(name string) Predicate {
	name_lower := strings.ToLower(name)
	return func(n *html.Node) bool { return hasClass(name_lower, n.Attr) }
}

func hasClass(name string, attrs []html.Attribute) bool {
	if val, ok := getAttr("class", attrs); ok {
		for _, v := range strings.Split(val, " ") {
			if strings.ToLower(v) == name {
				return true
			}
		}
	}
	return false
}

func getAttr(name string, attrs []html.Attribute) (string, bool) {
	for _, a := range attrs {
		if a.Key == name {
			return a.Val, true
		}
	}
	return "", false
}

type Transform func(*html.Node) (string, error)

func GetAllText() Transform {
	return func(n *html.Node) (string, error) {
		if n.Type == html.TextNode {
			return n.Data, nil
		}
		return "", nil
	}
}

func GetAllLinks() Transform {
	return func(n *html.Node) (string, error) {
		if n.Type == html.ElementNode {
			if h, ok := getAttr("href", n.Attr); ok {
				return h, nil
			}
		}
		return "", nil
	}
}

type Transformer struct {
	predicates []Predicate
	transform  Transform
	separator  string
}

func NewTransformer(preds []Predicate, xf Transform) *Transformer {
	return &Transformer{preds, xf, "\n"}
}

func (t *Transformer) Transform(r io.Reader) (string, error) {
	node, err := html.Parse(r)
	if err != nil {
		return "", err
	}
	res, err := t.processNodes(node, false)
	return strings.Join(res, t.separator), err
}

func (t *Transformer) processNodes(root *html.Node, match bool) ([]string, error) {
	var ret []string
	m := match || all(t.predicates, root)
	if m {
		r, err := t.transform(root)
		if err != nil {
			return ret, err
		}
		if r != "" {
			ret = append(ret, r)
		}
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		r, e := t.processNodes(c, m)
		if e != nil {
			return ret, e
		}
		ret = append(ret, r...)
	}
	return ret, nil
}

func all(p []Predicate, n *html.Node) bool {
	for _, pred := range p {
		if !pred(n) {
			return false
		}
	}
	return true
}
