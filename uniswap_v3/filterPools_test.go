package uniswap_v3

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

	FilterPoolsAndWriteToFile()
}
