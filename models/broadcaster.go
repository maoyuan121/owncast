package models

import "time"

// Broadcaster 表示有关入站广播连接的详细信息。
type Broadcaster struct {
	RemoteAddr    string               `json:"remoteAddr"`    // ip 地址
	StreamDetails InboundStreamDetails `json:"streamDetails"` // 入站流的详情
	Time          time.Time            `json:"time"`          // 时间
}

// 入站流详情
type InboundStreamDetails struct {
	Width          int     `json:"width"`
	Height         int     `json:"height"`
	VideoFramerate float32 `json:"framerate"`
	VideoBitrate   int     `json:"videoBitrate"`
	VideoCodec     string  `json:"videoCodec"`
	AudioBitrate   int     `json:"audioBitrate"`
	AudioCodec     string  `json:"audioCodec"`
	Encoder        string  `json:"encoder"`
	VideoOnly      bool    `json:"-"`
}

// RTMPStreamMetadata is the raw metadata that comes in with a RTMP connection.
type RTMPStreamMetadata struct {
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	VideoBitrate   float32     `json:"videodatarate"`
	VideoCodec     interface{} `json:"videocodecid"`
	VideoFramerate float32     `json:"framerate"`
	AudioBitrate   float32     `json:"audiodatarate"`
	AudioCodec     interface{} `json:"audiocodecid"`
	Encoder        string      `json:"encoder"`
}
