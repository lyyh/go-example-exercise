package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	json.Unmarshal(byt, &dat)
	fmt.Println(dat)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	json.Unmarshal([]byte(str), &dat)
	fmt.Println(dat)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"app": 5, "name": 1}
	enc.Encode(d)
}
