package main

import (
	"converterreportcsvtojson/app/legacy"
	"converterreportcsvtojson/app/strategy"
	"flag"
	"fmt"
	"os"

	toolkit "github.com/fseverino1981/golang-toolkit"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

	var logger toolkit.Logger

	inputFile := flag.String("inputfile", "csv_report.", "Input File Path - default: csv_report")
	outputFile := flag.String("outputfile", "output_report", "Output File Path - default: output_report")
	outputType := flag.String("outputtype", "json", "Output type report - options: json parquet" )

	flag.Parse()

	err := godotenv.Load()
	if err != nil{
		logger.Error("Erro ao carregar variveis de ambiente.", err, zap.String("Function", "Main"))
		return
	}

	legacyReport := legacy.NewLegacyReport(*inputFile)


	fullPath := os.Getenv("PATH_OUTPUT")+ *outputFile

	report, err := strategy.ReturnStrategy(*outputType)
	if err != nil{
		logger.Error(fmt.Sprintf("Erro ao gerar relatório de saído no formato %s.", *outputType), err, zap.String("Function","Main"))
	}

	report.GenerateReport(legacyReport, fullPath)
}
