package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Page struct {
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	Timezone  string     `json:"timezone"`
	Incidents []Incident `json:"incidents"`
}

type Incident struct {
	Id              string           `json:"id"`
	Name            string           `json:"name"`
	CreatedAt       string           `json:"created_at"`
	ResolvedAt      string           `json:"resolved_at"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
}

type IncidentUpdate struct {
	Id                 string              `json:"id"`
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

	// fmt.Println(string(body))
	var r Page
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	for _, incident := range r.Incidents {
		log.Printf("incident_id: %s,\n name: %s,\n created_at: %s,\n resolved_at: %s,\n status: %s,\n impact: %s,\n shortlink: %s,\n component_id: %s,\n component_name: %s\n\n\n", incident.Id, incident.Name, incident.CreatedAt, incident.ResolvedAt, incident.IncidentUpdates[0].Status, incident.Impact, incident.Shortlink, incident.IncidentUpdates[0].AffectedComponents[0].Code, incident.IncidentUpdates[0].AffectedComponents[0].Name)
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
