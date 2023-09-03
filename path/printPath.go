package path

import "fmt"

func PrintPath(path Path) {
	fmt.Println("Path id", path.Id)
	for i := 0; i < len(path.Pools); i++ {
		fmt.Println("Pool ", i)
		fmt.Println("Pool: ", path.Pools[i].GetAddress().String())
		fmt.Println("Token0: ", path.Pools[i].GetTokens()[0].String())
		fmt.Println("Token1: ", path.Pools[i].GetTokens()[1].String())
	}
}
