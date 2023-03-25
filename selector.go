package bprompt

import (
	//"fmt"
	//"bufio"
	//"fmt"
	"io"
	//"os"

	"github.com/fatih/color"
	//"github.com/mattn/go-tty"
)

type Selector struct{
	Writer io.Writer
	Prompt string
	PromptStyle []color.Attribute
	List []interface{}
	ListStyle []color.Attribute
	ActiveStyle []color.Attribute
	Selected interface{}
}

func (s *Selector)Len()(int){
	return len(s.List)
}

func (s *Selector)showPrompt(){
	color := color.New(s.PromptStyle...)
	color.Fprintln(s.Writer, s.Prompt)
}

func (s *Selector)showMenu(){

}


func (s *Selector)Ask()(*Selector, error){
	s.showPrompt()
	s.showMenu()
	return &Selector{}, nil
}
