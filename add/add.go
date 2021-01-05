package add

type Request struct {
	A int
	B int
}

func (r *Request) Add() (int, error) {
	return r.A + r.B, nil
}
