package scrapper

import (
	"github.com/gocolly/colly"
	"github.com/kmg7/ncap/internal/data"
)

func collectorv2(url string, ct *data.CarTest) (*data.CarTest, error) {
	c := colly.NewCollector()
	//Safety-equipment-table
	c.OnHTML("#tab2 > div:nth-child(1) > div.table-body", func(e *colly.HTMLElement) {
		safetyEquipmentv2(e, ct)
	})

	//Other Safety Systems
	c.OnHTML("#tab2 > div:nth-child(2) > div.table-body", func(h *colly.HTMLElement) {
		otherSafetyEquipmentv2(h, ct)
	})

	//Adult Tab
	c.OnHTML("#tab2-1", func(h *colly.HTMLElement) {
		adultTabv2(h, ct)
	})

	//Children Tab
	c.OnHTML("#tab2-2", func(h *colly.HTMLElement) {
		childTabv2(h, ct)
	})

	//VRU Tab
	c.OnHTML("#tab2-3", func(h *colly.HTMLElement) {
		vruTabv2(h, ct)
	})

	c.OnHTML("#tab2-4", func(h *colly.HTMLElement) {
		safetyAssistTabv2(h, ct)
	})
	err := c.Visit(url)
	return ct, err
}

func adultTabv2(e *colly.HTMLElement, ct *data.CarTest) {
	adultTabv1(e, ct)
}

func childTabv2(e *colly.HTMLElement, ct *data.CarTest) {
	childTabv1(e, ct)
}

func vruTabv2(e *colly.HTMLElement, ct *data.CarTest) {
	impTable := e.DOM.Find("div.pedestrian-protection-table")
	headImp := impTable.Find("p:nth-child(1) > span:nth-child(2)").Text()
	pelvisImp := impTable.Find("p:nth-child(2) > span:nth-child(2)").Text()
	legImp := impTable.Find("p:nth-child(3) > span:nth-child(2)").Text()
	ct.VRUResult.ImpactProtectionDetails.Head = floatFromPtsHeader(headImp)
	ct.VRUResult.ImpactProtectionDetails.Pelvis = floatFromPtsHeader(pelvisImp)
	ct.VRUResult.ImpactProtectionDetails.Leg = floatFromPtsHeader(legImp)

	aebPedest := e.DOM.Find("div.accordion > div:nth-child(3) > div.accordion-item-header.sub-item-header > span").Text()
	aebCyclist := e.DOM.Find("div.accordion > div:nth-child(4) > div.accordion-item-header.sub-item-header > span").Text()

	ct.VRUResult.AEBPedestrian = floatFromPtsHeader(aebPedest)
	ct.VRUResult.AEBCyclist = floatFromPtsHeader(aebCyclist)
}

func safetyAssistTabv2(e *colly.HTMLElement, ct *data.CarTest) {
	safetyAssistTabv1(e, ct)
}

func safetyEquipmentv2(e *colly.HTMLElement, ct *data.CarTest) {
	safetyEquipmentv1(e, ct)
}

func otherSafetyEquipmentv2(e *colly.HTMLElement, ct *data.CarTest) {
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
			ct.OtherSafetyEq.AEBCyclist = numericValue
		case 4:
			ct.OtherSafetyEq.AEBInterUrban = numericValue
		case 5:
			ct.OtherSafetyEq.SpeedAssistance = numericValue
		case 6:
			ct.OtherSafetyEq.LaneAssistSystem = numericValue
		}
	})
}
