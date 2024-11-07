package legacy

import (
	"io/ioutil"
)

type LegacyReport struct {
	FilePath string
}

func NewLegacyReport(filePath string) *LegacyReport{
	return &LegacyReport{FilePath: filePath}
}

func (l *LegacyReport) GeneratorCSVReport() (string,error) {
	data, err := ioutil.ReadFile(l.FilePath)
	if err != nil{
		return "", err
	}

	return string(data), nil
}