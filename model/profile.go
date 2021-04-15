package model

import "encoding/json"

type Profile struct {
	Name string // 姓名
	Gender string // 性别
	Age int // 年龄
	Height int //身高
	Weight int // 体重
	Income string // 收入
	Education string // 学历
	Marriage string // 婚况
	Occupation string // 职业
	Hokou string // 户口
	Xinzuo string // 星座
	House string // 房子
	Car string // 车子
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, nil
}