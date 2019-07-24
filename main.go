package main

import (
	"github.com/herbertabdillah/selection"
	"fmt"
)

const(
	Ipa = 0
	Ips = 1
)

func main(){
	fmt.Print("Halo")
	//s := selection.Selection{}
	//s.ParticipantList = append(s.ParticipantList, &participant.Participant{0, 90, 80, []int32{0,2}, selection.OnProcess})
	//s.ParticipantList = append(s.ParticipantList, &participant.Participant{1, 85, 80, []int32{2,0,1}, selection.OnProcess})
	//s.ParticipantList = append(s.ParticipantList, &participant.Participant{2, 80, 90, []int32{0,1}, selection.OnProcess})
	//s.ParticipantList = append(s.ParticipantList, &participant.Participant{3, 100, 90, []int32{0,1}, selection.OnProcess})
	//s.DepartmentList = append(s.DepartmentList, &department.Department{0,2,0,nil})
	//s.DepartmentList = append(s.DepartmentList, &department.Department{1,2,0,nil})
	//s.DepartmentList = append(s.DepartmentList, &department.Department{2,3,1,nil})
	//s.Execute()
	//s.Print()
	f := func(participantParams *interface{}, args interface{}) float32 {
		var score float32
		param := (*participantParams).(map[string]interface{})
		switch args.(int) {
		case Ipa:
			score = param["ipa"].(float32)
		case Ips:
			score = param["ips"].(float32)
		}
		return score
	}
	var s selection.SpmbJudgeContract = selection.NewSelection()
	s.SetScoringMethod(f)
	s.AddDepartment(0,2, Ipa)
	s.AddDepartment(1, 2, Ipa)
	s.AddDepartment(2, 3, Ips)
	s.AddParticipant(0, map[string]interface{}{"ipa" : float32(90), "ips" : float32(80)}, []int32{0,2})
	s.AddParticipant(1, map[string]interface{}{"ipa" : float32(85), "ips" : float32(80)}, []int32{2,0,1})
	s.AddParticipant(2, map[string]interface{}{"ipa" : float32(80), "ips" : float32(90)}, []int32{0,1})
	s.AddParticipant(3, map[string]interface{}{"ipa" : float32(100), "ips" : float32(90)}, []int32{0,1})
	s.Execute()
	s.GetResult()
}
// Contract



