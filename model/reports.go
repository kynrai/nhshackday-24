package model

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed reports
var reports embed.FS

func MtxReports() []SingleMtxReport {
	rs := []SingleMtxReport{}
	for i := 0; i < 9; i++ {
		b, err := reports.ReadFile(fmt.Sprintf("reports/week_%d_report.json", i))
		if err != nil {
			log.Fatal(err)
		}
		var r SingleMtxReport
		err = json.Unmarshal(b, &r)
		if err != nil {
			log.Fatal(err)
		}
		rs = append(rs, r)
	}
	return rs
}
