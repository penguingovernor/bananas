package bananas

import (
	"fmt"
	"net/url"
	"time"
)

const baseNutritionMenuLink = "https://nutrition.sa.ucsc.edu/longmenu.aspx"

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
