package json

import "encoding/json"

func Unmarshal[Out any, In string | []byte](in In) Out {
	var out Out

	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		panic(err)
	}

	return out
}
