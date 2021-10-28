package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//A struct with 2 fields
type response1 struct {
	Page   int
	Fruits []string
}

//A struct with 2 json fields
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	//Converting bool to json
	bolB, _ := json.Marshal(true) // json.Marshal returns 2 values i.e. 1 actual value and a blank operator to hold error if any
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher") //string comes with ""
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} //string array comes with ["",..]
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7} //map comes as same way {"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//Convert response1 to json
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// response1 to json
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//Converting a json to go type
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`) //sample input

	var dat map[string]interface{}                    //creating a map variable to hold the json values
	if err := json.Unmarshal(byt, &dat); err != nil { //Throw an error if issue with conversion
		panic(err)
	}
	fmt.Println(dat)

	//fetch value of num variable and convert that to float and assign to var
	num := dat["num"].(float64)
	fmt.Println(num)
	//fetch value of strs and convert it to string
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//Convert json to go type using Json Unmarshal
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//You can also stream json encodings directly to OS writers and Stdout or even http
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
