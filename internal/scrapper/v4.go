package scrapper

import (
	"github.com/gocolly/colly"
	"github.com/kmg7/ncap/internal/data"
)

func collectorv4(url string, ct *data.CarTest) (*data.CarTest, error) {
	c := colly.NewCollector()
	//Safety-equipment-table
	c.OnHTML("#tab2 > div:nth-child(1) > div.table-body", func(e *colly.HTMLElement) {
		safetyEquipmentv4(e, ct)
	})

	//Other Safety Systems
	c.OnHTML("#tab2 > div:nth-child(2) > div.table-body", func(h *colly.HTMLElement) {
		otherSafetyEquipmentv4(h, ct)
	})

	//Adult Tab
	c.OnHTML("#tab2-1", func(h *colly.HTMLElement) {
		adultTabv4(h, ct)
	})

	//Children Tab
	c.OnHTML("#tab2-2", func(h *colly.HTMLElement) {
		childTabv4(h, ct)
	})

	//VRU Tab
	c.OnHTML("#tab2-3", func(h *colly.HTMLElement) {
		vruTabv4(h, ct)
	})

	c.OnHTML("#tab2-4", func(h *colly.HTMLElement) {
		safetyAssistTabv4(h, ct)
	})
	err := c.Visit(url)
	return ct, err
}

func adultTabv4(e *colly.HTMLElement, ct *data.CarTest) {
	adultTabv3(e, ct)
}

func childTabv4(e *colly.HTMLElement, ct *data.CarTest) {
	childTabv1(e, ct)
}

func vruTabv4(e *colly.HTMLElement, ct *data.CarTest) {
	impTable := e.DOM.Find("div.pedestrian-protection-table")
	headImp := impTable.Find("p:nth-child(1) > span:nth-child(2)").Text()
	pelvisImp := impTable.Find("p:nth-child(2) > span:nth-child(2)").Text()
	femurImp := impTable.Find("p:nth-child(3) > span:nth-child(2)").Text()
	kneeTibiaImp := impTable.Find("p:nth-child(4) > span:nth-child(2)").Text()
	ct.VRUResult.ImpactProtectionDetails.Head = floatFromPtsHeader(headImp)
	ct.VRUResult.ImpactProtectionDetails.Pelvis = floatFromPtsHeader(pelvisImp)
	ct.VRUResult.ImpactProtectionDetails.Femur = floatFromPtsHeader(femurImp)
	ct.VRUResult.ImpactProtectionDetails.KneeAndTibia = floatFromPtsHeader(kneeTibiaImp)

	aebPedest := e.DOM.Find("div.accordion > div:nth-child(3) > div.accordion-item-header.sub-item-header > span").Text()
	aebCyclist := e.DOM.Find("div.accordion > div:nth-child(4) > div.accordion-item-header.sub-item-header > span").Text()
	cyclistDooring := e.DOM.Find("div.accordion > div:nth-child(5) > div.accordion-item-header.sub-item-header > span").Text()
	aebMotorcyclist := e.DOM.Find("div.accordion > div:nth-child(6) > div.accordion-item-header.sub-item-header > span").Text()
	laneSupMotorcy := e.DOM.Find("div.accordion > div:nth-child(7) > div.accordion-item-header.sub-item-header > span").Text()

	ct.VRUResult.AEBPedestrian = floatFromPtsHeader(aebPedest)
	ct.VRUResult.AEBCyclist = floatFromPtsHeader(aebCyclist)
	ct.VRUResult.CyclistDooring = floatFromPtsHeader(cyclistDooring)
	ct.VRUResult.AEBMotorcyclist = floatFromPtsHeader(aebMotorcyclist)
	ct.VRUResult.LaneSupMotorcyclist = floatFromPtsHeader(laneSupMotorcy)
}

func safetyAssistTabv4(e *colly.HTMLElement, ct *data.CarTest) {
	safetyAssistTabv3(e, ct)
}

func otherSafetyEquipmentv4(e *colly.HTMLElement, ct *data.CarTest) {
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
			ct.OtherSafetyEq.CyclistDooringPrev = numericValue
		case 4:
			ct.OtherSafetyEq.AEBMotorcyclist = numericValue
		case 5:
			ct.OtherSafetyEq.AEBCarToCar = numericValue
		case 6:
			ct.OtherSafetyEq.SpeedAssistance = numericValue
		case 7:
			ct.OtherSafetyEq.LaneAssistSystem = numericValue
		case 8:
			ct.OtherSafetyEq.FatigueDetection = numericValue
		}
	})
}

func safetyEquipmentv4(e *colly.HTMLElement, ct *data.CarTest) {
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
			ct.SafetyEq.ChildP.ChildPresenceDetection = columns
		case 12:
			ct.SafetyEq.SeatBeltReminder = columns
		}
	})
}
