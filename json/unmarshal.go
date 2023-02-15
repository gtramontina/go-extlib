package json

import (
	"encoding/json"
	"fmt"
)

// Unmarshal unmarshals the given JSON string or byte slice into the given type.
// It delegates to the standard library's json.Unmarshal function, leveraging
// generics to figure out the type of the output. If the unmarshaling fails,
// it panics. See also TryUnmarshal.
func Unmarshal[Out any, In string | []byte](in In) Out {
	out, err := TryUnmarshal[Out, In](in)
	if err != nil {
		panic(err)
	}

	return out
}

// TryUnmarshal unmarshals the given JSON string or byte slice into the given
// type. It delegates to the standard library's json.Unmarshal function,
// leveraging generics to figure out the type of the output. If the unmarshaling
// fails, it returns the error. See also Unmarshal.
func TryUnmarshal[Out any, In string | []byte](in In) (Out, error) {
	var out Out

	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		return out, fmt.Errorf("failed unmarshaling json: %w", err)
	}

	return out, nil
}
