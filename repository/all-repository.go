package repository

import "kereta/models"

func (r *RepoSecond) InsertData(i interface{}) error {
	result := r.app.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepoSecond) FindAll(i interface{}) error {
	result := r.app.Find(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepoSecond) ObjectByID(id int) (model models.Kategori, err error) {
	err = r.app.First(&model).Error
	return
}
func (r *RepoSecond) ObjectByIdDetail(id int) (model models.DetailKereta, err error) {
	err = r.app.First(&model).Error
	return
}

func (r *RepoSecond) UpdateData(i interface{}, where map[string]interface{}, data map[string]interface{}) error {
	result := r.app.Model(i).Where(where).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepoSecond) DeleteData(i interface{}) error {
	result := r.app.Delete(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
