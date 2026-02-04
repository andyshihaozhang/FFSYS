package ffproduction

// 员工性别常量
const (
	SexMale   = 1 // 男
	SexFemale = 2 // 女
)

var SexMap = map[int]string{
	SexMale:   "男",
	SexFemale: "女",
}

var SexReverseMap = map[string]int{
	"男": SexMale,
	"女": SexFemale,
}

// 员工角色常量
const (
	AuthorityLeader   = 1 // 组长
	AuthorityEmployee = 2 // 员工
)

var AuthorityMap = map[int]string{
	AuthorityLeader:   "组长",
	AuthorityEmployee: "员工",
}

var AuthorityReverseMap = map[string]int{
	"组长": AuthorityLeader,
	"员工": AuthorityEmployee,
}

// 员工状态常量
const (
	StatusWorking = 1 // 在职
	StatusLeft    = 2 // 离职
	StatusLeave   = 3 // 休假
)

var StatusMap = map[int]string{
	StatusWorking: "在职",
	StatusLeft:    "离职",
	StatusLeave:   "休假",
}

var StatusReverseMap = map[string]int{
	"在职": StatusWorking,
	"离职": StatusLeft,
	"休假": StatusLeave,
}

// 产品标志常量
const (
	ProductFlagNotProduced = 1 // 未生产
	ProductFlagProducing   = 2 // 生产中
	ProductFlagCompleted   = 3 // 已完成
)

var ProductFlagMap = map[int]string{
	ProductFlagNotProduced: "未生产",
	ProductFlagProducing:   "生产中",
	ProductFlagCompleted:   "已完成",
}

var ProductFlagReverseMap = map[string]int{
	"未生产": ProductFlagNotProduced,
	"生产中": ProductFlagProducing,
	"已完成": ProductFlagCompleted,
}
