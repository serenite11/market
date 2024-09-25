package failure

type LogicError struct {
	Desc     string `json:"description"`
	HttpCode int    `json:"code"`
	GrpcCode int
	inner    error
}

func (s *LogicError) Error() string {
	return s.Desc
}

func (s *LogicError) Wrap(err error) error {
	s.inner = err
	return s
}

func (s *LogicError) Unwrap() error {
	return s.inner
}

func (s *LogicError) Is(err error) bool {
	if _, ok := err.(*LogicError); ok {
		return true
	}
	return false
}
