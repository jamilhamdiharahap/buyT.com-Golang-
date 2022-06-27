package repository

func (r *Repo) Create(i interface{}) error {
	return r.app.Create(i).Error
}
