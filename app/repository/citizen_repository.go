package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"quest/model"

	"gorm.io/gorm"
)

const (
	urlAPI = "http://169.51.195.62:30174"
)

type CitizenRepository struct {
	db *gorm.DB
}

func NewCitizenRepository(db *gorm.DB) *CitizenRepository {
	return &CitizenRepository{db: db}
}

func (cr *CitizenRepository) GetAllCitizens() ([]*model.Citizen, error) {
	var citizens []*model.Citizen
	err := cr.db.Find(&citizens).Error
	if err != nil {
		return nil, err
	}

	return citizens, nil
}

func (cr *CitizenRepository) GetCitizenByID(id uint) (*model.Citizen, error) {
	var citizen model.Citizen
	err := cr.db.First(&citizen, id).Error
	if err != nil {
		return nil, err
	}

	return &citizen, nil
}

func (cr *CitizenRepository) GetCitizenDocuments(id uint) ([]*model.Document, error) {
	var documents []*model.Document
	err := cr.db.Where("citizen_id = ?", id).Find(&documents).Error
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (cr *CitizenRepository) CreateCitizen(citizen *model.Citizen) error {
	return cr.db.Create(citizen).Error
}

func (cr *CitizenRepository) UpdateCitizen(citizen *model.Citizen) error {
	return cr.db.Save(citizen).Error
}

func (cr *CitizenRepository) DeleteCitizen(citizen *model.Citizen) error {
	return cr.db.Delete(citizen).Error
}

func (cr *CitizenRepository) LoginCitizen(citizenID uint, password string) (error, bool) {
	var citizen *model.Citizen
	err := cr.db.Where("id = ? and password = ?", citizenID, password).Find(&citizen).Error
	if err != nil {
		return err, false
	}
	if citizen.ID == 0 {
		return nil, false
	}
	return nil, true
}

func (cr *CitizenRepository) CreateCitizenOnAPI(citizen *model.Citizen) error {
	data := struct {
		Id           uint   `json:"id"`
		Name         string `json:"name"`
		Address      string `json:"address"`
		Email        string `json:"email"`
		OperatorId   uint   `json:"operatorId"`
		OperatorName string `json:"operatorName"`
	}{
		Id:           citizen.ID,
		Name:         citizen.Name,
		Address:      citizen.Address,
		Email:        citizen.Email,
		OperatorId:   citizen.OperatorID,
		OperatorName: "test operator",
	}
	rawData, _ := json.Marshal(data)
	res, err := http.Post(urlAPI+"/apis/registerCitizen", "application/json", bytes.NewReader(rawData))
	if err != nil {
		return err
	}
	if res.StatusCode != 201 {
		body, _ := io.ReadAll(res.Body)
		return errors.New(string(body))
	}
	return nil
}
