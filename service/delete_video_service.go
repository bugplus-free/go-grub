package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct {
}

// Delete 删除视频
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	video:=model.Video{}
	err:=model.DB.First(&video,id).Error
	if err!=nil{
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error: err.Error(),
		}
	}

	err=model.DB.Delete(&video).Error
	if err!=nil{
		return serializer.Response{
			Status: 500,
			Msg:    "视频删除失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{}
}
