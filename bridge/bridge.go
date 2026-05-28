package bridge

import "fmt"

type Device interface {
	PowerOn()
	PowerOff()
	SetVolume(volume int32)
	GetVolume() int32
}


type TV struct {
	Volume int32
}

func (t *TV) PowerOn() {
	fmt.Println("TV is powered on")
}

func (t *TV) PowerOff() {
	fmt.Println("TV is powered off")
}

func (t *TV) SetVolume(volume int32) {
	t.Volume = volume
	fmt.Println("TV volume is set to", volume)
}

func (t *TV) GetVolume() int32 {
	return t.Volume
}

type DVDPlayer struct {
	Volume int32
}

func (d *DVDPlayer) PowerOn() {
	fmt.Println("DVD player is powered on")
}

func (d *DVDPlayer) PowerOff() {
	fmt.Println("DVD player is powered off")
}

func (d *DVDPlayer) SetVolume(volume int32) {
	d.Volume = volume
	fmt.Println("DVD player volume is set to", volume)
}

func (d *DVDPlayer) GetVolume() int32 {
	return d.Volume
}

type BasicControl struct {
	device Device
}

func NewBasicControl(device Device) *BasicControl {
	return &BasicControl{device: device}
}

func (b *BasicControl) PowerOn() {
	b.device.PowerOn()
}

func (b *BasicControl) PowerOff() {
	b.device.PowerOff()
}

func (b *BasicControl) VolumeUp() {
	b.device.SetVolume(b.device.GetVolume() + 10)
}

func (b *BasicControl) VolumeDown() {
	b.device.SetVolume(b.device.GetVolume() - 10)
}

type AdvancedControl struct {
	device Device
}

func NewAdvancedControl(device Device) *AdvancedControl {
	return &AdvancedControl{device: device}
}

func (a *AdvancedControl) PowerOn() {
	a.device.PowerOn()
}

func (a *AdvancedControl) PowerOff() {
	a.device.PowerOff()
}

func (a *AdvancedControl) VolumeUp() {
	a.device.SetVolume(a.device.GetVolume() + 5)
}

func (a *AdvancedControl) VolumeDown() {
	a.device.SetVolume(a.device.GetVolume() - 5)
}
