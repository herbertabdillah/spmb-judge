package main

import (
	"fmt"
)
const(
	Ipa = 0
	Ips = 1

)
func main(){
	fmt.Print("Halo")
	s := Selection{}
	s.ParticipantList = append(s.ParticipantList, &Participant{0, 90, 80, []int32{0,2}})
	s.ParticipantList = append(s.ParticipantList, &Participant{1, 85, 80, []int32{2,0,1}})
	s.ParticipantList = append(s.ParticipantList, &Participant{2, 80, 90, []int32{0,2}})
	s.ParticipantList = append(s.ParticipantList, &Participant{3, 100, 90, []int32{0,1}})
	s.DepartmentList = append(s.DepartmentList, &Department{0,2,0,nil})
	s.DepartmentList = append(s.DepartmentList, &Department{1,2,0,nil})
	s.DepartmentList = append(s.DepartmentList, &Department{2,3,1,nil})
	s.Execute()
	s.Print()
}
// Contract
type ParticipantContract interface {
	Add(id int32, choosenDepartment []int32)
}
type DepartmentContract interface {
	Add(id int32, quota int32, cluster int)
}
type SelectionContract interface {
	Execute()
	GetJson()
}

// Participant
type Participant struct {
	Id int32
	ipa float32
	ips float32
	ChosenDepartment []int32
	//status int
}

func (p Participant) GetScore(cluster int) float32 {
	var score float32
	switch cluster {
	case Ipa :
		score = p.ipa
	case Ips :
		score = p.ips
	}
	return score
}
func NewParticipant() Participant{
	return Participant{}
}

func (p Participant) Add(id int32, choosenDepartment []int32) {
	p.ChosenDepartment = choosenDepartment
	fmt.Println("Add ", id, choosenDepartment, "Success")
	p.Id = id
}

// Department
type Department struct {
	Id int32
	Quota int32
	Cluster int
	AcceptedParticipant []*Participant
}
func (d Department) Add(id int32, quota int32) {
	d.Id = id
	d.Quota = quota

	fmt.Println("Add ", id, quota, "Success")
}

type Selection struct{
	DepartmentList []*Department
	ParticipantList []*Participant
}
func (s Selection) InsertToDepProcess(p *Participant, d *Department, score float32){
	if d.AcceptedParticipant == nil {
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
			d.AcceptedParticipant = append(d.AcceptedParticipant[:i],append([]*Participant{p}, d.AcceptedParticipant[i:]...)...)
			break
		}
	}
}

func (s Selection) InsertToDep(p *Participant, d *Department) {
	score := p.GetScore(d.Cluster)
	isFull := len(d.AcceptedParticipant) >= int(d.Quota)
	//isScoreHigher := score > d.AcceptedParticipant[len(d.AcceptedParticipant)-1].GetScore(d.Cluster)
	if !isFull {
		s.InsertToDepProcess(p, d, score)
	} else if isFull && score > d.AcceptedParticipant[len(d.AcceptedParticipant)-1].GetScore(d.Cluster) {
		s.InsertToDepProcess(p, d, score)
		d.AcceptedParticipant = d.AcceptedParticipant[:len(d.AcceptedParticipant)-1]
	}
}
func (s Selection) Insert(p *Participant) {
	departmentIndex := p.ChosenDepartment[0]
	p.ChosenDepartment = p.ChosenDepartment[1:]
	department := s.DepartmentList[departmentIndex]
	s.InsertToDep(p, department)
}
func (s Selection) Execute(){
	isDone := false
	for !isDone {
		isDone = false
		for _, v := range s.ParticipantList {
			if len(v.ChosenDepartment) == 0 {
				isDone = true
				continue
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
