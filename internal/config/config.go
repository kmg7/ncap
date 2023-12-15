package config

const Url = "https://www.euroncap.com/Umbraco/EuroNCAP/SearchApi/GetAssessmentSearch?protocols=49446,45155,41776,40302,34803,30636,26061,24370&make=0&model=0&carClasses=1202,1199,1201,1196,1205,1203,1198,1179,40250,1197,1204,1180,34736,44997&driverAssistanceTechnologies&allProtocols=true&allClasses=true&allDriverAssistanceTechnologies=false&includeFullSafetyPackage=true&includeStandardSafetyPackage=true&showOnlyHybrid=false&showOnlyFleet=false&starNumber&thirdRowFitment=false"

type FetchOptions struct {
	Protocols                       []string
	Make                            string
	Model                           string
	CarClasses                      []string
	AllProtocols                    bool
	AllClasses                      bool
	AllDriverAssistanceTechnologies bool
	IncludeFullSafetyPackage        bool
	IncludeStandardSafetyPackage    bool
	ShowOnlyHybrid                  bool
	ShowOnlyFleet                   bool
	StarNumber                      string
	ThirdRowFitment                 bool
}
