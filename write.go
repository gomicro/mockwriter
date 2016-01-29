package penname

func (p *PenName) Write(b []byte) (n int, err error) {
	if p.returnError != nil {
		return 0, p.returnError
	}

	p.Written = b
	return len(p.Written), nil
}
