package selection

import (
	"../department"
	"../participant"
	"fmt"
)

const(
	OnProcess   int = -1
	NotAccepted int = - 2
)

type Selection struct{
	DepartmentList []*department.Department
	ParticipantList []*participant.Participant
}
func (s Selection) InsertToDepProcess(p *participant.Participant, d *department.Department, score float32){
	if d.AcceptedParticipant == nil {
		d.AcceptedParticipant = append(d.AcceptedParticipant, p)
		return
	} else if score < d.AcceptedParticipant[len(d.AcceptedParticipant) - 1].GetScore(d.Cluster) {
		d.AcceptedParticipant = append(d.AcceptedParticipant, p)
	}
	for i := len(d.AcceptedParticipant) - 1; i >= 0; i-- {
		thisScore := d.AcceptedParticipant[i].GetScore(d.Cluster)
		if i > 0 {
			NextScore := d.AcceptedParticipant[i - 1].GetScore(d.Cluster)
			if thisScore == NextScore {
				continue
			} else if score > NextScore {
				continue
			}
		}
		if score > thisScore {
			d.AcceptedParticipant = append(d.AcceptedParticipant[:i],append([]*participant.Participant{p}, d.AcceptedParticipant[i:]...)...)
			break
		}
	}
}

func (s Selection) InsertToDep(p *participant.Participant, d *department.Department) {
	score := p.GetScore(d.Cluster)
	isFull := len(d.AcceptedParticipant) >= int(d.Quota)
	//isScoreHigher := score > d.AcceptedParticipant[len(d.AcceptedParticipant)-1].GetScore(d.Cluster)
	if !isFull {
		s.InsertToDepProcess(p, d, score)
		p.Status = int(d.Id)
	} else if isFull && score > d.AcceptedParticipant[len(d.AcceptedParticipant)-1].GetScore(d.Cluster) {
		d.AcceptedParticipant[len(d.AcceptedParticipant) - 1].Status = OnProcess
		s.InsertToDepProcess(p, d, score)
		p.Status = int(d.Id)
		d.AcceptedParticipant = d.AcceptedParticipant[:len(d.AcceptedParticipant)-1]
	}
}
func (s Selection) Insert(p *participant.Participant) {
	departmentIndex := p.ChosenDepartment[0]
	p.ChosenDepartment = p.ChosenDepartment[1:]
	department := s.DepartmentList[departmentIndex]
	s.InsertToDep(p, department)
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

func (s Selection) Print(){
	for _,v := range s.DepartmentList {
		fmt.Print(v.Id, " : ")
		for _, p := range v.AcceptedParticipant {
			fmt.Print(p.Id, ", ")
		}
		fmt.Println()
	}
}

