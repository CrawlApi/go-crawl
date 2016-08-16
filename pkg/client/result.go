package client

type Result struct {
	Body string
	err  error
}

func (r *Result) HasError() bool {
	if r.err != nil {
		return true
	}
	return false
}

func (r *Result) GetError() error {
	return r.err
}

