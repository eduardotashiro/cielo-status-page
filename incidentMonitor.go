package main

import (
	"slices"
)

func CheckIncident(i Incident) bool {
	if i.Impact != "critical" && i.Impact != "major" {
		return false
	}
	if len(i.IncidentUpdates) == 0 || len(i.IncidentUpdates[0].AffectedComponents) == 0 {
		return false
	}
	//gxk94rbthwvb - Autorizador Físico removido
	componentesMonitorados := []string{
		"tpgr8ylqdt1j", //Antifraude
		"3zj9v894hf9y", //Autorizador Digital
		"zn71xcn0fgnl", //Bandeiras / Emissores / Parceiros
		"x5zhg1pgmzhw", //Cancelamento
		"6129pxjwtfvn", //e-Commerce - Parceiros
		"mh0f6fyn4xkh", //e-Commerce - PSP
		"dz8yg5x021b9", //Liquidação Financeira (PagFor)
		"n0n3lzdwjsdj", //Pix
	}
	for _, c := range i.IncidentUpdates[0].AffectedComponents {
		if slices.Contains(componentesMonitorados, c.Code) {
			return true
		}
	}
	return false
}
