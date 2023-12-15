package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/gocarina/gocsv"
)

func ReadAssessments() LatestAssessmentsSearchResult {
	var res LatestAssessmentsSearchResult
	data, err := os.ReadFile(todaysFile())
	panicOnErr(err)
	err = json.Unmarshal(data, &res)
	panicOnErr(err)
	return res
}

func SaveAssessments(results LatestAssessmentsSearchResult) {
	data, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(todaysFile())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Assessments saved")

}

func SaveCarTests(resulst []*CarTest) (*string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	out := path.Join(wd, "out", "tests")
	err = os.MkdirAll(out, 0755)
	if err != nil {
		return nil, err
	}
	f := path.Join(out, fmt.Sprintf("%v.csv", time.Now().Format(time.DateOnly)))
	testF, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer testF.Close()
	err = gocsv.MarshalFile(&resulst, testF)
	return &f, err
}

func todaysFile() string {
	wd, err := os.Getwd()
	panicOnErr(err)

	out := path.Join(wd, "out", "assessments")
	err = os.MkdirAll(out, 0755)
	panicOnErr(err)

	f := fmt.Sprintf("%v.json", time.Now().Format(time.DateOnly))
	return path.Join(out, f)
}

func TodaysFileExists() bool {
	if _, err := os.Stat(todaysFile()); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		panicOnErr(err)
	}
	return true
}

func panicOnErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
