package processor

import (
	m "ascii-web/pkg/models"
)

// the processor struct
type Processor struct{}

// The contructor to create a new processor
func CreateNewProcessor() *Processor {
	return &Processor{}
}

// The implementation of the Post method from the ascii art interface 
func (p *Processor) Post(request m.Ascii) (string, *m.Error) {
	response, err := p.ProcessTextIntoAscii(request)
	if err != nil {
		return "", err
	}

	return response, nil
}
