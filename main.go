package main

import (
	"encoding/json"
	"fmt"
	storage "github.com/GrooveCommunity/glib-cloud-storage/gcp"
	"github.com/GrooveCommunity/glib-noc-event-structs/gcs"
	"log"
	"time"
)

const JIRA_DATE_FORMAT = "2006-01-02T15:04:05.999-0700"

func main() {
	var metrics gcs.IssuesMetric

	payload := storage.GetObject("noc-paygo","jira-issues.json")

	json.Unmarshal(payload,&metrics)

	var (
		aguardando_sd,
		analise_sd,
		analise_squad,
		aguardando_atendimento,
		aguardando_fornecedor,
		aguardando_solicitante,
		resolvido,
		encerrado int
	)

	for _, issue := range metrics.Issues {
		for _, state := range issue.States {
			if state.Status == "Em analise SD" {
				if state.ChangeDate != "" {
					changeDateState := convertStrDate(state.ChangeDate)
					createDate := convertStrDate(state.CreateDate)

					duration := changeDateState.Sub(createDate)
					log.Println(duration.Minutes())

				}
			}
		}

		/*state := issue.States[len(issue.States)-1]

		switch state.Status {
			case "Aguardando SD":
				aguardando_sd += 1
				break
			case "Em analise SD":
				analise_sd += 1
				break
			case "Em analise Squad":
				analise_squad += 1
				break
			case "Aguardando Atendimento":
				aguardando_atendimento += 1
				break
			case "AGUARDANDO FORNECEDOR":
				aguardando_fornecedor += 1
				break
			case "AGUARDANDO SOLICITANTE":
				aguardando_solicitante += 1
				break
			case "Chamado Resolvido":
                resolvido += 1
                break
			case "Encerrado":
				encerrado += 1
				//TODO retirar da lista de issues
				break
			default:
				break
			}*/
	}

	fmt.Printf("Resolvido: %i\n", resolvido)
	fmt.Printf("Aguardando SD: %i\n", aguardando_sd)
	fmt.Printf("Em analise Squad: %i\n", analise_squad)
	fmt.Printf("Em analise SD: %i\n", analise_sd)
	fmt.Printf("Aguardando Atendimento: %i\n", aguardando_atendimento)
	fmt.Printf("AGUARDANDO FORNECEDOR: %i\n", aguardando_fornecedor)
	fmt.Printf("AGUARDANDO SOLICITANTE: %i\n", aguardando_solicitante)
	fmt.Printf("Encerrado: %i\n", encerrado)
}

func convertStrDate(dateStr string) time.Time {
	dateStruct, err := time.Parse(JIRA_DATE_FORMAT, dateStr)

	if err != nil {
		panic(err)
	}

	return dateStruct
}