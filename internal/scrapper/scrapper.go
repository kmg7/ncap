package scrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

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
	url := "https://www.euroncap.com/en" + asm.URL

	ct := asm.ParseToCarTest()

	switch ct.Version {
	case 1:
		return collectorv1(url, ct)
	case 2:
		return collectorv2(url, ct)
	case 3:
		return collectorv3(url, ct)
	default:
		return collectorv4(url, ct)
	}

}
