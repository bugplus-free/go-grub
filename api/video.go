package api

import (
	"go-crud/service"

	"github.com/gin-gonic/gin"
)

func CreateVideo(c *gin.Context){
	// var service service.CreateVideoService
	service :=service.CreateVideoService{}
	if err:=c.ShouldBind(&service);err==nil{
		res:=service.Create()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
	
}
func ShowVideo(c *gin.Context){
	// var service service.CreateVideoService
	service :=service.ShowVideoService{}
	if err:=c.ShouldBind(&service);err==nil{
		res:=service.Show(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
	
}

func ListVideo(c *gin.Context){
	// var service service.CreateVideoService
	service :=service.ListVideoService{}
	if err:=c.ShouldBind(&service);err==nil{
		res:=service.List()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
	
}

func UpdateVideo(c *gin.Context){
	// var service service.CreateVideoService
	service :=service.UpdateVideoService{}
	if err:=c.ShouldBind(&service);err==nil{
		res:=service.Update(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
	
}

func DeleteVideo(c *gin.Context){
	// var service service.CreateVideoService
	service :=service.DeleteVideoService{}
	if err:=c.ShouldBind(&service);err==nil{
		res:=service.Delete(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
	
}