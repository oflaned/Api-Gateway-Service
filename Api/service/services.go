package service

type CompileProgram interface {
	RunProgram(program []byte) (out string)
	//CheckProgram(program structs.Program) (check bool, err error)
}

type Service struct {
	CompileProgram
}

func NewService() *Service {
	return &Service{
		CompileProgram: NewCompileService(),
	}
}
