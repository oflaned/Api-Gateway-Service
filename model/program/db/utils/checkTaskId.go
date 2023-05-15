package utils

import "Mehmat/model/program"

const undefinedTask = "5190c60e-01a4-4bb6-9672-cf4d368d32bb"

func CheckTaskId(program *program.Program) {
	if program.TaskId == "" {
		program.TaskId = undefinedTask
	}
}
