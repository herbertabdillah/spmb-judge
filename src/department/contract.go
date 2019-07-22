package department

type DepartmentContract interface {
	AddDepartment(id int32, quota int32, cluster int)
}

