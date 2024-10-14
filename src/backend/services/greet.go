package services

type Hello struct{}

func (g *Hello) World(name string) string {
	return "Hello " + name + "!"
}
