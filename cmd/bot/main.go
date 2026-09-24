package main

import (
	"fmt"
	"log"

	"github.com/eduardotashiro/cielo-status-page/internal/client"
)

func main() {
	r, err := client.GetIncidents()
	if err != nil {
		log.Fatalf("Erro ao obter incidentes: %v", err)
	}

	for _, incident := range r.Incidents {
		if incident.IsRelevant() {
			fmt.Printf("name:\t%s\nurl:\t%s\ntimezone:\t%s\nincident_id:\t%s\nnincident_name:\t%s\nincident_created_at:\t%s\nincident_resolved_at:\t%s\nstatus:\t%s\nimpact:\t%s\nshortlink:\t%s\ncomponent_id:\t%s\ncomponent_name:\t%s\n\n\n", r.Page.Name, r.Page.Url, r.Page.Timezone, incident.Id, incident.Name, incident.CreatedAt, incident.ResolvedAt, incident.IncidentUpdates[0].Status, incident.Impact, incident.Shortlink, incident.IncidentUpdates[0].AffectedComponents[0].Code, incident.IncidentUpdates[0].AffectedComponents[0].Name)
		}
	}
}

// incident := r.Incidents[0]

/*
"id": "n0n3lzdwjsdj",
"name": "Pix",
----------------------
"code": "dz8yg5x021b9",
"name": "Transacional - Liquidação Financeira (PagFor)",
----------------------
"code": "mh0f6fyn4xkh",
"name": "Transacional - e-Commerce - PSP",
----------------------
"code": "6129pxjwtfvn",
"name": "Transacional - e-Commerce - Parceiros",
----------------------
"code": "x5zhg1pgmzhw",
"name": "Transacional - Cancelamento",
----------------------
"code": "zn71xcn0fgnl",
"name": "Transacional - Bandeiras / Emissores / Parceiros",
----------------------
"code": "gxk94rbthwvb", 					 desnecessario ?
"name": "Transacional - Autorizador Físico", desnecessario ?
----------------------
"code": "3zj9v894hf9y",						  full necessario !
"name": "Transacional - Autorizador Digital", full necessario !
----------------------
"id": "tpgr8ylqdt1j",
"name": "Antifraude",
----------------------

	{
	  "id": "3kpqfvh4jxlx",
	  "name": "Transacional",
	  "status": "operational",
	  "created_at": "2025-02-17T16:01:11.549-03:00",
	  "updated_at": "2025-02-17T16:01:11.549-03:00",
	  "position": 14,
	  "description": null,
	  "showcase": false,
	  "start_date": null,
	  "group_id": null,
	  "page_id": "mrnkr4vjrl2d",
	  "group": true,
	  "only_show_if_degraded": false,
	  "components": [
	    "tpgr8ylqdt1j",
	    "3zj9v894hf9y",
	    "gxk94rbthwvb", X
	    "zn71xcn0fgnl",
	    "x5zhg1pgmzhw",
	    "6129pxjwtfvn",
	    "mh0f6fyn4xkh",
	    "dz8yg5x021b9",
	    "n0n3lzdwjsdj"
	  ]
	},
*/

// name:   Cielo Status Page
// url:    https://status.cielo.com.br
// timezone:       America/Sao_Paulo
// incident_id:    8l7k6ghr6b75
// nincident_name: Indisponibilidade nos pagamentos PIX QrCode
// incident_created_at:    2026-08-23T07:06:51.423-03:00
// incident_resolved_at:   2026-08-23T09:13:46.317-03:00
// status: resolved
// impact: critical
// shortlink:      https://stspg.io/0lwmp3rn9qjc
// component_id:   n0n3lzdwjsdj
// component_name: Transacional - Pix
