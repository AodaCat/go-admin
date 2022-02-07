package example

import "go-admin/global"

// file struct, 文件结构体
type ExaFile struct {
	global.GA_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.GA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
