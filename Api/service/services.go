package service

import "Mehmat/model/program"

type CompileProgram interface {
	RunProgram(program []byte) (out string)
}

type Service struct {
	program.Repository
	CompileProgram
}

func NewService(rep program.Repository) *Service {
	return &Service{
		CompileProgram: NewCompileService(),
		Repository:     rep,
	}
}
