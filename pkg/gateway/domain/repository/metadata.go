package repository

type MetadataRepository interface {
	Register()
}

type metadataRepository struct {
}

func (mr *metadataRepository) Register() {}

func NewMetadataRepository() MetadataRepository {
	return &metadataRepository{}
}
