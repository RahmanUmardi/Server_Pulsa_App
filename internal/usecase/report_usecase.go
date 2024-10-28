package usecase

import (
	"fmt"
	"reflect"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"
	"server-pulsa-app/internal/shared/custom"

	"github.com/xuri/excelize/v2"
)

type ReportUseCase interface {
	FindAllTransactions(userId, startDate, endDate string) error
}

type reportUseCase struct {
	repo repository.ReportRepository
	log  *logger.Logger
}

func (r *reportUseCase) FindAllTransactions(userId, startDate, endDate string) error {
	r.log.Info("Starting to retrive report of all transactions in the usecase layer", nil)

	reportSlice, err := r.repo.List(userId, startDate, endDate)
	if err != nil {
		return err
	}

	f := excelize.NewFile()
	defer func() error {
		if err := f.Close(); err != nil {
			return err
		}
		return nil
	}()

	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return err
	}

	t := reflect.TypeOf(custom.ReportResp{})

	fields := make([]string, t.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = t.Field(i).Name
	}

	columns := make([]string, 0, t.NumField())
	var startingASCIINumber int = 64

	for i := 1; i <= t.NumField(); i++ {
		columns = append(columns, string(rune(startingASCIINumber+i)))
	}

	// Set value of a cell.
	for i := 0; i < len(fields); i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[i], 1), fields[i])
	}

	style, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"87CEFA"}, Pattern: 1},
	})
	if err != nil {
		return err
	}
	err = f.SetCellStyle("Sheet1", "A1", "B1", style)
	if err != nil {
		return err
	}

	for i := 0; i < len(reportSlice); i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[0], i+2), reportSlice[i].ProviderName)
		f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[1], i+2), reportSlice[i].Count)

		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[0], i+2), reportSlice[i].TransactionsId)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[1], i+2), reportSlice[i].CustomerName)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[2], i+2), reportSlice[i].DestinationNumber)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[3], i+2), reportSlice[i].TransactionDate)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[4], i+2), reportSlice[i].IdUser)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[5], i+2), reportSlice[i].Username)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[6], i+2), reportSlice[i].Role)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[7], i+2), reportSlice[i].IdMerchant)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[8], i+2), reportSlice[i].NameMerchant)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[9], i+2), reportSlice[i].Address)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[10], i+2), reportSlice[i].TransactionDetailId)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[11], i+2), reportSlice[i].ProductId)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[12], i+2), reportSlice[i].NameProvider)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[13], i+2), reportSlice[i].Nominal)
		// f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[14], i+2), reportSlice[i].Price)
		style, err = f.NewStyle(&excelize.Style{
			Fill: excelize.Fill{Type: "pattern", Color: []string{"90EE90"}, Pattern: 1},
		})
		if err != nil {
			return err
		}
		err = f.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", columns[0], i+2), fmt.Sprintf("%s%d", columns[1], i+2), style)
		if err != nil {
			return err
		}
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("./internal/assets/Report.xlsx"); err != nil {
		return err
	}

	return nil
}

func NewReportUseCase(repo repository.ReportRepository, log *logger.Logger) ReportUseCase {
	return &reportUseCase{repo: repo, log: log}
}
