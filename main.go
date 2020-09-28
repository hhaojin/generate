package main

import "github.com/hhaojin/generate/Commands"

func main() {
	//user := models.NewUsersModel()
	//porm.Model(user)
	//{
	//	porm.DebugMode = true
	//	users := make([]models.UsersModel, 0)
	//	err := porm.Build(`select user_id,name,password from users where name=?`).
	//		Args("user5").
	//		First(user)
	//
	//	//err := user.BuildX(`select user_id,name,{{set_password}} from {{.Table}} where name=?`,
	//	//	porm.F{
	//	//		"set_password" : func() string{
	//	//			return "password + 101 as password"
	//	//		},
	//	//	}, "user2")
	//	if err != nil {
	//		log.Println("ERR---", err)
	//	}
	//	fmt.Println(user)
	//	fmt.Println(users)
	//}
	Commands.Parse()
}
