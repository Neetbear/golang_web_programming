package internal

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (*CreateResponse, error) {
	membership, err := app.repository.Create(request)
	if err != nil {
		return nil, err
	}
	return &CreateResponse{membership.UserName, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (*UpdateResponse, error) {
	membership, err := app.repository.Update(request)
	if err != nil {
		return nil, err
	}
	return &UpdateResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	return app.repository.Delete(id)
}
