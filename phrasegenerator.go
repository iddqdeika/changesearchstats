package main

import (
	"errors"
	"math/rand"
)

type LexemTemplate interface {
	GetValues() []string
}

type PhraseTemplate interface {
	GetLexemTamplates() []LexemTemplate
}

type PhraseGenerator struct{
	PhraseTemplates map[string]PhraseTemplate
}

func (p *PhraseGenerator) addTemplate(name string, template PhraseTemplate) error{
	if p.PhraseTemplates == nil{
		p.PhraseTemplates = make(map[string]PhraseTemplate)
	}
	if _, ok := p.PhraseTemplates[name]; ok{
		return errors.New("already exists")
	}
	p.PhraseTemplates[name] = template
	return nil
}

func (p *PhraseGenerator) generate(name string) (*string, error){
	if _, ok := p.PhraseTemplates[name]; !ok{
		return nil, errors.New("cannot find template")
	}

	template := p.PhraseTemplates[name]

	result := ""
	for _, lexem := range template.GetLexemTamplates(){
		if len(result)>0{
			result += " "
		}
		result += lexem.GetValues()[rand.Int63n(int64(len(lexem.GetValues())))]
	}

	return &result, nil
}

