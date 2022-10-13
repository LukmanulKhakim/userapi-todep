package delivery

import "userapi/feature/user/domain"

type addUserFormat struct {
	Name   string `json:"name" form:"name"`
	HP     string `json:"hp" form:"hp"`
	Addres string `json:"addres"  form:"addres"`
}

type editUserFormat struct {
	Name   string `json:"name"  form:"name"`
	HP     string `json:"hp" form:"hp"`
	Addres string `json:"addres"  form:"addres"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case addUserFormat:
		cnv := i.(addUserFormat)
		return domain.Core{Name: cnv.Name, HP: cnv.HP, Addres: cnv.Addres}
	case editUserFormat:
		cnv := i.(editUserFormat)
		return domain.Core{Name: cnv.Name, HP: cnv.HP, Addres: cnv.Addres}
	}
	return domain.Core{}
}
