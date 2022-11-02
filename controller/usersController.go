package controller

import (
	"net/http"

	"github.com/daniel00k/medicalhistory/model"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}
func (uc UsersController) Index(ctx *gin.Context) {
	u := model.User{}
	ctx.JSON(http.StatusOK, u.All())
}

func (uc UsersController) Create(ctx *gin.Context) {
	var u model.User
	if err := ctx.BindJSON(&u); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, model.User.Create(u))
}

// func Show(router *gin.Engine, client *mongo.Client) {
// 	router.GET("/users", func(ctx *gin.Context) {
// 		u := model.User{
// 			Fname: "Juan",
// 			Lname: "ga",
// 			Email: "somemail",
// 		}
// 		collection := client.Database("historialmedico").Collection("users")
// 		ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer cancel()
// 		res, err := collection.InsertOne(ctx2, u)
// 		if err != nil {
// 			fmt.Printf("asahdjkasdhjkashdkjsa" + err.Error())
// 		}
// 		id := res.InsertedID
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": id,
// 		})
// 	})
// }
