package document

import (
    "gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock_repository.go -package=mocks

type FileRepository interface {
    Save(file *File) error
    GetByID(id uint) (*File, error)
    GetAll() ([]File, error)
    DeleteFile(id uint) error 
}


type fileRepo struct {
    db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
    return &fileRepo{db: db}
}

func (r *fileRepo) Save(file *File) error {
    return r.db.Create(file).Error
}

func (r *fileRepo) DeleteFile(id uint) error {
    result := r.db.Delete(&File{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *fileRepo) GetByID(id uint) (*File, error) {
    var file File
    if err := r.db.First(&file, id).Error; err != nil {
        return nil, err
    }
    return &file, nil
}

func (r *fileRepo) GetAll() ([]File, error) {
    var files []File
    if err := r.db.Find(&files).Error; err != nil {
        return nil, err
    }
    return files, nil
}
