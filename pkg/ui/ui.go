package ui

type Page struct {
	title string
}

func CreatePage() Page {
	page := Page{}
	page.title = "Default page"
	return page
}

func (p *Page) SetTitle(t string) {
	p.title = t
}

func (p Page) ToString() string {
	return "<html>Testing</html>"
}
