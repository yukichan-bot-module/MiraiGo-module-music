package service

import (
	"encoding/json"
	"fmt"

	"github.com/Mrs4s/MiraiGo/message"
	"github.com/yukichan-bot-module/MiraiGo-module-music/internal/pkg"
)

const cloudAPI = "https://music.163.com/api/search/get/web"

// CloudAPIResponse 网易云 API 返回内容
type CloudAPIResponse struct {
	Result struct {
		Songs []struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				ID        int           `json:"id"`
				Name      string        `json:"name"`
				PicURL    interface{}   `json:"picUrl"`
				Alias     []interface{} `json:"alias"`
				AlbumSize int           `json:"albumSize"`
				PicID     int           `json:"picId"`
				Img1V1URL string        `json:"img1v1Url"`
				Img1V1    int           `json:"img1v1"`
				Trans     interface{}   `json:"trans"`
			} `json:"artists"`
			Album struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Artist struct {
					ID        int           `json:"id"`
					Name      string        `json:"name"`
					PicURL    interface{}   `json:"picUrl"`
					Alias     []interface{} `json:"alias"`
					AlbumSize int           `json:"albumSize"`
					PicID     int           `json:"picId"`
					Img1V1URL string        `json:"img1v1Url"`
					Img1V1    int           `json:"img1v1"`
					Trans     interface{}   `json:"trans"`
				} `json:"artist"`
				PublishTime int64 `json:"publishTime"`
				Size        int   `json:"size"`
				CopyrightID int   `json:"copyrightId"`
				Status      int   `json:"status"`
				PicID       int64 `json:"picId"`
				Mark        int   `json:"mark"`
			} `json:"album"`
			Duration    int         `json:"duration"`
			CopyrightID int         `json:"copyrightId"`
			Status      int         `json:"status"`
			Alias       []string    `json:"alias"`
			Rtype       int         `json:"rtype"`
			Ftype       int         `json:"ftype"`
			Mvid        int         `json:"mvid"`
			Fee         int         `json:"fee"`
			RURL        interface{} `json:"rUrl"`
			Mark        int         `json:"mark"`
		} `json:"songs"`
		SongCount int `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}

// SearchCloudMusic 搜索网易云
func SearchCloudMusic(songName string) (*message.MusicShareElement, string) {
	respBody, err := pkg.GetRequest(cloudAPI, [][]string{{"s", songName}, {"type", "1"}})
	if err != nil {
		return nil, "网易云 API 访问超时。"
	}
	var apiResp CloudAPIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, "网易云 API 错误"
	}
	songCount := apiResp.Result.SongCount
	if songCount == 0 {
		return nil, "未找到歌曲" + songName
	}
	song := apiResp.Result.Songs[0]
	songName = song.Name
	songArtist := "U.A."
	if len(song.Artists) != 0 {
		songArtist = song.Artists[0].Name
	}
	songURL := fmt.Sprintf("https://y.music.163.com/m/song/%d/", song.ID)
	return &message.MusicShareElement{
		MusicType: message.CloudMusic,
		Title:     songName,
		Summary:   songArtist,
		Url:       songURL,
	}, ""
}
