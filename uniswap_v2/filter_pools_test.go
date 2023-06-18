package uniswap_v2

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterPoolsAndWriteToFile(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	pools, err := FilterPoolsAndWriteToFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//print length of pools
	fmt.Println("len(pools): ", len(pools))
}
