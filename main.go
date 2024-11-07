package main

import (
	"converterreportcsvtojson/app/adapter"
	"converterreportcsvtojson/app/legacy"
	"flag"
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

	switch *outputType{
	case "json":
		reportAdapter := adapter.NewJSONReportAdapter(legacyReport, fullPath)
		err = reportAdapter.GenerateReport()
		if err != nil {
			logger.Error("Erro ao gerar relatório JSON:", err, zap.String("Function", "Main"))
			return
		}
		logger.Info("Relatório JSON gerado com sucesso", zap.String("Function", "Main"))
		break
	case "parquet":
		reportAdapter := adapter.NewParquetReportAdapter(legacyReport, fullPath)
		err = reportAdapter.GenerateReport()
		if err != nil {
			logger.Error("Erro ao gerar relatório Parquet:", err, zap.String("Function", "Main"))
			return
		}
		logger.Info("Relatório Parquet gerado com sucesso", zap.String("Function", "Main"))
		break
	default:
		logger.Info("Parametro outputtype não reconhecido.", zap.String("Function", "Main"))
		break
	}	
}
