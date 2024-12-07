package example

import (
	"encoding/json"
)

func UnmarshalExample1() {
	person1Json := []byte(`{"name":"John","family":"lotfi","age ":32}`)

	var person1 = Persons{}
	err := json.Unmarshal(person1Json, &person1)
	if err != nil {
		panic(err)
	}

	println(person1)
}
