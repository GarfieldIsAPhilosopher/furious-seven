package v3action

import "github.com/cloudfoundry-incubator/furious-seven/api/v3"

type Actor struct {
	client v3api.Clientv3
}

func NewActor(client v3api.Clientv3) Actor {
	return Actor{
		client: client,
	}
}
