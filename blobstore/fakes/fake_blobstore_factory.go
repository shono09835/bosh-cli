package fakes

import (
	biblobstore "github.com/shono09835/bosh-cli/v7/blobstore"
)

type FakeBlobstoreFactory struct {
	CreateBlobstoreURL string
	CreateBlobstore    biblobstore.Blobstore
	CreateErr          error
}

func NewFakeBlobstoreFactory() *FakeBlobstoreFactory {
	return &FakeBlobstoreFactory{}
}

func (f *FakeBlobstoreFactory) Create(blobstoreURL string) (biblobstore.Blobstore, error) {
	f.CreateBlobstoreURL = blobstoreURL
	return f.CreateBlobstore, f.CreateErr
}
