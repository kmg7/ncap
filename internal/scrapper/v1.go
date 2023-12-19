package scrapper

import (
	"github.com/gocolly/colly"
	"github.com/kmg7/ncap/internal/data"
)

func collectorv1(url string, ct *data.CarTest) (*data.CarTest, error) {
	c := colly.NewCollector()
	//Safety-equipment-table
	c.OnHTML("#tab2 > div:nth-child(1) > div.table-body", func(e *colly.HTMLElement) {
		safetyEquipmentv1(e, ct)
	})

	//Other Safety Systems
	c.OnHTML("#tab2 > div:nth-child(2) > div.table-body", func(h *colly.HTMLElement) {
		otherSafetyEquipmentv1(h, ct)
	})

	//Adult Tab
	c.OnHTML("#tab2-1", func(h *colly.HTMLElement) {
		adultTabv1(h, ct)
	})

	//Children Tab
	c.OnHTML("#tab2-2", func(h *colly.HTMLElement) {
		childTabv1(h, ct)
	})

	//VRU Tab
	c.OnHTML("#tab2-3", func(h *colly.HTMLElement) {
		vruTabv1(h, ct)
	})

	c.OnHTML("#tab2-4", func(h *colly.HTMLElement) {
		safetyAssistTabv1(h, ct)
	})
	err := c.Visit(url)
	return ct, err
}

func adultTabv1(e *colly.HTMLElement, ct *data.CarTest) {
	frontalImp1 := e.DOM.Find("div.frame.w-1-2.frontal-block > div.frame-header > div.frame-points.no-frame-info").Text()
	frontaImp2 := e.DOM.Find("div.frame.w-1-2.frontal-full-width > div.frame-header > div.frame-points.no-frame-info").Text()
	lateralImp := e.DOM.Find("div:nth-child(5) > div.frame-header > div.frame-points.no-frame-info").Text()
	rearImp := e.DOM.Find("div.frame.w-1-2.whiplash-block > div.frame-header > div.frame-points.no-frame-info").Text()
	aebCity := e.DOM.Find("div.frame.w-1-1.aeb-city > div.frame-header > div.frame-points.no-frame-info").Text()
	ct.AdultOccupantRes.FrontalImpact = floatFromPtsHeader(frontalImp1) + floatFromPtsHeader(frontaImp2)
	ct.AdultOccupantRes.LateralImpact = floatFromPtsHeader(lateralImp)
	ct.AdultOccupantRes.RearImpact = floatFromPtsHeader(rearImp)
	ct.AdultOccupantRes.AEBCity = floatFromPtsHeader(aebCity)
}

func childTabv1(e *colly.HTMLElement, ct *data.CarTest) {
	frontalImp := e.DOM.Find("div.accordion-item-content > div:nth-child(1) > div.frame-header > div.frame-points.no-frame-info").Text()
	lateralImp := e.DOM.Find("div.accordion-item-content > div:nth-child(2) > div.frame-header > div.frame-points.no-frame-info").Text()
	safetyFeat := e.DOM.Find("div.accordion > div:nth-child(2) > div.accordion-item-header > span").Text()
	crsCheck := e.DOM.Find("div.accordion > div:nth-child(3) > div.accordion-item-header > span").Text()
	ct.ChildOccupantRes.FrontalImpact = floatFromPtsHeader(frontalImp)
	ct.ChildOccupantRes.LateralImpact = floatFromPtsHeader(lateralImp)
	ct.ChildOccupantRes.SafetyFeatures = floatFromPtsHeader(safetyFeat)
	ct.ChildOccupantRes.CRSInstallationCheck = floatFromPtsHeader(crsCheck)
}

func vruTabv1(e *colly.HTMLElement, ct *data.CarTest) {
	impTable := e.DOM.Find("div.pedestrian-protection-table")
	headImp := impTable.Find("p:nth-child(1) > span:nth-child(2)").Text()
	pelvisImp := impTable.Find("p:nth-child(2) > span:nth-child(2)").Text()
	legImp := impTable.Find("p:nth-child(3) > span:nth-child(2)").Text()
	ct.VRUResult.ImpactProtectionDetails.Head = floatFromPtsHeader(headImp)
	ct.VRUResult.ImpactProtectionDetails.Pelvis = floatFromPtsHeader(pelvisImp)
	ct.VRUResult.ImpactProtectionDetails.Leg = floatFromPtsHeader(legImp)

	aebPedest := e.DOM.Find("div.accordion > div:nth-child(2) > div.accordion-item-header > span").Text()

	ct.VRUResult.AEBPedestrian = floatFromPtsHeader(aebPedest)
}

func safetyAssistTabv1(e *colly.HTMLElement, ct *data.CarTest) {
	speedAssist := e.DOM.Find("div.accordion > div:nth-child(1) > div.accordion-item-header > span").Text()
	seatbeltRem := e.DOM.Find("div.accordion > div:nth-child(2) > div.accordion-item-header > span").Text()
	laneSup := e.DOM.Find("div.accordion > div:nth-child(3) > div.accordion-item-header > span").Text()
	aebInterUrban := e.DOM.Find("div.accordion > div:nth-child(4) > div.accordion-item-header > span").Text()
	ct.SafetyAssistRes.SpeedAssistance = floatFromPtsHeader(speedAssist)
	ct.SafetyAssistRes.SeatbeltReminder = floatFromPtsHeader(seatbeltRem)
	ct.SafetyAssistRes.LaneSupport = floatFromPtsHeader(laneSup)
	ct.SafetyAssistRes.AEBInterUrban = floatFromPtsHeader(aebInterUrban)
}

func otherSafetyEquipmentv1(e *colly.HTMLElement, ct *data.CarTest) {
	e.ForEach(".table-row:not(.title)", func(i int, row *colly.HTMLElement) {
		imageValue := row.ChildAttr(".table-row-c.c2 img", "data-src")
		numericValue := valueFromImages(imageValue)

		switch i {
		case 0:
			ct.OtherSafetyEq.ActiveBonnet = numericValue
		case 1:
			ct.OtherSafetyEq.AEBPedestrian = numericValue
		case 2:
			ct.OtherSafetyEq.AEBCity = numericValue
		case 3:
			ct.OtherSafetyEq.AEBInterUrban = numericValue
		case 4:
			ct.OtherSafetyEq.SpeedAssistance = numericValue
		case 5:
			ct.OtherSafetyEq.LaneAssistSystem = numericValue
		}
	})
}

func safetyEquipmentv1(e *colly.HTMLElement, ct *data.CarTest) {
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
			ct.SafetyEq.ChildP.IsofixISize = columns
		case 8:
			ct.SafetyEq.ChildP.IntegratedChildSeat = columns
		case 9:
			ct.SafetyEq.ChildP.AirbagCutOffSwitch = columns
		case 10:
			ct.SafetyEq.SeatBeltReminder = columns
		}
	})
}
