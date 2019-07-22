package main

import (
	"./department"
	"./participant"
	"./selection"
	"fmt"
)

func main(){
	fmt.Print("Halo")
	s := selection.Selection{}
	s.ParticipantList = append(s.ParticipantList, &participant.Participant{0, 90, 80, []int32{0,2}, selection.OnProcess})
	s.ParticipantList = append(s.ParticipantList, &participant.Participant{1, 85, 80, []int32{2,0,1}, selection.OnProcess})
	s.ParticipantList = append(s.ParticipantList, &participant.Participant{2, 80, 90, []int32{0,1}, selection.OnProcess})
	s.ParticipantList = append(s.ParticipantList, &participant.Participant{3, 100, 90, []int32{0,1}, selection.OnProcess})
	s.DepartmentList = append(s.DepartmentList, &department.Department{0,2,0,nil})
	s.DepartmentList = append(s.DepartmentList, &department.Department{1,2,0,nil})
	s.DepartmentList = append(s.DepartmentList, &department.Department{2,3,1,nil})
	s.Execute()
	s.Print()
}
// Contract



