package domain

import "slices"

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

// gxk94rbthwvb - Autorizador Físico removido
var componentesMonitorados = []string{
	"tpgr8ylqdt1j", //Antifraude
	"3zj9v894hf9y", //Autorizador Digital
	"zn71xcn0fgnl", //Bandeiras / Emissores / Parceiros
	"x5zhg1pgmzhw", //Cancelamento
	"6129pxjwtfvn", //e-Commerce - Parceiros
	"mh0f6fyn4xkh", //e-Commerce - PSP
	"dz8yg5x021b9", //Liquidação Financeira (PagFor)
	"n0n3lzdwjsdj", //Pix
}

func (i *Incident) IsRelevant() bool {
	if i.Impact != "critical" && i.Impact != "major" {
		return false
	}
	if len(i.IncidentUpdates) == 0 || len(i.IncidentUpdates[0].AffectedComponents) == 0 {
		return false
	}
	for _, c := range i.IncidentUpdates[0].AffectedComponents {
		if slices.Contains(componentesMonitorados, c.Code) {
			return true
		}
	}
	return false
}
