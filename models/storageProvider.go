package models

// StorageProvider 是块存储提供程序应该如何实现的接口。
// 目前有本地和 S3 两种存储方式
type StorageProvider interface {
	Setup() error
	Save(filePath string, retryCount int) (string, error)

	SegmentWritten(localFilePath string)
	VariantPlaylistWritten(localFilePath string)
	MasterPlaylistWritten(localFilePath string)
}
