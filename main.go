package main

import (
	"fmt"
	"time"

	"github.com/kmg7/ncap/internal/data"
	"github.com/kmg7/ncap/internal/scrapper"
)

func main() {
	println("Started")
	started := time.Now()
	res := getResults()
	asms := res.Filter([]int{
		24370, //2016
		26061, //2017
		30636, //2018
		34803, //2019
		40302, //2020
		41776, //2021
		45155, //2022
		49446, //2023
	})

	cts := scrapeAll(asms)

	data.SaveCarTests(cts)
	fmt.Printf("Finished in %v", time.Since(started))
}

func scrapeAll(asms []*data.Assessment) []*data.CarTest {
	fmt.Printf("For kindness before visiting every url, program will wait 5 seconds\n")
	fmt.Printf("Starting scraping test result for %v assessments.\n", len(asms))
	cts := []*data.CarTest{}
	for i, asm := range asms {
		time.Sleep(time.Second * 5)
		started := time.Now()
		ct, err := scrapper.CarResultFromAssessment(*asm)
		if err != nil {
			fmt.Printf("[%v]-[FAILED] %v [%vms]\n", i, err.Error(), time.Since(started))
		} else {
			fmt.Printf("[%v]-[SUCCESS] [%vms]\n", i, time.Since(started).Milliseconds())
			ct.ToFixedPoints()
			cts = append(cts, ct)
		}
	}
	return cts
}
func getResults() data.LatestAssessmentsSearchResult {
	assessments := data.LatestAssessmentsSearchResult{}
	if data.TodaysFileExists() {
		assessments = data.ReadAssessments()
	} else {
		assessments = scrapper.FetchAssessments()
		totalTests := 0
		for _, v := range assessments.AssessmentSearchResults {
			totalTests += len(v.Assessments)
			fmt.Printf("Protocol Id: %v, Year: %v, Total Numbe Of Tests: %v\n", v.ProtocolID, v.ProtocolYear, len(v.Assessments))
		}
		fmt.Printf("Fetching complete number of total tests is: %v\n", totalTests)
		data.SaveAssessments(assessments)
	}
	return assessments
}
