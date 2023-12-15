package data

import "slices"

type CarTest struct {
	AssessmentID        int                      `csv:"asm_id"`
	ProtocolID          int                      `csv:"prot_id"`
	ID                  int                      `csv:"car_id"`
	ClassID             int                      `csv:"class_id"`
	Make                string                   `csv:"make"`
	Model               string                   `csv:"model"`
	Year                int                      `csv:"year"`
	IsHybrid            int                      `csv:"hybrid"`
	StandardSafetyPack  int                      `csv:"standard"`
	FullSafetyPack      int                      `csv:"safety"`
	IsRatingExpired     int                      `csv:"expired"`
	Stars               int                      `csv:"start"`
	OverallRating       int                      `csv:"overall_r"`
	AdultOccupantRating int                      `csv:"adult_r"`
	ChildOccupantRating int                      `csv:"child_r"`
	PedestrianRating    int                      `csv:"pedest_r"`
	SafeAssistRating    int                      `csv:"assist_r"`
	SafetyEq            SafetyEquipment          `csv:"safety_eq"`
	AdultOccupantRes    AdultOccupantResult      `csv:"adult_res"`
	ChildOccupantRes    ChildOccupantResult      `csv:"child_res"`
	PedestrianRes       VulnerableRoadUserResult `csv:"vru_res"`
	SafetyAssistRes     SafetyAssistResult       `csv:"safety_res"`
}
type AdultOccupantResult struct {
	FrontalImpact      float32 `csv:"frontal_imp"`
	LateralImpact      float32 `csv:"lateral_imp"`
	RearImpact         float32 `csv:"rear_imp"`
	RescueExctrication float32 `csv:"res_ex"`
}

type ChildOccupantResult struct {
	FrontalImpact        float32 `csv:"frontal_imp"`
	LateralImpact        float32 `csv:"lateral_imp"`
	SafetyFeatures       float32 `csv:"safety_feat"`
	CRSInstallationCheck float32 `csv:"crs_inst_check"`
	// SafetyFeaturesFeatures       ChildSafetyFeatures `csv:"safety_feat"`
}

type VulnerableRoadUserResult struct {
	ImpactProtectionDetails VRUImpactProtection `csv:"impact_protect"`
	AEBPedestrian           float32             `csv:"aeb_pedest"`
	AEBCyclist              float32             `csv:"aeb_cyclist"`
	// ImpactProtection        float32             `csv:"impact_protect"`
}

type VRUImpactProtection struct {
	HeadImpact   float32 `csv:"head_impact"`
	PelvisImpact float32 `csv:"pelvis_impact"`
	LegImpact    float32 `csv:"leg_impact"`
}

type SafetyAssistResult struct {
	SpeedAssistance          float32 `csv:"speed_assist"`
	OccupantStatusMonitoring float32 `csv:"occupant_status_mon"`
	LaneSupport              float32 `csv:"lane_sup"`
	AEBCarToCar              float32 `csv:"aeb_car_to_car"`
}

type SafetyEquipment struct {
	FrontalCP              SEFrontalCrashProtection `csv:"frontal_crash_p"`
	LateralCP              SELateralCrashProtection `csv:"lateral_crash_p"`
	ChildP                 SEChildProtection        `csv:"child_p"`
	SeatBeltReminder       SEColumns                `csv:"belt_rem"`
	ActiveBonnet           int8                     `csv:"active_bonnet"`
	AEBVulnerableRoadUsers int8                     `csv:"aeb_vru"`
	AEBPedestrianReverse   int8                     `csv:"aeb_pedest"`
	AEBCarToCar            int8                     `csv:"aeb_car_to_car"`
	SpeedAssistance        int8                     `csv:"speed_assist"`
	LaneAssistSystem       int8                     `csv:"lane_assist"`
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
	IsofixISize         SEColumns `csv:"isofix_isize"`
	IntegratedChildSeat SEColumns `csv:"integr_child_seat"`
	AirbagCutOffSwitch  SEColumns `csv:"ab_cutoff_sw"`
}

type SEColumns struct {
	Driver    int8 `csv:"driver"`
	Passenger int8 `csv:"passenger"`
	Rear      int8 `csv:"rear"`
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

func (asm Assessment) ParseToCarTest() CarTest {
	return CarTest{
		ProtocolID:          asm.ProtocolID,
		AssessmentID:        asm.AssessmentID,
		ID:                  asm.ID,
		ClassID:             asm.ClassID,
		Make:                asm.Make,
		Model:               asm.Model,
		Year:                asm.Year,
		IsHybrid:            btoi(asm.IsHybrid),
		StandardSafetyPack:  btoi(asm.StandardSafetyPack),
		FullSafetyPack:      btoi(asm.FullSafetyPack),
		IsRatingExpired:     btoi(asm.IsRatingExpired),
		Stars:               asm.Stars,
		OverallRating:       asm.OverallRating,
		ChildOccupantRating: asm.ChildOccupantRating,
		AdultOccupantRating: asm.AdultOccupantRating,
		PedestrianRating:    asm.PedestrianRating,
		SafeAssistRating:    asm.SafeAssistRating,
	}
}

func btoi(b bool) int {
	i := 0
	if b {
		i = 1
	}
	return i
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
