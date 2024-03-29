package selection

import (

	"fmt"
)

const(
	OnProcess   int = -1
	NotAccepted int = - 2
)

type Selection struct{
	DepartmentList []*Department
	ParticipantList []*Participant
	GetScore func(participantParams *interface{}, args interface{}) float32
}
func (s *Selection) InsertToDepProcess(p *Participant, d *Department, score float32){
	if d.AcceptedParticipant == nil {
		d.AcceptedParticipant = append(d.AcceptedParticipant, p)
		return
	} else if score < d.AcceptedParticipant[len(d.AcceptedParticipant) - 1].GetScore(d.Params) {
		d.AcceptedParticipant = append(d.AcceptedParticipant, p)
	}
	for i := len(d.AcceptedParticipant) - 1; i >= 0; i-- {
		thisScore := d.AcceptedParticipant[i].GetScore(d.Params)
		if i > 0 {
			NextScore := d.AcceptedParticipant[i - 1].GetScore(d.Params)
			if thisScore == NextScore {
				continue
			} else if score > NextScore {
				continue
			}
		}
		if score > thisScore {
			d.AcceptedParticipant = append(d.AcceptedParticipant[:i],append([]*Participant{p}, d.AcceptedParticipant[i:]...)...)
			break
		}
	}
}

func (s *Selection) InsertToDep(p *Participant, d *Department) {
	score := p.GetScore(d.Params)
	isFull := len(d.AcceptedParticipant) >= int(d.Quota)
	//isScoreHigher := score > d.AcceptedParticipant[len(d.AcceptedParticipant)-1].GetScore(d.Cluster)
	if !isFull {
		s.InsertToDepProcess(p, d, score)
		p.Status = int(d.Id)
	} else if isFull && score > d.AcceptedParticipant[len(d.AcceptedParticipant)-1].GetScore(d.Params) {
		d.AcceptedParticipant[len(d.AcceptedParticipant) - 1].Status = OnProcess
		s.InsertToDepProcess(p, d, score)
		p.Status = int(d.Id)
		d.AcceptedParticipant = d.AcceptedParticipant[:len(d.AcceptedParticipant)-1]
	}
}
func (s *Selection) Insert(p *Participant) {
	departmentIndex := p.ChosenDepartment[0]
	p.ChosenDepartment = p.ChosenDepartment[1:]
	department := s.DepartmentList[departmentIndex]
	s.InsertToDep(p, department)
}

func (s *Selection) Print(){
	for _,v := range s.DepartmentList {
		fmt.Print(v.Id, " : ")
		for _, p := range v.AcceptedParticipant {
			fmt.Print(p.Id, ", ")
		}
		fmt.Println()
	}
}

