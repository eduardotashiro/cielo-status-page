package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type StatusResponse struct {
	Page      Page       `json:"page"`
	Incidents []Incident `json:"incidents"`
}

type Page struct {
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	Timezone  string     `json:"time_zone"`
	Incidents []Incident `json:"incidents"`
}

type Incident struct {
	Id              string           `json:"id"`
	Name            string           `json:"name"`
	CreatedAt       string           `json:"created_at"`
	ResolvedAt      string           `json:"resolved_at"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
}

type IncidentUpdate struct {
	Id                 string              `json:"incident_id"`
	Status             string              `json:"status"`
	AffectedComponents []AffectedComponent `json:"affected_components"`
}

type AffectedComponent struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("https://status.cielo.com.br/api/v2/incidents.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r StatusResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	for _, incident := range r.Incidents {
		if CheckIncident(incident) {
			fmt.Printf("name:\t%s\nurl:\t%s\ntimezone:\t%s\nincident_id:\t%s\nnincident_name:\t%s\nincident_created_at:\t%s\nincident_resolved_at:\t%s\nstatus:\t%s\nimpact:\t%s\nshortlink:\t%s\ncomponent_id:\t%s\ncomponent_name:\t%s\n\n\n", r.Page.Name, r.Page.Url, r.Page.Timezone, incident.Id, incident.Name, incident.CreatedAt, incident.ResolvedAt, incident.IncidentUpdates[0].Status, incident.Impact, incident.Shortlink, incident.IncidentUpdates[0].AffectedComponents[0].Code, incident.IncidentUpdates[0].AffectedComponents[0].Name)
		}
	}
}

// http.Client{Timeout: ...}
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
