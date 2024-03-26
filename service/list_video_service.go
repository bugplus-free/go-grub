package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// ListVideoService 视频队列展示的服务
type ListVideoService struct {
}

// List 展示所有视频
func (service *ListVideoService) List() serializer.Response {
	videos:=[]model.Video{}
	err:=model.DB.Find(&videos).Error
	if err!=nil{
		return serializer.Response{
			Status: 500,
			Msg:    "服务器故障",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
