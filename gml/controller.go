package gml

type gameController interface {
	GameStart()
	GamePreUpdate()
	GamePostUpdate()
	GamePreDraw()
	GamePostDraw()
	implementsController()
}

type Controller struct {
}

func (controller *Controller) GamePreUpdate()        {}
func (controller *Controller) GamePostUpdate()       {}
func (controller *Controller) GamePreDraw()          {}
func (controller *Controller) GamePostDraw()         {}
func (controller *Controller) implementsController() {}
