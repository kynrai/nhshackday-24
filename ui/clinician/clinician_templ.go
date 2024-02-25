// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package clinician

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/kynrai/nhshackday-24/model"
)

func ClinicianView(tabType string, data model.Data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"clinician\" class=\"flex h-screen\"><div class=\"w-[25%] h-screen border-r border-gray-400 pr-10 flex flex-col gap-4 items-stretch\"><button class=\"px-4 py-2 rounded-md bg-nhs-green-2 text-white border-b-2 border-b-white\" type=\"button\">Back to Patients</button><div class=\"flex flex-col gap-2 border-[1px] border-b-[4px] border-nhs-grey-2 p-6\"><div class=\"text-lg font-semibold\">Patient Profile</div><div class=\"flex justify-between\"><div class=\"font-medium\">Name: </div><div>Alan</div></div><div class=\"flex justify-between\"><div class=\"font-medium\">Age: </div><div>33</div></div><div class=\"flex justify-between\"><div class=\"font-medium\">DOB: </div><div>1/1/1999</div></div><div class=\"flex justify-between\"><div class=\"font-medium\">Gender: </div><div>M</div></div><div class=\"flex justify-between\"><div class=\"font-medium\">Height: </div><div>175cm</div></div><div class=\"flex justify-between\"><div class=\"font-medium\">Weight: </div><div>65kg</div></div><div class=\"flex justify-between\"><div class=\"font-medium\">Ethnicity: </div><div>White</div></div><div class=\"flex justify-between\"><div class=\"font-medium\">Language: </div><div>English</div></div></div></div><div id=\"tab\" class=\"flex flex-col w-[85%]\"><nav class=\"h-10 w-full border-b border-gray-400\"><ul class=\"ml-10 h-10 flex text-gray-500 cursor-pointer\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{"px-4 py-2 border border-gray-400 bg-gray-200 rounded-tl", templ.KV("border-b-white bg-white", tabType == "prescription"), templ.KV("hover:bg-blue-100", tabType != "prescription")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/hx/nav/" + "prescription"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#clinician\" hx-swap=\"morph:outerHTML\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Prescriptions</li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 = []any{"px-4 py-2 border-y border-gray-400 bg-gray-200", templ.KV("border-b-white bg-white", tabType == "lab"), templ.KV("hover:bg-blue-100", tabType != "lab")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/hx/nav/" + "lab"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#clinician\" hx-swap=\"morph:outerHTML\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var3).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Lab Results</li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 = []any{"px-4 py-2 border border-gray-400 bg-gray-200 rounded-tr", templ.KV("border-b-white bg-white", tabType == "appointment"), templ.KV("hover:bg-blue-100", tabType != "appointment")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/hx/nav/" + "appointment"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#clinician\" hx-swap=\"morph:outerHTML\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var4).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Appointments</li></ul></nav><div class=\"mx-10 my-6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if tabType == "prescription" {
			templ_7745c5c3_Err = PrescriptionTable(data.Composition.MtxReport.Prescription).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if tabType == "lab" {
			templ_7745c5c3_Err = TestTable(data.Composition.MtxReport.LabResults).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PrescriptionTable(results model.Prescription) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col border-[1px] border-b-4 border-nhs-purple-2 w-full mb-6 p-4 gap-2\"><div>Patient's next blood test for <span class=\"font-medium\">Methotrexate</span> is due on <span class=\"text-red-500 font-medium\">27/3/24</span>. Send a reminder to the patient to book an appointment?</div><div class=\"flex gap-4\"><button hx-get=\"/hx/sms-reminder\" hx-swap=\"none\" class=\"px-4 py-2 rounded-md w-fit bg-nhs-green-2 text-white border-b-2 border-b-white\" type=\"button\">Send SMS</button> <button class=\"px-4 py-2 rounded-md w-fit bg-nhs-green-2 text-white border-b-2 border-b-white\" type=\"button\">Send via App</button></div></div><table class=\"table-fixed w-full mt-1 text-sm rounded-lg\"><thead><tr class=\"p-2 font-semibold bg-zinc-300 tracking-wide\"><td class=\"p-2 border border-zinc-400 text-zinc-600\">Medication</td><td class=\"p-2 border border-zinc-400 text-zinc-600\">Date started</td><td class=\"p-2 border border-zinc-400 text-zinc-600\">Date ended</td><td class=\"p-2 border border-zinc-400 text-zinc-600\">Dose</td><td class=\"p-2 border border-zinc-400 text-zinc-600\">Prescribed by</td><td class=\"p-2 border border-zinc-400 text-zinc-600\">Dispensed by</td><td class=\"p-2 border border-zinc-400 text-zinc-600\">Last blood test</td></tr></thead> <tbody><tr class=\"align-top\"><td class=\"p-2 border border-zinc-400\">Methotrexate</td><td class=\"p-2 border border-zinc-400\">1/2/24</td><td class=\"p-2 border border-zinc-400\">-</td><td class=\"p-2 border border-zinc-400\">15mg</td><td class=\"p-2 border border-zinc-400\">Dr Andrew</td><td class=\"p-2 border border-zinc-400\">Esther</td><td class=\"p-2 border border-zinc-400 text-zinc-600\">13/3/24</td></tr></tbody></table>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TestTable(results model.LabResults) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col border-[1px] border-b-4 border-red-500 w-full mb-6 p-4 gap-2\"><div>Patient's blood test results are not within the normal range. Send an alert to the patient to stop taking the medication immediately.</div><div class=\"flex gap-4\"><button hx-get=\"/hx/sms-alert\" hx-swap=\"none\" class=\"px-4 py-2 rounded-md w-fit bg-nhs-green-2 text-white border-b-2 border-b-white\" type=\"button\">Send SMS</button> <button class=\"px-4 py-2 rounded-md w-fit bg-nhs-green-2 text-white border-b-2 border-b-white\" type=\"button\">Send via App</button></div></div><table class=\"table-fixed w-full mt-1 text-sm rounded-lg\"><thead><tr class=\"p-2 font-semibold bg-zinc-300 tracking-wide\"><td class=\"p-2 border border-zinc-400 text-zinc-600 w-1/6\">Lab</td><td class=\"p-2 border border-zinc-400 text-zinc-600 w-1/6\">Reference Range</td><td class=\"p-2 border border-zinc-400 text-zinc-600 w-1/6\">1/2/24 - 8.00am</td><td class=\"p-2 border border-zinc-400 text-zinc-600 w-1/6\">14/2/24 - 8.00am</td><td class=\"p-2 border border-zinc-400 text-zinc-600 w-1/6\">28/2/24 - 8.00am</td><td class=\"p-2 border border-zinc-400 text-zinc-600 w-1/6\">13/3/24 - 8.00am</td></tr></thead> <tbody>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestSection("Full Blood Count").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Hemoglobin", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].FullBloodCount[0].Hb[0].AnalyteResult[0].Magnitude, results[0].FullBloodCount[0].Hb[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("White blood cell", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].FullBloodCount[0].TotalWhiteCellCount[0].AnalyteResult[0].Magnitude, results[0].FullBloodCount[0].TotalWhiteCellCount[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Lymphocyte", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].FullBloodCount[0].Platelets[0].AnalyteResult[0].Magnitude, results[0].FullBloodCount[0].Platelets[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Neutrophils", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].FullBloodCount[0].Neutrophils[0].AnalyteResult[0].Magnitude, results[0].FullBloodCount[0].Neutrophils[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Platelets", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].FullBloodCount[0].Lymphocytes[0].AnalyteResult[0].Magnitude, results[0].FullBloodCount[0].Lymphocytes[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestSection("Liver Function Test").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("AST", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].LiverFunctionTests[0].Ast[0].AnalyteResult[0].Magnitude, results[0].LiverFunctionTests[0].Ast[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("ALT", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].LiverFunctionTests[0].Alt[0].AnalyteResult[0].Magnitude, results[0].LiverFunctionTests[0].Alt[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestSection("Urea & Electrolytes").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Creatinine", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].UreaAndElectrolytes[0].Creatinine[0].AnalyteResult[0].Magnitude, results[0].UreaAndElectrolytes[0].Creatinine[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Sodium", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].UreaAndElectrolytes[0].Sodium[0].AnalyteResult[0].Magnitude, results[0].UreaAndElectrolytes[0].Sodium[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Potassium", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].UreaAndElectrolytes[0].Potassium[0].AnalyteResult[0].Magnitude, results[0].UreaAndElectrolytes[0].Potassium[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TestRow("Urea", results[0].FullBloodCount[0].Hb[0].ReferenceRangeGuidance[0], fmt.Sprintf("%v %v", results[0].UreaAndElectrolytes[0].Urea[0].AnalyteResult[0].Magnitude, results[0].UreaAndElectrolytes[0].Urea[0].AnalyteResult[0].Unit)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TestSection(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"p-2 font-medium bg-zinc-200\"><td class=\"p-2 w-1/6 text-zinc-600 border border-zinc-400\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/clinician/clinician.templ`, Line: 178, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 0; i < 5; i++ {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"p-2 border border-zinc-400 w-1/6\"></td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TestRow(k, refRange, v string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"align-top\"><td class=\"p-2 font-medium border border-zinc-400\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(k)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/clinician/clinician.templ`, Line: 187, Col: 56}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-2 border border-zinc-400 w-1/6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(refRange)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/clinician/clinician.templ`, Line: 188, Col: 57}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 0; i < 4; i++ {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"p-2 border border-zinc-400 w-1/6\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(v)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/clinician/clinician.templ`, Line: 190, Col: 51}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
