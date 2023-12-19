package data

import "slices"

type CarTest struct {
	Version                  int                      `csv:"vers"`
	ClassID                  int                      `csv:"class_id"`
	Make                     string                   `csv:"make"`
	Model                    string                   `csv:"model"`
	Year                     int                      `csv:"year"`
	IsHybrid                 int                      `csv:"hybrid"`
	FullSafetyPack           int                      `csv:"safety"`
	Stars                    int                      `csv:"stars"`
	OverallRating            int                      `csv:"rating"`
	AdultOccupantRating      int                      `csv:"adult_rating"`
	ChildOccupantRating      int                      `csv:"child_rating"`
	VulnerableRoadUserRating int                      `csv:"vru_rating"`
	SafeAssistRating         int                      `csv:"assist_rating"`
	SafetyEq                 SafetyEquipment          `csv:"seq"`
	OtherSafetyEq            OtherSafetyEquipment     `csv:"o_seq"`
	AdultOccupantRes         AdultOccupantResult      `csv:"adult_result"`
	ChildOccupantRes         ChildOccupantResult      `csv:"child_result"`
	VRUResult                VulnerableRoadUserResult `csv:"vru_result"`
	SafetyAssistRes          SafetyAssistResult       `csv:"safety_result"`
}

type AdultOccupantResult struct {
	FrontalImpact      float32 `csv:"frontal_imp"`
	LateralImpact      float32 `csv:"lateral_imp"`
	RearImpact         float32 `csv:"rear_imp"`
	RescueExctrication float32 `csv:"res_ex"`
	AEBCity            float32 `csv:"aeb_city"`
}

type ChildOccupantResult struct {
	FrontalImpact        float32 `csv:"frontal_imp"`
	LateralImpact        float32 `csv:"lateral_imp"`
	SafetyFeatures       float32 `csv:"safety_feat"`
	CRSInstallationCheck float32 `csv:"crs_check"`
}

type VulnerableRoadUserResult struct {
	ImpactProtectionDetails VRUImpactProtection `csv:"impact_protect"`
	AEBPedestrian           float32             `csv:"aeb_pedest"`
	AEBCyclist              float32             `csv:"aeb_cyclist"`
	CyclistDooring          float32             `csv:"cyclist_dooring"`
	AEBMotorcyclist         float32             `csv:"aeb_motorcyclist"`
	LaneSupMotorcyclist     float32             `csv:"lane_sup_motocyc"`
}

type VRUImpactProtection struct {
	Head         float32 `csv:"head"`
	Pelvis       float32 `csv:"pelvis"`
	Leg          float32 `csv:"leg"`
	Femur        float32 `csv:"femur"`
	KneeAndTibia float32 `csv:"knee-tibia"`
}

type SafetyAssistResult struct {
	SpeedAssistance  float32 `csv:"speed_assist"`
	SeatbeltReminder float32 `csv:"seatbelt_rem"`
	DriverMonitoring float32 `csv:"driver_monitoring"`
	LaneSupport      float32 `csv:"lane_sup"`
	AEBInterUrban    float32 `csv:"aeb_inter_urban"`
	AEBCarToCar      float32 `csv:"aeb_car_to_car"`
}

type SafetyEquipment struct {
	FrontalCP        SEFrontalCrashProtection `csv:"frontal_p"`
	LateralCP        SELateralCrashProtection `csv:"lateral_p"`
	ChildP           SEChildProtection        `csv:"child_p"`
	SeatBeltReminder SEColumns                `csv:"belt_rem"`
}

type OtherSafetyEquipment struct {
	ActiveBonnet         int8 `csv:"active_bonnet"`
	SpeedAssistance      int8 `csv:"speed_assist"`
	LaneAssistSystem     int8 `csv:"lane_assist"`
	AEBPedestrian        int8 `csv:"aeb_pedest"`
	AEBCity              int8 `csv:"aeb_city"`
	AEBCyclist           int8 `csv:"aeb_cyclist"`
	AEBInterUrban        int8 `csv:"aeb_inter_urban"`
	AEBVRU               int8 `csv:"aeb_vru"`
	AEBPedestrianReverse int8 `csv:"aeb_pedest_reverse"`
	AEBCarToCar          int8 `csv:"aeb_car_to_car"`
	CyclistDooringPrev   int8 `csv:"cyclist_door_prev"`
	AEBMotorcyclist      int8 `csv:"aeb_motorcycle"`
	FatigueDetection     int8 `csv:"fatigue_detect"`
}

type SEFrontalCrashProtection struct {
	FrontAirbag      SEColumns `csv:"front_ab"`
	BeltPretensioner SEColumns `csv:"belt_pretension"`
	BeltLoadLimiter  SEColumns `csv:"belt_load_lim"`
	KneeAirbag       SEColumns `csv:"knee_ab"`
}

