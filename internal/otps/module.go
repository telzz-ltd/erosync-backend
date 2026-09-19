package otps

type Module struct {
	Service Service
}

func New(repo Repository) *Module {
	return &Module{
		Service: Service{
			repo: repo,
		},
	}
}
