package main

import (
	"github.com/Hayao0819/bprompt"
)


func main(){
	s:=bprompt.Select("Hoge??", "A", "B", "C")
	s.Ask()
}

