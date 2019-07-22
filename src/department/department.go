package department
import (
	"../participant"
	"fmt"
)

type Department struct {
	Id int32
	Quota int32
	Cluster int
	AcceptedParticipant []*participant.Participant
}
func (d Department) AddDepartment(id int32, quota int32) {
	d.Id = id
	d.Quota = quota

	fmt.Println("AddDepartment ", id, quota, "Success")
}


