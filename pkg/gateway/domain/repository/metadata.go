package repository

import (
	metadataclient "github.com/Bo0km4n/arc/pkg/metadata/client"
)

type MetadataRepository interface {
	Register()
}

type metadataRepository struct {
	metadataclient.Client
}

func (mr *metadataRepository) Register() {}

func NewMetadataRepository(client metadataclient.Client) MetadataRepository {
	return &metadataRepository{
		Client: client,
	}
}
