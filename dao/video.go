package dao

import "time"

type Video struct {
	Id        int64 `json:"id"`
	AuthorId  int64
	VideoPath string `json:"video_path"`
	CoverPath string `json:"cover_path"`
	CreateAt  time.Time
	Title     string `json:"title"`
}

// TableName
//	将TableVideo映射到videos，
//	这样我结构体到名字就不需要是Video了，防止和我Service层到结构体名字冲突
func (Video) TableName() string {
	return "video"
}
