package annotation

type MethodOptions interface {
	GetSkipAuth() bool
	GetReqMaskPaths() []string
	GetRespMaskPaths() []string
}

var emptyMethodOptions = &methodOptionsImpl{}

type methodOptionsImpl struct {
}

var _ MethodOptions = (*methodOptionsImpl)(nil)

func (methodOptionsImpl) GetSkipAuth() bool          { return false }
func (methodOptionsImpl) GetReqMaskPaths() []string  { return nil }
func (methodOptionsImpl) GetRespMaskPaths() []string { return nil }
