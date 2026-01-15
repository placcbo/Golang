package main

import (
	"fmt"
	"work_reservation/utils"
)

func main() {

	// Accuracy
	workerAverageAccuracy, err := utils.WorkerAccuracyAverage(86, 90, 80) // partially working
	fmt.Printf("Accruracy: %.2f, %v \n ", workerAverageAccuracy, err)

	//Average Tpt

	totaWorkerlTpt, err := utils.Tpt(200, 70, 100)
	fmt.Println("Average Tpt:", totaWorkerlTpt, nil)

	workerScore, err := utils.Score(totaWorkerlTpt, workerAverageAccuracy)

	fmt.Println(workerScore, nil)
}
