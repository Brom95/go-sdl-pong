package base

import "sync"

var (
	World *world = &world{}
)

type world struct {
	Screen          *Screen
	KeyState        []uint8
	ObjectsToDraw   map[string]Drawer
	DrawMutex       *sync.RWMutex
	UpdateMutex     *sync.RWMutex
	ObjectsToUpdate map[string]Updater
}

func NewWorld(screen *Screen, keyState []uint8) {
	World.Screen = screen
	World.KeyState = keyState
	World.ObjectsToDraw = make(map[string]Drawer)
	World.ObjectsToUpdate = make(map[string]Updater)
	World.DrawMutex = &sync.RWMutex{}
	World.UpdateMutex = &sync.RWMutex{}
}
func (w *world) Update() {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	w.UpdateMutex.Lock()
	defer w.UpdateMutex.Unlock()
	objects := w.ObjectsToUpdate
	for _, updater := range objects {
		updater := updater
		wg.Add(1)
		go func() {
			defer wg.Done()
			updater.Update()
		}()
	}
}
func (w *world) Draw() {
	drawers := make([]Drawer, 0, len(w.ObjectsToDraw))
	wg := sync.WaitGroup{}
	w.DrawMutex.Lock()
	defer w.DrawMutex.Unlock()
	defer wg.Wait()
	for _, drawer := range w.ObjectsToDraw {
		drawer := drawer
		drawers = append(drawers, drawer)
		wg.Add(1)
		go func() {
			defer wg.Done()
			drawer.Draw()

		}()
	}
	for i, drawer := range drawers {
		for j := i + 1; j < len(drawers); j++ {
			wg.Add(1)
			drawer := drawer
			j := j

			go func() {
				defer wg.Done()
				drawer.Collide(drawers[j])
				drawers[j].Collide(drawer)
			}()
		}
	}
}
func (w *world) AddObjectToDraw(name string, drawer Drawer) {
	w.DrawMutex.RLock()
	defer w.DrawMutex.RUnlock()
	w.ObjectsToDraw[name] = drawer
}
func (w *world) AddObjectToUpdate(name string, updater Updater) {
	w.UpdateMutex.RLock()
	defer w.UpdateMutex.RUnlock()

	w.ObjectsToUpdate[name] = updater
}
func (w *world) RemoveObjectToDraw(name string) {
	delete(w.ObjectsToDraw, name)
}
func (w *world) GetObjectToDraw(name string) Drawer {
	w.DrawMutex.Lock()
	defer w.DrawMutex.Unlock()
	return w.ObjectsToDraw[name]
}
func (w *world) GetObjectToUpdate(name string) Updater {
	w.UpdateMutex.Lock()
	defer w.UpdateMutex.Unlock()

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
