package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liuyibao/migration/model"
	"fmt"
	"bytes"
)

var DbGymSecond *sql.DB
var DbGym *sql.DB

const SPAN  = 30

func init() {
	var err error
	DbGymSecond, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ran_gym_second")
	if err != nil {
		panic(err)
	}

	DbGym, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ran_gym")
	if err != nil {
		panic(err)
	}
}

func main()  {
	var maxId, offset uint32
	err := DbGym.QueryRow("SELECT MAX(id) FROM shop_member_follow_history").Scan(&maxId);
	if err != nil {
		panic(err)
	}

	for offset=0; offset<maxId; offset+=SPAN {
		modelList := readData(offset, SPAN)
		groups := groupData(modelList)
		for index,value := range groups {
			writeData(index, value)
		}
	}
}

func readData(offset uint32, span uint32) (modelList []model.ShopMemberFollowHistory) {
	rows, err := DbGym.Query("SELECT * FROM shop_member_follow_history WHERE id>? AND id<=?", offset, offset+span)
	if err != nil {
		panic(err)
	}


	for rows.Next() {
		m := model.ShopMemberFollowHistory{}
		err = rows.Scan(
			&m.Id,
			&m.MemberId,
			&m.Category,
			&m.Type,
			&m.ContactPurpose,
			&m.ContactStatus,
			&m.ContactResult,
			&m.OrderType,
			&m.OrderStatus,
			&m.SalesId,
			&m.VisitTime,
			&m.ReceiveTime,
			&m.LeaveTime,
			&m.Duration,
			&m.ShopId,
			&m.CabinetId,
			&m.CabinetStatus,
			&m.ParentId,
			&m.CourseCatId,
			&m.OrderId,
			&m.ContractId,
			&m.Unit,
			&m.FailReason,
			&m.NextContact,
			&m.Content,
			&m.AssignContent,
			&m.Text,
			&m.Uid,
			&m.StaffId,
			&m.UpdatedTime,
			&m.CreatedTime,
			&m.ReceiveSalesId,
			&m.SpaceId,
		)
		if err != nil {
			panic(err)
		}

		modelList = append(modelList, m)
	}

	return
}

func groupData(modelList []model.ShopMemberFollowHistory) (map[string][]model.ShopMemberFollowHistory) {
	groups := make(map[string][]model.ShopMemberFollowHistory)
	for _, m := range modelList {
		mod := fmt.Sprintf("%02d", m.Uid%100)
		groups[mod] = append(groups[mod], m)
	}
	return groups
}

func writeData(suffix string, modelList []model.ShopMemberFollowHistory) {
	insertSql := "INSERT INTO shop_member_follow_history_" + suffix + "(id,member_id,category,type,contact_purpose,contact_status,contact_result,order_type,order_status,sales_id,visit_time,receive_time,leave_time,duration,shop_id,cabinet_id,cabinet_status,parent_id,course_cat_id,order_id,contract_id,unit,fail_reason,next_contact,content,assign_content,text,uid,staff_id,updated_time,created_time,receive_sales_id,space_id) VALUES"
	var buf bytes.Buffer
	buf.WriteString(insertSql)
	buf.WriteString("(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	for i:=1;i<len(modelList);i++ {
		buf.WriteString(",(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	}

	stmt,err := DbGymSecond.Prepare(buf.String())
	if err != nil {
		panic(err)
	}

	var args []interface{}
	for _,m := range modelList {
		args = append(args,
			m.Id,
			m.MemberId,
			m.Category,
			m.Type,
			m.ContactPurpose,
			m.ContactStatus,
			m.ContactResult,
			m.OrderType,
			m.OrderStatus,
			m.SalesId,
			m.VisitTime,
			m.ReceiveTime,
			m.LeaveTime,
			m.Duration,
			m.ShopId,
			m.CabinetId,
			m.CabinetStatus,
			m.ParentId,
			m.CourseCatId,
			m.OrderId,
			m.ContractId,
			m.Unit,
			m.FailReason,
			m.NextContact,
			m.Content,
			m.AssignContent,
			m.Text,
			m.Uid,
			m.StaffId,
			m.UpdatedTime,
			m.CreatedTime,
			m.ReceiveSalesId,
			m.SpaceId,
		)
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("插入数量", rowCnt)
}
