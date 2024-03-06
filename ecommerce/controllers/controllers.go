package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HashPassword (password string) string{

}

func VerifyPassword (usersPassword string, givenPassword string) (bool, string){

}

func Signup()  gin.HandlerFunc {

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments (ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return 
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error":"user already exist"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error":"this phone number is already in use"})
			return
		}
	}
}

func Login() gin.HandlerFunc {

}

func ProductViewerAdmin() gin.HandlerFunc{

}

func SearchProduct() gin.HandlerFunc{

}

func SearchProductByQuery() gin.HandlerFunc{

}