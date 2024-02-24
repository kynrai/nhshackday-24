package model

import "time"

type Data struct {
	Meta struct {
		Href string `json:"href"`
	} `json:"meta"`
	CompositionUID string `json:"compositionUid"`
	Format         string `json:"format"`
	TemplateID     string `json:"templateId"`
	Composition    struct {
		MtxReport struct {
			UID      []string `json:"_uid"`
			Language []struct {
				Code        string `json:"|code"`
				Terminology string `json:"|terminology"`
			} `json:"language"`
			Territory []struct {
				Code        string `json:"|code"`
				Terminology string `json:"|terminology"`
			} `json:"territory"`
			Context []struct {
				StartTime []time.Time `json:"start_time"`
				Setting   []struct {
					Code        string `json:"|code"`
					Value       string `json:"|value"`
					Terminology string `json:"|terminology"`
				} `json:"setting"`
			} `json:"context"`
			LabResults []struct {
				UID            []string `json:"_uid"`
				FullBloodCount []struct {
					UID      []string `json:"_uid"`
					TestName []string `json:"test_name"`
					Hb       []struct {
						AnalyteName   []string `json:"analyte_name"`
						AnalyteResult []struct {
							Magnitude float64 `json:"|magnitude"`
							Unit      string  `json:"|unit"`
						} `json:"analyte_result"`
						ReferenceRangeGuidance []string `json:"reference_range_guidance"`
					} `json:"hb"`
					Wbc []struct {
						AnalyteName   []string `json:"analyte_name"`
						AnalyteResult []struct {
							Magnitude float64 `json:"|magnitude"`
							Unit      string  `json:"|unit"`
						} `json:"analyte_result"`
						ReferenceRangeGuidance []string `json:"reference_range_guidance"`
					} `json:"wbc"`
					Time     []time.Time `json:"time"`
					Language []struct {
						Code        string `json:"|code"`
						Terminology string `json:"|terminology"`
					} `json:"language"`
					Encoding []struct {
						Code        string `json:"|code"`
						Terminology string `json:"|terminology"`
					} `json:"encoding"`
				} `json:"full_blood_count"`
			} `json:"lab_results"`
			Prescription []struct {
				UID                    []string `json:"_uid"`
				MedicationUseStatement []struct {
					UID                          []string    `json:"_uid"`
					MedicationName               []string    `json:"medication_name"`
					OverallDirectionsDescription []string    `json:"overall_directions_description"`
					RouteOfAdministration        []string    `json:"route_of_administration"`
					ClinicalIndication           []string    `json:"clinical_indication"`
					Time                         []time.Time `json:"time"`
					Language                     []struct {
						Code        string `json:"|code"`
						Terminology string `json:"|terminology"`
					} `json:"language"`
					Encoding []struct {
						Code        string `json:"|code"`
						Terminology string `json:"|terminology"`
					} `json:"encoding"`
				} `json:"medication_use_statement"`
			} `json:"prescription"`
			BodyWeight []struct {
				UID    []string `json:"_uid"`
				Weight []struct {
					Magnitude float64 `json:"|magnitude"`
					Unit      string  `json:"|unit"`
				} `json:"weight"`
				Time     []time.Time `json:"time"`
				Language []struct {
					Code        string `json:"|code"`
					Terminology string `json:"|terminology"`
				} `json:"language"`
				Encoding []struct {
					Code        string `json:"|code"`
					Terminology string `json:"|terminology"`
				} `json:"encoding"`
			} `json:"body_weight"`
			HeightLength []struct {
				UID                []string `json:"_uid"`
				OtherParticipation []struct {
					Function string `json:"|function"`
					Mode     string `json:"|mode"`
					Name     string `json:"|name"`
				} `json:"_other_participation"`
				HeightLength []struct {
					Magnitude float64 `json:"|magnitude"`
					Unit      string  `json:"|unit"`
				} `json:"height_length"`
				Time     []time.Time `json:"time"`
				Language []struct {
					Code        string `json:"|code"`
					Terminology string `json:"|terminology"`
				} `json:"language"`
				Encoding []struct {
					Code        string `json:"|code"`
					Terminology string `json:"|terminology"`
				} `json:"encoding"`
			} `json:"height_length"`
			Category []struct {
				Code        string `json:"|code"`
				Value       string `json:"|value"`
				Terminology string `json:"|terminology"`
			} `json:"category"`
			Composer []struct {
				Name string `json:"|name"`
			} `json:"composer"`
		} `json:"mtx_report"`
	} `json:"composition"`
	Deleted            bool      `json:"deleted"`
	LastVersion        bool      `json:"lastVersion"`
	EhrID              string    `json:"ehrId"`
	LifecycleState     string    `json:"lifecycleState"`
	CommitterName      string    `json:"committerName"`
	ChangeType         string    `json:"changeType"`
	CommittedTimestamp time.Time `json:"committedTimestamp"`
}
