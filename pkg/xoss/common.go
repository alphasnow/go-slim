package xoss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func NewClient(config *Config) (*oss.Client, error) {
	return oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret, oss.Timeout(30, 60))
}

func NewBucket(client *oss.Client, config *Config) (*oss.Bucket, error) {
	return client.Bucket(config.Bucket)
}
