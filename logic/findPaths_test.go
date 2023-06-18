package logic

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func printPool(pool interfaces.PoolInterface) {
// 	fmt.Println("POOL")
// 	fmt.Println("pool type", pool.GetType())
// 	fmt.Println("pool.Address", pool.GetAddress().Hex())
// 	fmt.Println("pool token0", pool.GetTokens()[0].Hex())
// 	fmt.Println("pool token1", pool.GetTokens()[1].Hex())

// }

// func TestPrintAllPools(t *testing.T) {
// 	err := os.Chdir("../")
// 	fmt.Println("err: ", err)
// 	assert.NoError(t, err)
// 	pools, err := GetFilteredPools()
// 	assert.NoError(t, err)

// 	for i, pool := range pools {
// 		fmt.Println("i: ", i)
// 		printPool(pool)
// 	}
// }

func TestFindAllPathsAndWriteToFile(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	paths, err := FindAllPathsAndWriteToFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	fmt.Println("len(paths): ", len(paths))
}

func TestReadPathsFromFile(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//print a path

	PrintPath(paths[100])

	fmt.Println("len(paths): ", len(paths))
}
