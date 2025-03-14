package cmd

import datamodel "github.com/munnaMia/Combi-Tracker/Model"

type Application struct {
	dataModel datamodel.Model
	commands  []string
}
