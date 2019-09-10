package bananas

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const baseNutritionMenuLink = "https://nutrition.sa.ucsc.edu/longmenu.aspx"
const baseNutritionLink = "https://nutrition.sa.ucsc.edu"

func generateURL(dh DiningHall, ml Meal, date time.Time) (string, error) {
	// Make a URL.
	baseMenuURL, err := url.Parse(baseNutritionMenuLink)
	if err != nil {
		return "", fmt.Errorf("could not generate URL: %v", err)
	}
	// See the documentation for (time.Time).Format for more info
	// on this var.
	referenceDate := "01/02/2006"

	// Add parameters to the URL.
	//
	// These are all the standard parameters
	// that my web browser generates when visiting the
	// menu.
	//
	// For ease of reading, I've decoded the URL values.
	//
	// Any parameter values that are hard-coded shouldn't
	// (in my testing, anyways) change across calls.
	parameters := url.Values{}
	parameters.Add("sName", "UC Santa Cruz Dining")
	parameters.Add("locationNum", dh.LocationNumber)
	parameters.Add("locationName", dh.LocationName)
	parameters.Add("naFlag", "1")
	parameters.Add("WeeksMenus", "UCSC - This Week's Menus")
	parameters.Add("dtdate", date.Format(referenceDate))
	parameters.Add("mealName", string(ml))
	baseMenuURL.RawQuery = parameters.Encode()

	return baseMenuURL.String(), nil
}

func generateRequest(menuURL string) (*http.Request, error) {
	// Grab the cookies necessary for the request.
	responseWithCookies, err := http.Get(baseNutritionLink)
	if err != nil {
		return nil, fmt.Errorf("failed to get necessary cookies: %v", err)
	}
	// Make a request...
	req, err := http.NewRequest(http.MethodGet, menuURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	// and add the cookies to it.
	for _, ck := range responseWithCookies.Cookies() {
		req.AddCookie(ck)
	}
	return req, nil
}

func menuFromRequest(req *http.Request) ([]string, error) {
	// Do the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to get menu: %v", err)
	}
	// Parse the request
	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("unable to parse response: %v", err)
	}
	// Find all a tags that have a parent with a class name of longmenucoldispname.
	menuSelections := document.Find(".longmenucoldispname a")
	// Make a slice that has a length equal to the number of a tags we found.
	// And add all the a tags text to that slice.
	tMenuItems := make([]string, menuSelections.Length())
	menuSelections.Each(func(index int, selection *goquery.Selection) {
		tMenuItems[index] = selection.Text()
	})
	return tMenuItems, nil
}

// MenuFor returns a slice of menu items given
// a dining hall, a meal, and a date.
func MenuFor(dh DiningHall, ml Meal, date time.Time) ([]string, error) {
	// Generate the URL
	menuURL, err := generateURL(dh, ml, date)
	if err != nil {
		return nil, err
	}
	// Generate the request.
	req, err := generateRequest(menuURL)
	if err != nil {
		return nil, err
	}

	// Return a menu from that request.
	return menuFromRequest(req)
}
