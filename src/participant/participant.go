package participant

import(
	"fmt"
)
const(
	Ipa = 0
	Ips = 1
)

type Participant struct {
	Id               int32
	Ipa              float32
	Ips              float32
	ChosenDepartment []int32
	Status int
}

func (p Participant) GetScore(cluster int) float32 {
	var score float32
	switch cluster {
	case Ipa :
		score = p.Ipa
	case Ips :
		score = p.Ips
	}
	return score
}
func NewParticipant() Participant{
	return Participant{}
}

func (p Participant) AddParticipant(id int32, choosenDepartment []int32) {
	p.ChosenDepartment = choosenDepartment
	fmt.Println("AddDepartment ", id, choosenDepartment, "Success")
	p.Id = id
}


