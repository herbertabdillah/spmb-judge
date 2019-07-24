package selection

const(
	Ipa = 0
	Ips = 1
)

type Participant struct {
	Id               int32
	Param            interface{}
	ChosenDepartment []int32
	Status           int
}

func NewParticipant(id int32, param interface{}, chosenDepartment []int32) *Participant {
	return &Participant{Id: id, Param: param, ChosenDepartment: chosenDepartment, Status:OnProcess}
}


func (p *Participant) GetScore(args interface{}) float32 {
	return SelectionPtr.GetScore(&p.Param, args)
}

