package bprompt

type Selector struct{
	Prompt string
	List []interface{}
}

func (s *Selector)Len()(int){
	return len(s.List)
}

func (s *Selector)Ask()(*Selector, error){
	return &Selector{}, nil
}
