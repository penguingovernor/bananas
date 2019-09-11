package bananas

import (
	"fmt"
	"net/http"
	"time"
)

// Client is meant to be used
// when multiple menus want to be accessed.
//
// It is preferred over MenuFrom because
// it will avoid making a call to the main
// UCSC nutritional page for cookies.
type Client struct {
	cookies []*http.Cookie
}

func (c *Client) updateCookies() error {
	// Grab the cookies necessary for the request.
	responseWithCookies, err := http.Get(baseNutritionLink)
	if err != nil {
		return fmt.Errorf("failed to get necessary cookies: %v", err)
	}
	// Copy over the cookies.
	goodCookies := responseWithCookies.Cookies()
	c.cookies = make([]*http.Cookie, len(goodCookies))
	copy(c.cookies, goodCookies)
	return nil
}

func (c Client) needsUpdate() bool {
	// For right now, the nutrition
	// website only sets session cookies.
	// Therefore they are good until
	// we "close" a session.
	//
	// Since we're storing these
	// cookies in memory we're never
	// "closing" the session.
	//
	// If UCSC updates the website
	// to require non-session cookies
	// then this is where we'll do more
	// checks.
	return c.cookies == nil
}

// MenuFor is equivalent to the standalone MenuFor function.
//
// The primary difference between the two is that the one that
// uses a client maintains cookies across calls and only updates them
// when need be.
func (c *Client) MenuFor(dh DiningHall, ml Meal, date time.Time) ([]string, error) {
	// Generate the URL
	menuURL, err := generateURL(dh, ml, date)
	if err != nil {
		return nil, err
	}
	// Make a request...
	req, err := http.NewRequest(http.MethodGet, menuURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	// Check if we need to update our cookies.
	// And update them if need be.
	if c.needsUpdate() {
		if err := c.updateCookies(); err != nil {
			return nil, err
		}
	}
	// Add our cookies to it.
	for _, ck := range c.cookies {
		req.AddCookie(ck)
	}
	return menuFromRequest(req)
}
