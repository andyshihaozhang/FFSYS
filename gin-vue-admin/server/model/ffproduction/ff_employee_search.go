package ffproduction

// FfEmployeeSearch 员工搜索结构体
type FfEmployeeSearch struct {
	EmployeeName   string `json:"employeeName" form:"employeeName"`
	Phone          string `json:"phone" form:"phone"`
	Sex            int    `json:"sex" form:"sex"`
	AuthorityIds   int    `json:"authorityIds" form:"authorityIds"`
	EmployeeStatus int    `json:"employeeStatus" form:"employeeStatus"`
	PageInfo
}
