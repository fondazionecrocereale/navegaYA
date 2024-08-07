package repository

import (
	"errors"

	"github.com/fondazionecrocereale/navegayaapi/data/request"
	"github.com/fondazionecrocereale/navegayaapi/helper"
	"github.com/fondazionecrocereale/navegayaapi/model"

	"gorm.io/gorm"
)

type ClipsRepositoryImpl struct {
	Db *gorm.DB
}

func NewClipsREpositoryImpl(Db *gorm.DB) ClipsRepository {
	return &ClipsRepositoryImpl{Db: Db}
}

// Delete implements ClipsRepository
func (t *ClipsRepositoryImpl) Delete(clipsId int) {
	var clips model.Clips
	result := t.Db.Where("id = ?", clipsId).Delete(&clips)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ClipsRepository
func (t *ClipsRepositoryImpl) FindAll() []model.Clips {
	var clips []model.Clips
	result := t.Db.Find(&clips)
	helper.ErrorPanic(result.Error)
	return clips
}

// FindById implements ClipsRepository
func (t *ClipsRepositoryImpl) FindById(clipsId int) (clips model.Clips, err error) {
	var tag model.Clips
	result := t.Db.Find(&tag, clipsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements ClipsRepository
func (t *ClipsRepositoryImpl) Save(clips model.Clips) {
	result := t.Db.Create(&clips)
	helper.ErrorPanic(result.Error)
}

// Update implements ClipsRepository
func (t *ClipsRepositoryImpl) Update(clips model.Clips) {
	var updateTag = request.UpdateClipsRequest{
		Id:   clips.Id,
		Name: clips.Name,
	}
	result := t.Db.Model(&clips).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
