package xoss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type Factory struct {
	config *Config
}

func (f *Factory) Client() (*oss.Client, error) {
	client, err := oss.New(f.config.Endpoint, f.config.AccessKeyID, f.config.AccessKeySecret, oss.Timeout(30, 60))

	return client, err
}

func (f *Factory) Bucket() (*oss.Bucket, error) {
	client, err := f.Client()
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(f.config.Bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
