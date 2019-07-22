package participant

type ParticipantContract interface {
	AddParticipant(id int32, choosenDepartment []int32)
}
