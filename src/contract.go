package main
import (
	"./department"
	"./participant"
	"./selection"
)
type SpmbJudgeContract interface{
	department.DepartmentContract
	participant.ParticipantContract
	selection.SelectionContract
}
