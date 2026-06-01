package facade

import "fmt"

type TV struct {
}

func (t *TV) PowerOn() {
	fmt.Println("TV is powered on")
}

func (t *TV) PowerOff() {
	fmt.Println("TV is powered off")
}

type DVDPlayer struct {
}

func (d *DVDPlayer) PowerOn() {
	fmt.Println("DVD player is powered on")
}

func (d *DVDPlayer) PowerOff() {
	fmt.Println("DVD player is powered off")
}

type SoundSystem struct {
}

func (s *SoundSystem) PowerOn() {
	fmt.Println("Sound system is powered on")
}

func (s *SoundSystem) PowerOff() {
	fmt.Println("Sound system is powered off")
}

type HomeTheaterFacade struct {
	TV *TV
	DVDPlayer *DVDPlayer
	SoundSystem *SoundSystem
}

func NewHomeTheaterFacade(tv *TV, dvdPlayer *DVDPlayer, soundSystem *SoundSystem) *HomeTheaterFacade {
	return &HomeTheaterFacade{TV: tv, DVDPlayer: dvdPlayer, SoundSystem: soundSystem}
}

func (h *HomeTheaterFacade) WatchMovie() {
	h.TV.PowerOn()
	h.DVDPlayer.PowerOn()
	h.SoundSystem.PowerOn()
}

func (h *HomeTheaterFacade) EndMovie() {
	h.TV.PowerOff()
	h.DVDPlayer.PowerOff()
	h.SoundSystem.PowerOff()
}
