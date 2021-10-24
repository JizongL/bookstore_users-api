package user

import (
	"net/http"
	"strconv"

	"github.com/JizongL/bookstore_users-api/domain/users"
	"github.com/JizongL/bookstore_users-api/services"
	"github.com/JizongL/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){	
	var user users.User
	
	// bytes,err:= ioutil.ReadAll(c.Request.Body)
	// if err != nil{
	// 	// TODO Handle Error
	// 	fmt.Println(err)
	// }
	// if err:= json.Unmarshal(bytes,&user);err!=nil{
	// 	// TODO Handle json error
	// 	fmt.Println(err)
	// 	return
	// }
	// c.ShouldBindJSON does all the above commented out
	if err:= c.ShouldBindJSON(&user);err!=nil{
		restErr:= errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		// TODO: handle json error
		return
	}

	result,saveErr:=services.CreateUser(user)
	if saveErr !=nil{
		// TODO hanle user creation error
		c.JSON(saveErr.Status,saveErr)
		return
	}	
	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context){
	userId,userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr!=nil{
		err:= errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)
		return
	}
	result,getErr:= services.GetUser(userId)
	if getErr!=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK,result)
}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"implement me")
}
