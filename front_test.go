package front_test

import (
	"fmt"
	"log"

	"github.com/cmr-informatik/front"
)

var markdown = []byte(`---
title: Ferrets
authors:
- Tobi
- Loki
- Jane
---
Some content here, so
interesting, you just
want to keep reading.`)

type article struct {
	Title   string
	Authors []string
}

func ExampleUnmarshal() {
	var a article

	content, err := front.Unmarshal(markdown, &a)
	if err != nil {
		log.Fatalf("error unmarshalling: %s", err)
	}

	fmt.Printf("%#v\n", a)
	fmt.Printf("%s\n", string(content))
	// Output:
	// front_test.article{Title:"Ferrets", Authors:[]string{"Tobi", "Loki", "Jane"}}
	// Some content here, so
	// interesting, you just
	// want to keep reading.
}

func ExampleMarshal() {
	a := article{Title: "Ferrets", Authors: []string{"Tobi", "Loki", "Jane"}}
	content := []byte(`Some content here, so
interesting, you just
want to keep reading.`)

	data, err := front.Marshal(a, content)
	if err != nil {
		log.Fatalf("error marshalling: %s", err)
	}

	fmt.Printf("%v\n", data)
	fmt.Printf("%s\n", data)
	// Output:
	// [45 45 45 10 116 105 116 108 101 58 32 70 101 114 114 101 116 115 10 97 117 116 104 111 114 115 58 10 45 32 84 111 98 105 10 45 32 76 111 107 105 10 45 32 74 97 110 101 10 45 45 45 10 83 111 109 101 32 99 111 110 116 101 110 116 32 104 101 114 101 44 32 115 111 10 105 110 116 101 114 101 115 116 105 110 103 44 32 121 111 117 32 106 117 115 116 10 119 97 110 116 32 116 111 32 107 101 101 112 32 114 101 97 100 105 110 103 46]
	// ---
	// title: Ferrets
	// authors:
	// - Tobi
	// - Loki
	// - Jane
	// ---
	// Some content here, so
	// interesting, you just
	// want to keep reading.
}
