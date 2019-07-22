package main
import(
	"fmt"
	"github.com/go-delve/delve/service/api"
)
const(
	Agm = 1
	Ipa = 2
	Ips = 3
	Process = -1
	NotAccept = -2
)
func main(){
	fmt.Print("Halo")
	var ParticipantList []Participant
	var DepartmentList []Department
	ti := Department{Id:0, Quota:2}
	si := Department{Id:1, Quota:3}
	mtk := Department{Id:2, Quota:4}
	a := Participant{Id:0, ChosenDepartment:0}
	var i ParticipantContract = p
	i.Add(1, []int32{1, 2, 3})

	ParticipantList = append(ParticipantList, i.(Participant))

}
// Contract
type ParticipantContract interface {
	Add(id int32, choosenDepartment []int32)
	GetScore() float32
}
type DepartmentContract interface {
	Add(id int32, quota int32)
}
type SelectionContract interface {
	Execute()
	GetJson()
}

// Participant
type Participant struct {
	Id int32
	tpa float32
	ipa float32
	ips float32
	agm float32
	ChosenDepartment []int32
	status int
}

func (p Participant) GetScore(cluster int) float32 {
	var score float32
	switch cluster {
	case Agm :
		score = (p.tpa + p.agm) / 2
	case Ipa :
		score = (p.tpa + p.ipa) / 2
	case Ips :
		score = (p.tpa + p.ips) / 2
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
}
func (d Department) Add(id int32, quota int32) {
	d.Id = id
	d.Quota = quota

	fmt.Println("Add ", id, quota, "Success")
}
// Selection
type Selection struct {
	DepartmentParticipant []node
	NotAccept []Participant
	ProcessQueue []Participant
}
type node struct {
	department *Department
	queue []*Participant
}
func (s Selection) Execute() {

}
func (s Selection) GetJson() {

}
func (s Selection) Print() {
	for _, v := range s.DepartmentParticipant {
		fmt.Println("")
	}
}
func (s Selection) Init(DepartmentList []Department, ParticipantList []Participant) {
	s.DepartmentParticipant = make([]node, len(DepartmentList))
	for _, participant := range ParticipantList{
		s.ProcessQueue = append(s.ProcessQueue, participant)
		participant.status = Process
	}
}
func (s Selection) Insert(queue []*Participant, score float32) {
	for i:= len(queue) - 1; i >=0; i-- {
		if(score > queue[i].GetScore(1)) {
			continue
		} else {
			queue = append(queue[:i+1], queue[i+1:len(queue)]...)
		}
	}
}
func (s Selection) Asdf() {
	i := 0
	for len(s.ProcessQueue) > 0 {
		participant := s.ProcessQueue[i]
		if participant.status != Process {
			continue
		}

		departmentIndex := participant.ChosenDepartment[0]
		departmentParticipant := s.DepartmentParticipant[departmentIndex]
		queue := departmentParticipant.queue
		isFull := len(departmentParticipant.queue) >= int(departmentParticipant.department.Quota)
		isScorePass := participant.GetScore(1) > queue[len(queue) - 1].GetScore(1)

		if !isFull {
			queue = append(queue, &participant)
			participant.ChosenDepartment = participant.ChosenDepartment[1:]
			participant.status = int(departmentIndex)
		} else if isFull && isScorePass {

		} else {

		}
		i++
		if  i >= len(s.ProcessQueue) {
			i = 0
		}
	}
}