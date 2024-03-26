package serializer
import "go-crud/model"
// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title  string `json:"title"`
	Info  string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}


func BuildVideo(video model.Video) Video {
	return Video{
		ID:        video.ID,
		Title:  video.Title,
		Info:  video.Info,
		CreatedAt: video.CreatedAt.Unix(),
	}
}
func BuildVideos(videos []model.Video) (terms []Video) {
	for _ ,term :=range videos{
		terms=append(terms,BuildVideo(term))
	}
	return terms
}