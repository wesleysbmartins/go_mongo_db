package main

import (
	"fmt"
	"go_mongo_db/internal/adapters"
	"go_mongo_db/internal/entities"
	"go_mongo_db/internal/repository"
)

func init() {
	mongo := adapters.MongoDb{}
	mongo.Connect()
}

func main() {
	reportRepository := repository.ReportRepository{}

	reports := &[]entities.Report{}

	reportA := entities.Report{
		Title:       "Resultado Mensal",
		Content:     "Sucesso com as vendas acima da média, possível estabelecer nova meta.",
		Responsible: "João da Silva",
	}

	reportB := entities.Report{
		Title:       "Hardenização de Maquinás",
		Content:     "Falha total, todas as maquinas devem passar por vistoria.",
		Responsible: "Gustavo de Souza",
	}

	*reports = append(*reports, reportA, reportB)

	for _, report := range *reports {
		err := reportRepository.Create(report)
		if err != nil {
			fmt.Println("CREATE ERROR\n", report)
		} else {
			fmt.Println("CREATE SUCCESS!")
		}
	}

	allReports, err := reportRepository.Find(nil)

	if err != nil {
		fmt.Println("FIND ERRROR\n", err)
	} else {
		fmt.Println("FIND SUCCESS!\nLENGTH: ", len(*allReports))
	}

	for _, report := range *allReports {
		newReport := &entities.Report{
			Responsible: "Wesley Martins",
			Content:     report.Content,
			Title:       report.Title,
		}

		err := reportRepository.Update(report.Id, *newReport)
		if err != nil {
			fmt.Println("UPDATE ERROR\n", err)
		} else {
			fmt.Println("UPDATE SUCCESS!")
		}
	}

	allReports, err = reportRepository.Find(nil)

	if err != nil {
		fmt.Println("FIND ERRROR\n", err)
	} else {
		fmt.Println("FIND SUCCESS!\nLENGTH: ", len(*allReports))
	}

	for _, report := range *allReports {

		err := reportRepository.Delete(report.Id)
		if err != nil {
			fmt.Println("DELETE ERROR\n", err)
		} else {
			fmt.Println("DELETE SUCCESS!")
		}
	}

	allReports, err = reportRepository.Find(nil)

	if err != nil {
		fmt.Println("FIND ERRROR\n", err)
	} else {
		fmt.Println("FIND SUCCESS!\nLENGTH: ", len(*allReports))
	}
}
