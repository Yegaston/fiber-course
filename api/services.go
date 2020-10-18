package api

type Services struct {
	search MovieSearch
}

func newServices() Services {
	return Services{
		search: &MovieService{},
	}
}

type WebServices struct {
	S Services
}

func Start() *WebServices {
	return WebServices(S: NewServices)
}