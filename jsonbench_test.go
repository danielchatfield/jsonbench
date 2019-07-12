package jsonbench

import (
	"encoding/json"
	"log"
	"testing"
)

func BenchmarkJSONUnmarshal(b *testing.B) {
	f := NewFoo()
	data, err := json.Marshal(f)
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		var f Foo
		if err := json.Unmarshal(data, &f); err != nil {
			log.Fatal(err)
		}
	}
}