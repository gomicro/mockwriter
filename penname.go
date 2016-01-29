package penname

import ()

type PenName struct {
	Written     []byte
	returnError error
}

func New() *PenName {
	return &PenName{}
}

func (p *PenName) ReturnError(err error) {
	p.returnError = err
}
