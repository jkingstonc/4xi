package src
import "C"

type Client struct {
	ECS *ECS
}

func NewClient(addr string) *Client{
	// create a client by registering all the relevant ECS systems
	ecs := NewECS(CLIENT)
	ecs.RegisterSystem(&NetworkSys{SystemBase: NewSysBase(ecs), ServerAddress: addr})
	ecs.RegisterSystem(&PlayerSys{SystemBase: NewSysBase(ecs)})
	ecs.RegisterSystem(&RendererSys{SystemBase: NewSysBase(ecs)})
	ecs.RegisterSystem(&WorldSys{SystemBase: NewSysBase(ecs)})
	ecs.RegisterSystem(&EmpireSys{SystemBase: NewSysBase(ecs)})
	ecs.RegisterSystem(&SettlementSys{SystemBase: NewSysBase(ecs)})
	return &Client{
		ECS:        ecs,
	}
}

func (Client *Client) Init(){
	Client.ECS.Init()
}

// process all updatable entities
func (Client *Client) Process() {
	Client.ECS.Update()
}

func (Client *Client) Close(){
	Client.ECS.Close()
}