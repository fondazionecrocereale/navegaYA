package service

import (
	"github.com/fondazionecrocereale/navegayaapi/data/request"
	"github.com/fondazionecrocereale/navegayaapi/data/response"
	"github.com/fondazionecrocereale/navegayaapi/helper"
	"github.com/fondazionecrocereale/navegayaapi/model"
	"github.com/fondazionecrocereale/navegayaapi/repository"

	"github.com/go-playground/validator/v10"
)

type ClipsServiceImpl struct {
	ClipsRepository repository.ClipsRepository
	Validate        *validator.Validate
}

func NewClipsServiceImpl(clipRepository repository.ClipsRepository, validate *validator.Validate) ClipsService {
	return &ClipsServiceImpl{
		ClipsRepository: clipRepository,
		Validate:        validate,
	}
}

// Create implements ClipsService
func (t *ClipsServiceImpl) Create(clips request.CreateClipsRequest) {
	err := t.Validate.Struct(clips)
	helper.ErrorPanic(err)
	clipModel := model.Clips{
		Name: clips.Name,
	}
	t.ClipsRepository.Save(clipModel)
}

// Delete implements ClipsService
func (t *ClipsServiceImpl) Delete(clipsId int) {
	t.ClipsRepository.Delete(clipsId)
}

// FindAll implements ClipsService
func (t *ClipsServiceImpl) FindAll() []response.ClipsResponse {
	result := t.ClipsRepository.FindAll()

	var clips []response.ClipsResponse
	for _, value := range result {
		clip := response.ClipsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		clips = append(clips, clip)
	}

	return clips
}

// FindById implements ClipsService
func (t *ClipsServiceImpl) FindById(clipsId int) response.ClipsResponse {
	clipData, err := t.ClipsRepository.FindById(clipsId)
	helper.ErrorPanic(err)

	clipResponse := response.ClipsResponse{
		Id:   clipData.Id,
		Name: clipData.Name,
	}
	return clipResponse
}

// Update implements ClipsService
func (t *ClipsServiceImpl) Update(clips request.UpdateClipsRequest) {
	clipData, err := t.ClipsRepository.FindById(clips.Id)
	helper.ErrorPanic(err)
	clipData.Name = clips.Name
	t.ClipsRepository.Update(clipData)
}