type SELateralCrashProtection struct {
	SideHeadAirbag   SEColumns `csv:"side_head_ab"`
	SideChestAirbag  SEColumns `csv:"side_chest_ab"`
	SidePelvisAirbag SEColumns `csv:"side_pelvis_ab"`
	CentreAirbag     SEColumns `csv:"centre_ab"`
}

type SEChildProtection struct {
	IsofixISize            SEColumns `csv:"isofix_isize"`
	IntegratedChildSeat    SEColumns `csv:"integr_child_seat"`
	AirbagCutOffSwitch     SEColumns `csv:"ab_cutoff_sw"`
	ChildPresenceDetection SEColumns `csv:"child_detect"`
}

type SEColumns struct {
	Driver    int8 `csv:"driver"`
	Passenger int8 `csv:"passenger"`
	Rear      int8 `csv:"rear"`
}

func (ct *CarTest) ToFixedPoints() {
	switch ct.Version {
	case 1:
		ct.v1ToFixesPoints()
	case 2:
		ct.v2ToFixesPoints()
	case 3:
		ct.v3ToFixesPoints()
	case 4:
		ct.v4ToFixesPoints()
	}

}

func (ct *CarTest) v1ToFixesPoints() {

	ct.AdultOccupantRes.FrontalImpact /= 16
	ct.AdultOccupantRes.LateralImpact /= 16
	ct.AdultOccupantRes.RearImpact /= 3
	ct.AdultOccupantRes.AEBCity /= 3

	ct.ChildOccupantRes.FrontalImpact /= 16
	ct.ChildOccupantRes.LateralImpact /= 8
	ct.ChildOccupantRes.SafetyFeatures /= 13
	ct.ChildOccupantRes.CRSInstallationCheck /= 12

	ct.VRUResult.ImpactProtectionDetails.Head /= 24
	ct.VRUResult.ImpactProtectionDetails.Pelvis /= 6
	ct.VRUResult.ImpactProtectionDetails.Leg /= 6
	ct.VRUResult.AEBPedestrian /= 6

	ct.SafetyAssistRes.SpeedAssistance /= 3
	ct.SafetyAssistRes.SeatbeltReminder /= 3
	ct.SafetyAssistRes.LaneSupport /= 3
	ct.SafetyAssistRes.AEBInterUrban /= 3
}

func (ct *CarTest) v2ToFixesPoints() {

	ct.AdultOccupantRes.FrontalImpact /= 16
	ct.AdultOccupantRes.LateralImpact /= 16
	ct.AdultOccupantRes.RearImpact /= 2
	ct.AdultOccupantRes.AEBCity /= 4

	ct.ChildOccupantRes.FrontalImpact /= 16
	ct.ChildOccupantRes.LateralImpact /= 8
	ct.ChildOccupantRes.SafetyFeatures /= 13
	ct.ChildOccupantRes.CRSInstallationCheck /= 12

	ct.VRUResult.ImpactProtectionDetails.Head /= 24
	ct.VRUResult.ImpactProtectionDetails.Pelvis /= 6
	ct.VRUResult.ImpactProtectionDetails.Leg /= 6
	ct.VRUResult.AEBPedestrian /= 6
	ct.VRUResult.AEBCyclist /= 6

	ct.SafetyAssistRes.SpeedAssistance /= 3
	ct.SafetyAssistRes.SeatbeltReminder /= 3
	ct.SafetyAssistRes.LaneSupport /= 4
	ct.SafetyAssistRes.AEBInterUrban /= 3
}

func (ct *CarTest) v3ToFixesPoints() {

	ct.AdultOccupantRes.FrontalImpact /= 16
	ct.AdultOccupantRes.LateralImpact /= 16
	ct.AdultOccupantRes.RearImpact /= 4
	ct.AdultOccupantRes.RescueExctrication /= 2

	ct.ChildOccupantRes.FrontalImpact /= 16
	ct.ChildOccupantRes.LateralImpact /= 8
	ct.ChildOccupantRes.SafetyFeatures /= 13
	ct.ChildOccupantRes.CRSInstallationCheck /= 12

	ct.VRUResult.ImpactProtectionDetails.Head /= 24
	ct.VRUResult.ImpactProtectionDetails.Pelvis /= 6
	ct.VRUResult.ImpactProtectionDetails.Leg /= 6
	ct.VRUResult.AEBPedestrian /= 9
	ct.VRUResult.AEBCyclist /= 9

	ct.SafetyAssistRes.SpeedAssistance /= 3
	ct.SafetyAssistRes.SeatbeltReminder /= 2
	ct.SafetyAssistRes.DriverMonitoring /= 1
	ct.SafetyAssistRes.LaneSupport /= 4
	ct.SafetyAssistRes.AEBCarToCar /= 6
}

