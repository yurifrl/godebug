package godebug

import (
	"encoding/json"
	"fmt"
)

// Println ...
func Println(i interface{}) {
	j, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))
}
