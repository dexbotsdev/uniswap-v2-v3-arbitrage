package executor

import (
	"log"
	"math/big"
	"os"
)

func logProfitablePath(pathExecutionValues PathExecutionValues, blockNumber uint64, revenue *big.Int, minGasCost uint64, profit *big.Int) {
	//path execution values
	//block number
	//time stamp
	//revenue
	//min gas cost
	//profit

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This log message will go to a file.")

}
