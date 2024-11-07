package legacy

type ReportGenerator interface{
	GenerateReport() (error)
}