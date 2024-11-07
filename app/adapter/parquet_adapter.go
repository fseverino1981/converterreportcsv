package adapter

import (
	"encoding/csv"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"github.com/xitongsys/parquet-go/writer"
	"github.com/xitongsys/parquet-go-source/local"
	"converterreportcsvtojson/app/legacy"
)


type ParquetReportAdapter struct{
	legacyReport *legacy.LegacyReport
	outputPath    string
}

type Person struct {
	ID        int32  `parquet:"name=id, type=INT32"`
	Nome      string `parquet:"name=nome, type=BYTE_ARRAY, convertedtype=UTF8"`
	Sobrenome string `parquet:"name=sobrenome, type=BYTE_ARRAY, convertedtype=UTF8"`
	Idade     int32  `parquet:"name=idade, type=INT32"`
	Email     string `parquet:"name=email, type=BYTE_ARRAY, convertedtype=UTF8"`
	Cidade    string `parquet:"name=cidade, type=BYTE_ARRAY, convertedtype=UTF8"`
	Estado    string `parquet:"name=estado, type=BYTE_ARRAY, convertedtype=UTF8"`
}

func NewParquetReportAdapter(legacyReport *legacy.LegacyReport, outputPath string) *ParquetReportAdapter{
	return &ParquetReportAdapter{legacyReport: legacyReport,
		outputPath: outputPath}
}

func (a *ParquetReportAdapter) GenerateReport() error {
    csvData, err := a.legacyReport.GeneratorCSVReport()
    if err != nil {
        return fmt.Errorf("Erro ao obter relatório CSV: %v", err)
    }

    reader := csv.NewReader(strings.NewReader(csvData))

    // Leia o cabeçalho
    header, err := reader.Read()
    if err != nil {
        return fmt.Errorf("Erro ao ler o cabeçalho do CSV: %v", err)
    }

    // Abra o arquivo Parquet para escrita
    fw, err := local.NewLocalFileWriter(a.outputPath)
    if err != nil {
        log.Fatalf("Erro ao criar o arquivo Parquet: %v", err)
        return err
    }
    defer fw.Close()

    pw, err := writer.NewParquetWriter(fw, new(Person), 4)
    if err != nil {
        log.Fatalf("Erro ao criar writer do Parquet: %v", err)
        return err
    }
    defer pw.WriteStop()

    pw.RowGroupSize = 128 * 1024 * 1024
    pw.PageSize = 8 * 1024

    // Leia as linhas e converta para o struct User
    for {
        record, err := reader.Read()
        if err != nil {
            break // Fim do arquivo
        }

        person := Person{}

        // Preencha os campos do struct com os dados da linha
        for i, value := range record {
            switch header[i] {
            case "ID":
                id, _ := strconv.Atoi(value)
                person.ID = int32(id)
            case "Nome":
                person.Nome = value
            case "Sobrenome":
                person.Sobrenome = value
            case "Idade":
                idade, _ := strconv.Atoi(value)
                person.Idade = int32(idade)
            case "Email":
                person.Email = value
            case "Cidade":
                person.Cidade = value
            case "Estado":
                person.Estado = value
            }
        }

        if err := pw.Write(person); err != nil {
            return err
        }
    }

    return nil
}


func detectType(value string) reflect.Type {
	if _, err := strconv.Atoi(value); err == nil {
		return reflect.TypeOf(int32(0))
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return reflect.TypeOf(float64(0))
	}
	return reflect.TypeOf("")
}

func convertValue(value string, dataType reflect.Type) interface{} {
	switch dataType.Kind() {
	case reflect.Int32:
		v, _ := strconv.Atoi(value)
		return int32(v)
	case reflect.Float64:
		v, _ := strconv.ParseFloat(value, 64)
		return v
	default:
		return value
	}
}
