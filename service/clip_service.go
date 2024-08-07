package service

import (
	"github.com/fondazionecrocereale/navegayaapi/data/request"
	"github.com/fondazionecrocereale/navegayaapi/data/response"
)

type ClipsService interface {
	Create(clips request.CreateClipsRequest)
	Update(clips request.UpdateClipsRequest)
	Delete(clipsId int)
	FindById(clipsId int) response.ClipsResponse
	FindAll() []response.ClipsResponse
}
