package bost

import (
	"errors"
	"fmt"

	"github.com/mozgunovdm/omp-bot/internal/model/mdi"
)

var ErrDataIDRange = errors.New("id out of range data")
var ErrEmptyModel = errors.New("Список продуктов пуст")
var ErrPosOutRange = errors.New("Конец списка")

type BostService interface {
	Describe(bostID uint64) (*mdi.Bost, error)
	List(cursor uint64, limit uint64) ([]mdi.Bost, error)
	Create(mdi.Bost) (uint64, error)
	Update(bostID uint64, bost mdi.Bost) error
	Remove(bostID uint64) (bool, error)
}

type DummyBostService struct{}

func NewDummyBostService() BostService {
	return &DummyBostService{}
}

func (d *DummyBostService) Describe(bostID uint64) (*mdi.Bost, error) {
	if bostID < 1 || bostID > uint64(len(mdi.DataModel)) {
		return nil, ErrDataIDRange
	}
	return &mdi.DataModel[bostID-1], nil
}

func (d *DummyBostService) List(cursor uint64, limit uint64) ([]mdi.Bost, error) {
	if len(mdi.DataModel) == 0 {
		return nil, ErrEmptyModel
	}
	pos := cursor - 1
	if int(pos) >= len(mdi.DataModel) {
		return nil, ErrPosOutRange
	}
	if int(pos+limit) >= len(mdi.DataModel) {
		return mdi.DataModel[pos:], nil
	}
	return mdi.DataModel[pos : pos+limit], nil
}

func (d *DummyBostService) Create(bost mdi.Bost) (uint64, error) {
	mdi.DataModel = append(mdi.DataModel, bost)
	return uint64(len(mdi.DataModel)), nil
}

func (d *DummyBostService) Update(bostID uint64, bost mdi.Bost) error {
	fmt.Println(bost)
	if bostID < 1 || bostID > uint64(len(mdi.DataModel)) {
		return ErrDataIDRange
	}
	fmt.Println(mdi.DataModel[bostID-1])
	mdi.DataModel[bostID-1].Name = bost.Name
	fmt.Println(mdi.DataModel[bostID-1])
	return nil
}

func (d *DummyBostService) Remove(bostID uint64) (bool, error) {
	if bostID < 1 || bostID > uint64(len(mdi.DataModel)) {
		return false, ErrDataIDRange
	}
	mdi.DataModel = append(mdi.DataModel[:bostID-1], mdi.DataModel[bostID:]...)
	return true, nil
}