func (ct *CarTest) v4ToFixesPoints() {

	ct.AdultOccupantRes.FrontalImpact /= 16
	ct.AdultOccupantRes.LateralImpact /= 16
	ct.AdultOccupantRes.RearImpact /= 4
	ct.AdultOccupantRes.RescueExctrication /= 4

	ct.ChildOccupantRes.FrontalImpact /= 16
	ct.ChildOccupantRes.LateralImpact /= 8
	ct.ChildOccupantRes.SafetyFeatures /= 13
	ct.ChildOccupantRes.CRSInstallationCheck /= 12

	ct.VRUResult.ImpactProtectionDetails.Head /= 18
	ct.VRUResult.ImpactProtectionDetails.Pelvis /= 4.5
	ct.VRUResult.ImpactProtectionDetails.Femur /= 4.5
	ct.VRUResult.ImpactProtectionDetails.KneeAndTibia /= 9

	ct.VRUResult.AEBPedestrian /= 9
	ct.VRUResult.AEBCyclist /= 8
	ct.VRUResult.CyclistDooring /= 1
	ct.VRUResult.AEBMotorcyclist /= 6
	ct.VRUResult.LaneSupMotorcyclist /= 3

	ct.SafetyAssistRes.SpeedAssistance /= 3
	ct.SafetyAssistRes.SeatbeltReminder /= 1
	ct.SafetyAssistRes.DriverMonitoring /= 2
	ct.SafetyAssistRes.LaneSupport /= 3
	ct.SafetyAssistRes.AEBCarToCar /= 9
}
func (las LatestAssessmentsSearchResult) Filter(pid []int) []*Assessment {
	filtered := []*Assessment{}
	for _, asr := range las.AssessmentSearchResults {
		if slices.Contains(pid, asr.ProtocolID) {
			filtered = append(filtered, asr.Assessments...)
		}
	}
	return filtered
}

func (asm Assessment) ParseToCarTest() *CarTest {
	return &CarTest{
		Version:                  int(protocolToVers(asm.ProtocolID)),
		ClassID:                  asm.ClassID,
		Make:                     asm.Make,
		Model:                    asm.Model,
		Year:                     asm.Year,
		IsHybrid:                 btoi(asm.IsHybrid),
		FullSafetyPack:           btoi(asm.FullSafetyPack),
		Stars:                    asm.Stars,
		OverallRating:            asm.OverallRating,
		ChildOccupantRating:      asm.ChildOccupantRating,
		AdultOccupantRating:      asm.AdultOccupantRating,
		VulnerableRoadUserRating: asm.PedestrianRating,
		SafeAssistRating:         asm.SafeAssistRating,
	}
}

func btoi(b bool) int {
	i := 0
	if b {
		i = 1
	}
	return i
}

// Determines protocol version by protocol id
// v1=2016-17, v2=2018-19, v3=2020-22, v4=2023
func protocolToVers(protocol int) uint8 {
	v1 := []int{
		24370, //2016
		26061, //2017
	}
	v2 := []int{
		30636, //2018
		34803, //2019
	}
	v3 := []int{
		40302, //2020
		41776, //2021
		45155, //2022
	}
	v4 := []int{
		49446, //2023
	}
	if slices.Contains(v1, protocol) {
		return 1
	}
	if slices.Contains(v2, protocol) {
		return 2
	}
	if slices.Contains(v3, protocol) {
		return 3
	}
	if slices.Contains(v4, protocol) {
		return 4
	}
	return 0
}

type LatestAssessmentsSearchResult struct {
	AssessmentSearchResults []struct {
		ProtocolID   int           `json:"ProtocolId"`
		ProtocolYear int           `json:"ProtocolYear"`
		Assessments  []*Assessment `json:"Assessments"`
	} `json:"AssessmentSearchResults"`
	PreAssessmentSearchResults []any `json:"PreAssessmentSearchResults"`
}

