package processor

import (
	m "ascii-web/pkg/models"
)

type Processor struct{}

func CreateNewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Get() *m.Error {
	return nil
}

func (p *Processor) Post(request m.Ascii) (m.Response, *m.Error) {
	response, err := p.ProcessTextIntoAscii(request)
	if err != nil {
		return m.Response{}, err
	}

	output := m.Response{
		Reponse: response,
	}

	return output, nil
}
