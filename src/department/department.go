package department
import (
	"../participant"
)

type Department struct {
	Id int32
	Quota int32
	Params interface{}
	AcceptedParticipant []*participant.Participant
}

func NewDepartment(id int32, quota int32, params interface{}) *Department {
	return &Department{Id: id, Quota: quota, Params: params, AcceptedParticipant:nil}
}



