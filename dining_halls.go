package bananas

// DiningHall represents a dining hall
// according to UCSC's nutritional website.
type DiningHall struct {
	LocationName   string
	LocationNumber string
}

var (
	// NineTen is the Nine & Ten Dining Hall
	NineTen = DiningHall{
		LocationName:   "Colleges Nine & Ten Dining Hall",
		LocationNumber: "40",
	}
	// CowellStevenson is the Cowell Stevenson Dining Hall
	CowellStevenson = DiningHall{
		LocationName:   "Cowell Stevenson Dining Hall",
		LocationNumber: "05",
	}
	// CrownMerrill is the Crown Merrill Dining Hall
	CrownMerrill = DiningHall{
		LocationName:   "Crown Merrill Dining Hall",
		LocationNumber: "20",
	}
	// PorterKresge is the Porter Kresge Dining Hall
	PorterKresge = DiningHall{
		LocationName:   "Porter Kresge Dining Hall",
		LocationNumber: "25",
	}
	// CarsonOakes is the Rachel Carson Oakes Dining Hall
	CarsonOakes = DiningHall{
		LocationName:   "Rachel Carson Oakes Dining Hall",
		LocationNumber: "30",
	}
)
