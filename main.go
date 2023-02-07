package main

import (
	"selefra-demo/internal/model"
	"selefra-demo/internal/utils"
)

func main() {
	utils.InitSetting.Do(utils.Init)
	model.InitDB.Do(model.InitDb)
}
