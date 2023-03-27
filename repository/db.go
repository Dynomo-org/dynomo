package repository

import "context"

const (
	collectionMstApp = "mst_app"
)

func (r *Repository) GetAllCollection(ctx context.Context) (interface{}, error) {
	var data interface{}
	err := r.db.NewRef(collectionMstApp).Get(ctx, data)
	return data, err
}

func (r *Repository) InsertMasterApp(ctx context.Context, app MasterApp) error {
	_, err := r.db.NewRef(collectionMstApp).Push(ctx, app)
	return err
}

func (r *Repository) CreateNewMasterApp(ctx context.Context, name string) error {
	_, err := r.db.NewRef(collectionMstApp).Push(ctx, MasterApp{Name: name})
	return err
}