type Assessment struct {
	URL                 string `json:"Url"`
	AssessmentID        int    `json:"AssessmentId"`
	ID                  int    `json:"Id"`
	ClassID             int    `json:"ClassId"`
	Stars               int    `json:"Stars"`
	Make                string `json:"Make"`
	Model               string `json:"Model"`
	Year                int    `json:"Year"`
	OverallRating       int    `json:"OverallRating"`
	ChildOccupantRating int    `json:"ChildOccupantRating"`
	AdultOccupantRating int    `json:"AdultOccupantRating"`
	PedestrianRating    int    `json:"PedestrianRating"`
	SafeAssistRating    int    `json:"SafeAssistRating"`
	StandardSafetyPack  bool   `json:"StandardSafetyPack"`
	FullSafetyPack      bool   `json:"FullSafetyPack"`
	IsRatingExpired     bool   `json:"IsRatingExpired"`
	IsHybrid            bool   `json:"IsHybrid"`
	CreateDate          string `json:"CreateDate"`
	ReleaseDate         string `json:"ReleaseDate"`
	FirstPublished      string `json:"FirstPublished"`
	ProtocolID          int    `json:"ProtocolId"`
	// DriverAssistanceTechnologyIds    string `json:"DriverAssistanceTechnologyIds"`
	// AdvancedRewardTechnologyIds      string `json:"AdvancedRewardTechnologyIds"`
	// AdvancedRewardTechnologies       string `json:"AdvancedRewardTechnologies"`
	// AdvancedRewardTechnologiesCount  int    `json:"AdvancedRewardTechnologiesCount"`
	// FullSafetyRatingAssessmentID     int    `json:"FullSafetyRatingAssessmentId"`
	// StandardSafetyRatingAssessmentID int    `json:"StandardSafetyRatingAssessmentId"`
	// IsPreProtocol      bool   `json:"IsPreProtocol"`
	// IsFleet            bool   `json:"IsFleet"`
	// StarsWithOverallRating           int    `json:"StarsWithOverallRating"`// Another useless field Start+Score string addification
	// Content                          string `json:"Content"`// All Contents Test Already
	// IsPublished                      bool   `json:"IsPublished"`// Never gets value of false
	// DriverAssistanceTechnologies     string `json:"DriverAssistanceTechnologies"`
	// BestInClassCarClassID            int    `json:"BestInClassCarClassId"`
	// ModelID                          int    `json:"ModelId"`
	// MakeID                           int    `json:"MakeId"`
	// BestInClassCarClassEN            any    `json:"BestInClassCarClassEN"`
	// BestInClassCarClassDE            any    `json:"BestInClassCarClassDE"`
	// BestInClassCarClassNL            any    `json:"BestInClassCarClassNL"`
	// BestInClassCarClassES            any    `json:"BestInClassCarClassES"`
	// BestInClassCarClassFR            any    `json:"BestInClassCarClassFR"`
	// BestInClassCarClassIT            any    `json:"BestInClassCarClassIT"`
	// BestInClassCarClassRU            any    `json:"BestInClassCarClassRU"`
	// BestInClassCarClassSV            any    `json:"BestInClassCarClassSV"`
	// BestInClassCarClassTR            any    `json:"BestInClassCarClassTR"`
	// BestInClassCarClassZH            any    `json:"BestInClassCarClassZH"`
	// XML                              any    `json:"Xml"`
	// XMLUploaded                      bool   `json:"XmlUploaded"`
	// HideFromHomepageCarousel         bool   `json:"HideFromHomepageCarousel"`
	// Covid19Alert                     bool   `json:"Covid19Alert"`
	// FrontCarImage                    string `json:"FrontCarImage"`
	// CrashCarImage                    string `json:"CrashCarImage"`
	// Variants              string `json:"Variants"`
	// Title                 string `json:"Title"`
	// Name                  string `json:"Name"`
	// ClassEN                          string `json:"ClassEN"`
	// ClassDE                          string `json:"ClassDE"`
	// ClassNL                          string `json:"ClassNL"`
	// ClassES                          string `json:"ClassES"`
	// ClassFR                          string `json:"ClassFR"`
	// ClassIT                          string `json:"ClassIT"`
	// ClassRU                          string `json:"ClassRU"`
	// ClassSV                          string `json:"ClassSV"`
	// ClassTR                          string `json:"ClassTR"`
	// ClassZH                          string `json:"ClassZH"`
	// MakeImageLarge                   string `json:"MakeImageLarge"`
	// MakeImageSmall                   string `json:"MakeImageSmall"`
	// Images                           string `json:"Images"`
	// YoutubeVideos                    string `json:"YoutubeVideos"`
	// YoukuVideos                      string `json:"YoukuVideos"`
	// SortOrder       int    `json:"SortOrder"`
	// LifecycleStatus int    `json:"LifecycleStatus"`
}

// type ChildSafetyFeatures struct {
// 	Isofix        ChildSFColumns `csv:"isofix"`
// 	ISize         ChildSFColumns `csv:"isize"`
// 	IntegratedCRS ChildSFColumns `csv:"integ_crs"`
// }

// type ChildSFColumns struct {
// 	FrontPassenger    int8 `csv:"front_pass"`
// 	SecondRowOutBoard int8 `csv:"second_row_out_board"`
// 	SecondRowCenter   int8 `csv:"second_row_center"`
// 	ThirdRowOutboard  int8 `csv:"third_row_out_board"`
// }
