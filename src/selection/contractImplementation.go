package selection

import (
	"../department"
	"../participant"
)
var SelectionPtr *Selection
func NewSelection() *Selection {
	SelectionPtr = &Selection{}
	return SelectionPtr
}

func (s Selection) SetScoringMethod(f func(participantParams interface{}, args interface{}) float32){
	s.GetScore = f
}

func (s Selection) AddDepartment(id int32, quota int32, params interface{}) {
	n := department.NewDepartment(id, quota, params)
	s.DepartmentList = append(s.DepartmentList, n)
}

func (s Selection) AddParticipant(id int32, param interface{}, chosenDepartment []int32) {
	n := participant.NewParticipant(id, param, chosenDepartment)
	s.ParticipantList = append(s.ParticipantList, n)
}
func (s Selection) Execute(){
	isDone := false
	for !isDone {
		isDone = true
		for _, v := range s.ParticipantList {
			if  v.Status >= 0{
				continue
			} else if len(v.ChosenDepartment) == 0 && v.Status == OnProcess {
				v.Status = NotAccepted
				continue
			} else {
				isDone = false
			}
			s.Insert(v)
		}
	}
}
func (s Selection) GetResult() {

}