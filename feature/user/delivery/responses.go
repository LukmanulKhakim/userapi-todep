package delivery

import "userapi/feature/user/domain"

func SuccessResponses(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponses(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type addUserRespons struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	HP     string `json:"hp"`
	Addres string `json:"addres"`
}

type editUserRespons struct {
	Name   string `json:"name"`
	HP     string `json:"hp"`
	Addres string `json:"addres"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "insert":
		cnv := core.(domain.Core)
		res = addUserRespons{ID: cnv.ID, Name: cnv.Name, HP: cnv.HP, Addres: cnv.Addres}
	case "edit":
		cnv := core.(domain.Core)
		res = editUserRespons{Name: cnv.Name, HP: cnv.HP, Addres: cnv.Addres}
	case "all":
		var arr []addUserRespons
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, addUserRespons{ID: val.ID, Name: val.Name, HP: val.HP, Addres: val.Addres})
		}
		res = arr
	}

	return res
}
