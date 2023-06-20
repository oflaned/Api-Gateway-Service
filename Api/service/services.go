package service

import (
	"Mehmat/model/mossStatus"
	"Mehmat/model/program"
)

type CompileProgram interface {
	RunProgram(program []byte) (out string)
}

type Service struct {
	ProgramRep program.Repository
	StatusRep  mossStatus.Repository
	CompileProgram
}

func NewService(repProgram program.Repository, repStatus mossStatus.Repository) *Service {
	return &Service{
		CompileProgram: NewCompileService(),
		ProgramRep:     repProgram,
		StatusRep:      repStatus,
	}
}
