package basic

import "github.com/4p00rv/pinch-ui/pkg/ui"

func BasicPage() string {
	page := ui.CreatePage()
	page.SetTitle("Example page")
	return page.ToString()
}
