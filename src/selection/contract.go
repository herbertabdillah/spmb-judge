package selection
type SpmbJudgeContract interface{
	SetScoringMethod(f func(participantParams interface{}, args interface{}) float32)
	AddDepartment(id int32, quota int32, params interface{})
	AddParticipant(id int32, param interface{}, chosenDepartment []int32)
	Execute()
	GetResult()
}
