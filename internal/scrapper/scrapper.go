package scrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/kmg7/ncap/internal/config"
	"github.com/kmg7/ncap/internal/data"
)

func FetchAssessments() data.LatestAssessmentsSearchResult {
	result := data.LatestAssessmentsSearchResult{}
	response, err := http.Get(config.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func CarResultFromAssessment(asm data.Assessment) (*data.CarTest, error) {
	adultFrontal := true
	adultLateral := true
	url := "https://www.euroncap.com/en" + asm.URL
	ct := asm.ParseToCarTest()

	c := colly.NewCollector()
	// Set up a callback to find and process the safety-equipment-table
	c.OnHTML("#tab2 > div:nth-child(1) > div.table-body", func(e *colly.HTMLElement) {
		e.ForEach(".table-row:not(.title)", func(i int, row *colly.HTMLElement) {

			driverValue := row.ChildAttr(".table-row-c.c2 img", "data-src")
			passengerValue := row.ChildAttr(".table-row-c.c3 img", "data-src")
			rearValue := row.ChildAttr(".table-row-c.c4 img", "data-src")
			columns := data.SEColumns{
				Driver:    valueFromImages(driverValue),
				Passenger: valueFromImages(passengerValue),
				Rear:      valueFromImages(rearValue),
			}
			switch i {
			case 0:
				ct.SafetyEq.FrontalCP.FrontAirbag = columns
			case 1:
				ct.SafetyEq.FrontalCP.BeltPretensioner = columns
			case 2:
				ct.SafetyEq.FrontalCP.BeltLoadLimiter = columns
			case 3:
				ct.SafetyEq.FrontalCP.KneeAirbag = columns
			case 4:
				ct.SafetyEq.LateralCP.SideHeadAirbag = columns
			case 5:
				ct.SafetyEq.LateralCP.SideChestAirbag = columns
			case 6:
				ct.SafetyEq.LateralCP.SidePelvisAirbag = columns
			case 7:
				ct.SafetyEq.LateralCP.CentreAirbag = columns
			case 8:
				ct.SafetyEq.ChildP.IsofixISize = columns
			case 9:
				ct.SafetyEq.ChildP.IntegratedChildSeat = columns
			case 10:
				ct.SafetyEq.ChildP.AirbagCutOffSwitch = columns
			case 11:
				ct.SafetyEq.SeatBeltReminder = columns
			}
			// category := row.ChildText(".table-row-c.c1"
			// fmt.Printf("Category: %s\n", category)
			// fmt.Printf("Driver: %d, Passenger: %d, Rear: %d\n", driverNumericValue, passengerNumericValue, rearNumericValue)
			// fmt.Println(strings.Repeat("-", 30))
		})
	})
	// Other Safety Systems
	c.OnHTML("#tab2 > div:nth-child(2) > div.table-body", func(e *colly.HTMLElement) {
		e.ForEach(".table-row:not(.title)", func(i int, row *colly.HTMLElement) {
			imageValue := row.ChildAttr(".table-row-c.c2 img", "data-src")
			numericValue := valueFromImages(imageValue)

			switch i {
			case 0:
				ct.SafetyEq.ActiveBonnet = numericValue
			case 1:
				ct.SafetyEq.AEBVulnerableRoadUsers = numericValue
			case 2:
				ct.SafetyEq.AEBPedestrianReverse = numericValue
			case 3:
				ct.SafetyEq.AEBCarToCar = numericValue
			case 4:
				ct.SafetyEq.SpeedAssistance = numericValue
			case 5:
				ct.SafetyEq.LaneAssistSystem = numericValue
			}

			// category := row.ChildText(".table-row-c.c1")
			// fmt.Printf("Category: %s\n", category)
			// fmt.Printf("Value: %d\n", numericValue)
			// fmt.Println(strings.Repeat("-", 30))
		})
	})

	// Set up a collector for the second type of element
	c.OnHTML("div.frame-header", func(e *colly.HTMLElement) {
		dataPoint := strings.TrimSpace(e.Text)
		points := e.DOM.Find("div.frame-points.no-frame-info").Text()
		pts := floatFromPtsHeader(points)
		if strings.Contains(dataPoint, "Frontal Impact") {
			if adultFrontal {
				// fmt.Printf("Adult Frontal Impact-%v\n", strings.TrimSpace(points))
				ct.AdultOccupantRes.FrontalImpact = pts
				adultFrontal = false
			} else {
				ct.ChildOccupantRes.FrontalImpact = pts
				// fmt.Printf("Child Frontal Impact-%v\n", strings.TrimSpace(points))
			}
		}

		if strings.Contains(dataPoint, "Lateral Impact") {
			if adultLateral {
				ct.AdultOccupantRes.LateralImpact = pts
				// fmt.Printf("Adult Lateral Impact-%v\n", strings.TrimSpace(points))
				adultLateral = false
			} else {
				ct.ChildOccupantRes.LateralImpact = pts
				// fmt.Printf("Child Lateral Impact-%v\n", strings.TrimSpace(points))
			}
		}

		if strings.Contains(dataPoint, "Rear Impact") {
			ct.AdultOccupantRes.RearImpact = pts
			// fmt.Printf("Rear Impact-%v\n", strings.TrimSpace(points))
		}
	})

	// Set up a collector for the rows
	c.OnHTML("div.accordion-item-header", func(e *colly.HTMLElement) {
		// Extract text content of the current div element
		dataPoint := strings.TrimSpace(e.Text)

		// Clean up and format the output
		dataPoint = strings.Join(strings.Fields(dataPoint), " ")

		assignPoints(e, dataPoint, &ct)
	})

	//Rescue & Extrication
	c.OnHTML("div.accordion-item-header-frame", func(e *colly.HTMLElement) {
		dataPoint := strings.TrimSpace(e.Text)
		if strings.Contains(dataPoint, "Rescue and Extrication") {
			points := e.DOM.Find("span.points-in-frame").Text()
			ct.AdultOccupantRes.RescueExctrication = floatFromPtsHeader(strings.TrimSpace(points))
			// fmt.Printf("Rescue and Extrication-%v\n", strings.TrimSpace(points))
		}
	})

	//pedestrian protection table
	c.OnHTML("div.pedestrian-protection-table", func(e *colly.HTMLElement) {
		points1 := strings.TrimSpace(e.DOM.Find("p:nth-of-type(1) span:last-child").Text())
		points2 := strings.TrimSpace(e.DOM.Find("p:nth-of-type(2) span:last-child").Text())
		points3 := strings.TrimSpace(e.DOM.Find("p:nth-of-type(3) span:last-child").Text())

		ct.PedestrianRes.ImpactProtectionDetails.HeadImpact = floatFromPtsHeader(points1)
		ct.PedestrianRes.ImpactProtectionDetails.PelvisImpact = floatFromPtsHeader(points2)
		ct.PedestrianRes.ImpactProtectionDetails.LegImpact = floatFromPtsHeader(points3)
		// category1 := strings.TrimSpace(e.DOM.Find("p:nth-of-type(1) span:first-child").Text())
		// category2 := strings.TrimSpace(e.DOM.Find("p:nth-of-type(2) span:first-child").Text())
		// category3 := strings.TrimSpace(e.DOM.Find("p:nth-of-type(3) span:first-child").Text())
		// fmt.Printf("%v -%v\n", category1, points1)
		// fmt.Printf("%v -%v\n", category2, points2)
		// fmt.Printf("%v -%v\n", category3, points3)
	})
	err := c.Visit(url)
	return &ct, err
}

func valueFromImages(imageVal string) int8 {
	switch imageVal {
	case "/gfx/picto_standard.svg":
		return 5
	case "/gfx/picto_safety.svg":
		return 4
	case "/gfx/picto_not-safety.svg":
		return 3
	case "/gfx/picto_not-available.svg":
		return 2
	case "/gfx/picto_not-applicable.svg":
		return 1
	default:
		return -1
	}
}

func assignPoints(e *colly.HTMLElement, dp string, ct *data.CarTest) {
	if fp := filterPoints(e, dp, "Safety Features"); fp != nil {
		ct.ChildOccupantRes.SafetyFeatures = floatFromPts(dp)
		return
		// fmt.Printf("Safety Features-%v\n", *fp)
	}
	if fp := filterPoints(e, dp, "CRS Installation Check"); fp != nil {
		ct.ChildOccupantRes.CRSInstallationCheck = floatFromPts(dp)
		return
		// fmt.Printf("CRS Installation Check-%v\n", *fp)
	}
	if fp := filterPoints(e, dp, "AEB Pedestrian"); fp != nil {
		ct.PedestrianRes.AEBPedestrian = floatFromPts(dp)
		return
		// fmt.Printf("AEB Pedestrian-%v\n", *fp)
	}
	if fp := filterPoints(e, dp, "AEB Cyclist"); fp != nil {
		ct.PedestrianRes.AEBCyclist = floatFromPts(dp)
		return
		// fmt.Printf("AEB Cyclist-%v\n", *fp)
	}
	if fp := filterPoints(e, dp, "Speed Assistance"); fp != nil {
		ct.SafetyAssistRes.SpeedAssistance = floatFromPts(dp)
		return
		// fmt.Printf("Speed Assistance-%v\n", *fp)
	}
	if fp := filterPoints(e, dp, "Occupant Status Monitoring"); fp != nil {
		ct.SafetyAssistRes.OccupantStatusMonitoring = floatFromPts(dp)
		return
		// fmt.Printf("Occupant Status Monitoring-%v\n", *fp)
	}
	if fp := filterPoints(e, dp, "Lane Support"); fp != nil {
		ct.SafetyAssistRes.LaneSupport = floatFromPts(dp)
		return
		// fmt.Printf("Lane Support-%v\n", *fp)
	}
	if fp := filterPoints(e, dp, "AEB Car-to-Car"); fp != nil {
		ct.SafetyAssistRes.AEBCarToCar = floatFromPts(dp)
		return
		// fmt.Printf("AEB Car-to-Car-%v\n", *fp)
	}
	// if fp := filterPoints(e, dataPoint, "VRU Impact Mitigation"); fp != nil {
	// fmt.Printf("VRU Impact Mitigation-%v\n", *fp)
	// }
	// if fp := filterPoints(e, dataPoint, "Seatbelt Reminder"); fp != nil {
	// fmt.Printf("Seatbelt Reminder-%v\n", *fp)
	// }
	// if fp := filterPoints(e, dataPoint, "Driver Monitoring"); fp != nil {
	// fmt.Printf("Safety Features-%v\n", *fp)
	// }

}
func filterPoints(e *colly.HTMLElement, dp, want string) *string {
	if strings.Contains(dp, want) {
		// Find the corresponding span.points element within the current div
		points := strings.TrimSpace(e.DOM.Find("span.points").Text())
		// Print the extracted data
		return &points

	}
	return nil
}

func floatFromPtsHeader(s string) float32 {
	s = strings.TrimSpace(s)
	bef, _, f := strings.Cut(s, " ")
	if f {
		v, err := strconv.ParseFloat(bef, 32)
		if err == nil {
			return float32(v)
		}
	}
	return -1.0
}

func floatFromPts(s string) float32 {
	firstSplit := strings.Split(s, " / ")[0]
	splits := strings.Split(firstSplit, " ")
	l := len(splits)
	if l != 0 {
		v, err := strconv.ParseFloat(splits[l-1], 32)
		if err == nil {
			return float32(v)
		}
	}
	return -1.0
}

// // Child Safety Features Cancelled
// 	c.OnHTML("div.accordion-item-content > div.safety-equipment-table", func(e *colly.HTMLElement) {
// 		// Determine the number of columns by checking the presence of a specific class
// 		numColumns := 4
// 		if e.DOM.Find(".c5").Length() > 0 {
// 			numColumns = 5
// 		}

// 		// Iterate over each table row
// 		e.ForEach(".table-row", func(rowIdx int, row *colly.HTMLElement) {
// 			// Extract text content of the first column (c1)
// 			category := strings.TrimSpace(row.ChildText(".c1"))

// 			// Initialize variables for column values
// 			columnValues := make([]int, numColumns)

// 			// Function to determine numeric value for the image
// 			getImageValue := func(imgSrc string) int {
// 				switch imgSrc {
// 				case "/gfx/picto_standard.svg":
// 					return 2
// 				case "/gfx/picto_not-safety.svg":
// 					return 1
// 				case "/gfx/picto_not-available.svg":
// 					return 0
// 				default:
// 					return -1 // Default value or error handling
// 				}
// 			}

// 			// Extract and determine values for each column (c2, c3, c4, c5)
// 			for colIdx := 0; colIdx < numColumns; colIdx++ {
// 				colSelector := fmt.Sprintf(".c%d", colIdx+2)
// 				imgSrc := row.ChildAttr(colSelector+" img", "data-src")
// 				columnValues[colIdx] = getImageValue(imgSrc)
// 			}

// 			// Print the extracted data for each row
// 			fmt.Printf("%v: %v\n", category, columnValues)
// 			if strings.Contains(category, "sofix") {
// 				ct.ChildOccupantRes.SafetyFeatures.Isofix = data.ChildSFColumns{
// 					FrontPassenger:    int8(columnValues[0]),
// 					SecondRowOutBoard: int8(columnValues[1]),
// 					ThirdRowOutboard:  imn,
// 				}
// 			}

// 		})
// 	})
