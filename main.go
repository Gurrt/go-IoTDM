package main

import (
	"IoTDM"
	"fmt"
)

func main() {
	dbHandle, err := IoTDM.GetDatabaseConnection("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err)
		return
	}
	runtime := IoTDM.RuntimeInfo{DB: dbHandle}
	runtime.StartAPI()
}
