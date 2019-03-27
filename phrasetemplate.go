package main

func NewTemplate(phrases [][]string) PhraseTemplate{
	t := &phraseTemplate{}
	for _, wordSet := range phrases{
		t.lexems = append(t.lexems,&lexemTemplate{words:wordSet})
	}
	return t
}

type phraseTemplate struct {
	lexems []LexemTemplate
}

func (p *phraseTemplate) GetLexemTamplates() []LexemTemplate {
	return p.lexems
}

type lexemTemplate struct {
	words []string
}

func (l *lexemTemplate) GetValues() []string{
	return l.words
}
