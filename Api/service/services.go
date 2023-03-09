package service

import "Mehmat/structs"

type CompileProgram interface {
	RunProgram(program structs.Program) (out string, err error)
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
