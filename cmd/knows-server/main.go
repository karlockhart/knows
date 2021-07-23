package main

import (
	"fmt"

	"github.com/karlockhart/knows"
)

func main() {
	s, err := knows.NewServer()
	if err != nil {
		panic(err)
	}

	u, _ := s.Create(*knows.NewKnow("title 1", []string{"simple", "first", "karl"}, []byte("a whole long string")))
	s.Create(*knows.NewKnow("title 2", []string{"simple", "second", "karl"}, []byte("a whole long string")))
	s.Create(*knows.NewKnow("title 3", []string{"simple", "third", "karl"}, []byte("a whole long string")))

	kn := s.Read(u)
	fmt.Println(kn.String(), "\n")

	fmt.Println(s.Dump(), "\n")

	kn.Title = "New Title"
	kn.Tags = append(kn.Tags, "all")
	s.Update(u, *kn)

	kns := s.FindByTag("all")
	for _, k := range kns {
		fmt.Println(k.String())
	}

	kn = s.Read(u)
	fmt.Println(kn.String(), "\n")

	fmt.Println(s.Dump(), "\n")

	s.Delete(u)
	fmt.Println(s.Dump(), "\n")

	kns = s.FindByTag("karl")
	for _, k := range kns {
		fmt.Println(k.String())
	}

}
