package model

type ShopMemberFollowHistory struct {
	Id uint32
	MemberId int32
	Category int8
	Type int8
	ContactPurpose int8
	ContactStatus int8
	ContactResult int8
	OrderType int8
	OrderStatus int8
	SalesId uint32
	VisitTime string
	ReceiveTime string
	LeaveTime string
	Duration int32
	ShopId int32
	CabinetId int32
	CabinetStatus int8
	ParentId int32
	CourseCatId int32
	OrderId int32
	ContractId int32
	Unit int32
	FailReason int8
	NextContact string
	Content string
	AssignContent string
	Text string
	Uid int32
	StaffId int32
	UpdatedTime string
	CreatedTime string
	ReceiveSalesId int32
	SpaceId int32
}
