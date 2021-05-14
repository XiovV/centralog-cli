package main

import (
	"fmt"
	"github.com/XiovV/centralog/sorter"
)


func main() {
	//portainer := collector.New("75fe3da0cd24fb7e1517f7ddec5b7ea9b4c89427231288c93157fd88f4b3393c", "portainer")
	//
	//logs := portainer.CollectLogs()

	logs := sorter.GetNewBatch()



	for _, log := range logs {
		fmt.Printf("%+v\n", log)
	}

}
