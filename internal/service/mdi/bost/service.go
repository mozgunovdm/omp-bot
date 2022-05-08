package bost

import (
	"errors"

	"github.com/mozgunovdm/omp-bot/internal/model/mdi"
)

var ErrDataIDRange = errors.New("id out of range data")
var ErrEmptyModel = errors.New("empty list")
var ErrPosOutRange = errors.New("position out of range")

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
	if int(cursor) >= len(mdi.DataModel) {
		return nil, ErrPosOutRange
	}
	if int(cursor+limit) >= len(mdi.DataModel) {
		return mdi.DataModel[cursor:], nil
	}
	return mdi.DataModel[cursor : cursor+limit], nil
}

func (d *DummyBostService) Create(bost mdi.Bost) (uint64, error) {
	mdi.DataModel = append(mdi.DataModel, bost)
	return uint64(len(mdi.DataModel)), nil
}

func (d *DummyBostService) Update(bostID uint64, bost mdi.Bost) error {
	if bostID < 1 || bostID > uint64(len(mdi.DataModel)) {
		return ErrDataIDRange
	}
	mdi.DataModel[bostID-1] = bost
	return nil
}

func (d *DummyBostService) Remove(bostID uint64) (bool, error) {
	if bostID < 1 || bostID > uint64(len(mdi.DataModel)) {
		return false, ErrDataIDRange
	}
	mdi.DataModel = append(mdi.DataModel[:bostID-1], mdi.DataModel[bostID:]...)
	return true, nil
}
