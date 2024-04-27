package server

import (
	"shards-backend/db"
	"shards-backend/utils"
)

// Init <function>
// is used to initialize server and all the corresponding services such as DB, Utils, Workers
func Init() {
	// utils
	utils.InitEnvVars()

	// services
	db.InitService()
	db.GetClient().FillSeedsInformation() // a bit ugly but everything is going to work, only needed to fill seeds information

	// workers
	//do stuff

	r := NewRouter()
	// r.Run(":6969")
	err := r.Run(":80")
	if err != nil {
		return 
	}
}