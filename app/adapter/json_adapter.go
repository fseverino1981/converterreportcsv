package adapter

import (
	"converterreportcsvtojson/app/legacy"
	"encoding/json"
	"io/ioutil"
	"strings"
	"os"
    "path/filepath"
)


type JSONReportAdapter struct{
	legacyReport *legacy.LegacyReport
	outputPath    string
}

func NewJSONReportAdapter(legacyReport *legacy.LegacyReport, outputPath string) *JSONReportAdapter{
	return &JSONReportAdapter{legacyReport: legacyReport,
		outputPath: outputPath}
}

func (a *JSONReportAdapter) GenerateReport() error{
	
	csvData, err := a.legacyReport.GeneratorCSVReport()
	if err != nil{
		return err
	}

	lines := strings.Split(csvData, "\n")
	if len(lines) < 2{
		return err
	}

	headers := strings.Split(lines[0], ",")

	var jsonData []map[string]string

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		values := strings.Split(line, ",")
		rowData := make(map[string]string)
		for i, header := range headers {
			rowData[header] = values[i]
		}
		jsonData = append(jsonData, rowData)
	}

	jsonBytes, err := json.MarshalIndent(jsonData, "", " ")
	if err != nil {
		return err
	}

	if err := a.outputFileCreate(); err != nil{
		return err
	}
	
	if err := ioutil.WriteFile(a.outputPath,jsonBytes, 0644); err != nil{
		return err
	}

	return nil
}

func (a JSONReportAdapter) outputFileCreate()error{
	if err := os.MkdirAll(filepath.Dir(a.outputPath), os.ModePerm); err != nil {
        return err
    }

    file, err := os.Create(a.outputPath)
    if err != nil {
        return err
    }
    defer file.Close()
	
	return nil
}