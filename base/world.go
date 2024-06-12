package base

var (
	World *world = &world{}
)

type world struct {
	Screen          *Screen
	KeyState        []uint8
	ObjectsToDraw   map[string]Drawer
	ObjectsToUpdate map[string]Updater
}

func NewWorld(screen *Screen, keyState []uint8) {
	World.Screen = screen
	World.KeyState = keyState
	World.ObjectsToDraw = make(map[string]Drawer)
	World.ObjectsToUpdate = make(map[string]Updater)
}
func (w *world) Update() {
	for _, updater := range w.ObjectsToUpdate {
		updater.Update()
	}

}
func (w *world) Draw() {
	drawers := make([]Drawer, 0, len(w.ObjectsToDraw))
	for _, drawer := range w.ObjectsToDraw {
		drawer.Draw()
		drawers = append(drawers, drawer)
	}
	for i, drawer := range drawers {
		for j := i + 1; j < len(drawers); j++ {
			drawer.Collide(drawers[j])

		}
	}
}
func (w *world) AddObjectToDraw(name string, drawer Drawer) {
	w.ObjectsToDraw[name] = drawer
}
func (w *world) AddObjectToUpdate(name string, updater Updater) {
	w.ObjectsToUpdate[name] = updater
}
func (w *world) RemoveObjectToDraw(name string) {
	delete(w.ObjectsToDraw, name)
}
func (w *world) GetObjectToDraw(name string) Drawer {
	return w.ObjectsToDraw[name]
}
func (w *world) GetObjectToUpdate(name string) Updater {
	return w.ObjectsToUpdate[name]
}
func (w *world) AddDrawerUpdater(name string, updater DrawerUpdater) {
	w.AddObjectToDraw(name, updater)
	w.AddObjectToUpdate(name, updater)

}

type Drawer interface {
	Draw()
	Collide(Drawer)
}
type Updater interface {
	Update()
}

type DrawerUpdater interface {
	Drawer
	Updater
}
