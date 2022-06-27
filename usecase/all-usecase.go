package usecase

import "kereta/models"

func (r *UcSecond) GetDataKategori() ([]models.Kategori, error) {
	var Modelmhs []models.Kategori
	err := r.query.FindAll(&Modelmhs)
	if err != nil {
		return Modelmhs, err
	}
	return Modelmhs, nil
}

func (r *UcSecond) InsertDataKategori(data models.Kategori) error {
	err := r.query.InsertData(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UcSecond) UpdateDataKategori(data models.Kategori) error {
	var Where map[string]interface{}
	Where["id"] = data.Id

	var dataUpdates map[string]interface{}
	dataUpdates["nama"] = data.Nama

	err := r.query.UpdateData(&data, Where, dataUpdates)
	if err != nil {
		return err
	}
	return nil
}

func (r *UcSecond) DeleteDataKereta(id int) error {
	detail, err := r.query.ObjectByIdDetail(id)
	if err != nil {
		return err
	}
	err = r.query.DeleteData(detail)
	if err != nil {
		return err
	}
	return nil
}
