package json

import "encoding/json"

// Unmarshal unmarshals the given JSON string or byte slice into the given type.
// It delegates to the standard library's json.Unmarshal function, leveraging
// generics to figure out the type of the output. If the unmarshaling fails,
// it panics.
func Unmarshal[Out any, In string | []byte](in In) Out {
	var out Out

	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		panic(err)
	}

	return out
}
