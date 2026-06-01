package flyweight

import (
	"strconv"
	"time"
	"fmt"
)

type Stock struct {
	Total int32
	Used  int32
}

func NewStock(total int32, used int32) *Stock {
	return &Stock{Total: total, Used: used}
}

type Activity struct {
	Id    int32
	Name  string
	Desc  string
	StartDate time.Time
	EndDate   time.Time
	Stock     *Stock
}

func (a *Activity) String() string {
	return fmt.Sprintf("Activity{Id: %d, Name: %s, Desc: %s, StartDate: %s, EndDate: %s, Stock: %d}", a.Id, a.Name, a.Desc, a.StartDate, a.EndDate, a.Stock.Used)
}

func NewActivity(id int32, name string, desc string, startDate time.Time, endDate time.Time, stock *Stock) *Activity {
	return &Activity{Id: id, Name: name, Desc: desc, StartDate: startDate, EndDate: endDate, Stock: stock}
}

var Activities =make(map[int32]*Activity)

func GetActivity(id int32) *Activity {
	if activity, ok := Activities[id]; ok {
		return activity
	}
	activity := NewActivity(id, "活动"+strconv.Itoa(int(id)), "活动描述", time.Now(), time.Now().Add(time.Hour*24), NewStock(100, 0))
	Activities[id] = activity
	return activity
}

func QueryActivity(id int32) *Activity {
	activity := GetActivity(id)
	stock:=NewStock(1000, GetStock())
	activity.Stock = stock
	return activity
}
