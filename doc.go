/*
Package bananas is a Go library for interfacing with the dining hall menus at the University of California Santa Cruz

Origins

At UCSC there are five (5) dining halls.
One for each pair of colleges in the ten college system.

Each dining hall has it's own menu that can be viewed online here: https://nutrition.sa.ucsc.edu
Many students have built third party menu viewers -- each with their own method of parsing the website
(manual web scraping is necessary as there is no public API).
And this is where bananas comes in.
Bananas will serve as the engine for a new API server that is in development, so that students can build custom web views, apps etc. .

Bananas will also be available for everyone, so that they too can build awesome stuff with it.

Usage

To use bananas simply call MenuFor with one item from each of the following lists:

Dining Halls (Choose one (1)):

	bananas.PorterKresge // Porter and Kresge Dining Hall
	bananas.NineTen // College Nine and Ten Dining Hall
	bananas.CowellStevenson // Cowell and Steventson Dining Hall
	bananas.CarsonOakes // Rachel Carson and Oakes Dining Hall
	bananas.CrownMerrill // Crown and Merrill Dining Hall

Meals (Choose one (1)):

	bananas.Breakfast // For breakfast
	bananas.Lunch // For lunch
	bananas.Dinner // For dinner
	bananas.LateNight // For late night

Time:

Any struct that is time.Time will satisfy this parameter.
Note that the actual time (e.g 12:00PM) is ignored and only the date is used.

Examples

Here are some common uses for bananas.

	// Get the lunch menu for today's meal at RCC and Oakes Dining Hall.
	banana.MenuFor(bananas.CarsonOakes, bananas.Lunch, time.Now())

	// For tomorrow's dinner menu at College 9 & 10.
	banana.MenuFor(bananas.NineTen, bananas.Dinner, time.Now().Add(24 * time.Hour))

Note that each of these calls will make two (2) http requests.
One for the menu data, and another to set the cookies.

If you would like to only have one http request per call use bananas.Client

	// Make a client.
	c := bananas.Client{}
	// Get the menu as normal.
	c.MenuFor(bananas.CarsonOakes, bananas.Lunch, time.Now())
*/
package bananas
