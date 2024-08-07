package repository

import "github.com/fondazionecrocereale/navegayaapi/model"

type ClipsRepository interface {
	Save(clips model.Clips)
	Update(clips model.Clips)
	Delete(clipsId int)
	FindById(clipsId int) (tags model.Clips, err error)
	FindAll() []model.Clips
}
