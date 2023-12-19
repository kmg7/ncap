package scrapper

import (
	"github.com/gocolly/colly"
	"github.com/kmg7/ncap/internal/data"
)

func collectorv3(url string, ct *data.CarTest) (*data.CarTest, error) {
	c := colly.NewCollector()
	//Safety-equipment-table
	c.OnHTML("#tab2 > div:nth-child(1) > div.table-body", func(e *colly.HTMLElement) {
		safetyEquipmentv3(e, ct)
	})

	//Other Safety Systems
	c.OnHTML("#tab2 > div:nth-child(2) > div.table-body", func(h *colly.HTMLElement) {
		otherSafetyEquipmentv3(h, ct)
	})

	//Adult Tab
	c.OnHTML("#tab2-1", func(h *colly.HTMLElement) {
		adultTabv3(h, ct)
	})

	//Children Tab
	c.OnHTML("#tab2-2", func(h *colly.HTMLElement) {
		childTabv3(h, ct)
	})

	//VRU Tab
	c.OnHTML("#tab2-3", func(h *colly.HTMLElement) {
		vruTabv3(h, ct)
	})

	c.OnHTML("#tab2-4", func(h *colly.HTMLElement) {
		safetyAssistTabv3(h, ct)
	})
	err := c.Visit(url)
	return ct, err
}
func adultTabv3(e *colly.HTMLElement, ct *data.CarTest) {
	frontalImp := e.DOM.Find("div.frame.w-1-1.frontal-block > div.frame-header > div.frame-points.no-frame-info").Text()
	lateralImp := e.DOM.Find("div:nth-child(5) > div.frame-header > div.frame-points.no-frame-info").Text()
	rearImp := e.DOM.Find("div.frame.w-1-1.whiplash-block > div.frame-header > div.frame-points.no-frame-info").Text()
	rescueExt := e.DOM.Find("div:nth-child(7) > div > div > div.accordion-item-header.accordion-item-header-frame > span").Text()
	ct.AdultOccupantRes.FrontalImpact = floatFromPtsHeader(frontalImp)
	ct.AdultOccupantRes.LateralImpact = floatFromPtsHeader(lateralImp)
	ct.AdultOccupantRes.RearImpact = floatFromPtsHeader(rearImp)
	ct.AdultOccupantRes.RescueExctrication = floatFromPtsHeader(rescueExt)
}

func childTabv3(e *colly.HTMLElement, ct *data.CarTest) {
	childTabv1(e, ct)
}

func vruTabv3(e *colly.HTMLElement, ct *data.CarTest) {
	vruTabv2(e, ct)
}

func safetyAssistTabv3(e *colly.HTMLElement, ct *data.CarTest) {
	speedAssist := e.DOM.Find("div.accordion > div:nth-child(1) > div.accordion-item-header > span").Text()
	seatbeltRem := e.DOM.Find("div.accordion > div:nth-child(2) > div.accordion-item-content > div:nth-child(1) > div.accordion-item-header.no-background > span.points").Text()
	driverMon := e.DOM.Find("div.accordion > div:nth-child(2) > div.accordion-item-content > div:nth-child(1) > div.accordion-item-header.no-background > span.points").Text()
	laneSup := e.DOM.Find("div.accordion > div:nth-child(3) > div.accordion-item-header > span").Text()
	aebCarToCar := e.DOM.Find("div.accordion > div:nth-child(4) > div.accordion-item-header > span").Text()
	ct.SafetyAssistRes.SpeedAssistance = floatFromPtsHeader(speedAssist)
	ct.SafetyAssistRes.SeatbeltReminder = floatFromPtsHeader(seatbeltRem)
	ct.SafetyAssistRes.DriverMonitoring = floatFromPtsHeader(driverMon)
	ct.SafetyAssistRes.LaneSupport = floatFromPtsHeader(laneSup)
	ct.SafetyAssistRes.AEBCarToCar = floatFromPtsHeader(aebCarToCar)
}

func otherSafetyEquipmentv3(e *colly.HTMLElement, ct *data.CarTest) {
	e.ForEach(".table-row:not(.title)", func(i int, row *colly.HTMLElement) {
		imageValue := row.ChildAttr(".table-row-c.c2 img", "data-src")
		numericValue := valueFromImages(imageValue)

		switch i {
		case 0:
			ct.OtherSafetyEq.ActiveBonnet = numericValue
		case 1:
			ct.OtherSafetyEq.AEBVRU = numericValue
		case 2:
			ct.OtherSafetyEq.AEBPedestrianReverse = numericValue
		case 3:
			ct.OtherSafetyEq.AEBCarToCar = numericValue
		case 4:
			ct.OtherSafetyEq.SpeedAssistance = numericValue
		case 5:
			ct.OtherSafetyEq.LaneAssistSystem = numericValue
		}
	})
}

func safetyEquipmentv3(e *colly.HTMLElement, ct *data.CarTest) {
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
	})
}
