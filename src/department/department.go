package department
import (
	"../participant"
	"fmt"
)
type DepartmentContract interface {
	Add(id int32, quota int32, cluster int)
}

// Department
type Department struct {
	Id int32
	Quota int32
	Cluster int
	AcceptedParticipant []*participant.Participant
}
func (d Department) Add(id int32, quota int32) {
	d.Id = id
	d.Quota = quota

	fmt.Println("Add ", id, quota, "Success")
}


