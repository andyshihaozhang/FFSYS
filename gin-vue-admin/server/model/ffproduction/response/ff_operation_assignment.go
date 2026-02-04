package response

// FfOperationAssignmentResponse 工序分配响应结构体
type FfOperationAssignmentResponse struct {
	ID                 uint    `json:"ID"`
	ProductId          uint    `json:"productId"`
	OperationId        uint    `json:"operationId"`
	EmployeeId         uint    `json:"employeeId"`
	EmployeeName       string  `json:"employeeName"`
	UnitPrice          float64 `json:"unitPrice"`
	AssignmentQuantity int     `json:"assignmentQuantity"`
	CreatedAt          int64   `json:"createdAt"`
	UpdatedAt          int64   `json:"updatedAt"`
}
