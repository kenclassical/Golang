package main

type MyModel struct {
	Index int    `json:"index"`
	value string `value:"value"`
}

func convertToMyModel(array []string) []MyModel {
	var models []MyModel
	for index, value := range array {
		models = append(models, MyModel{index, value})
	}
	return models
}

func main() {
	array := []string{"red", "green", "blue", "yellow", "white"}
	var results = convertToMyModel(array)
	println(results[0].Index)
	println(results[0].value)
	println(results[1].Index)
	println(results[1].value)
	println(results[2].Index)
	println(results[2].value)
	println(results[3].Index)
	println(results[3].value)
	println(results[4].Index)
	println(results[4].value)
}
