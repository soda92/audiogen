package main

// typedef unsigned char Uint8;
// void SineWave(void *userdata, Uint8 *stream, int len);
import "C"
import (
	// "log"
	"math"
	"reflect"
	"unsafe"

	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
)

const (
	toneHz   = 440
	sampleHz = 48000
	dPhase   = 2 * math.Pi * toneHz / sampleHz
)

//export SineWave
func SineWave(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	var phase float64
	for i := 0; i < n; i += 2 {
		phase += dPhase
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128)
		buf[i] = sample
		buf[i+1] = sample
	}
}

// func sdl2_play() {
// 	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer sdl.Quit()

// 	spec := &sdl.AudioSpec{
// 		Freq:     sampleHz,
// 		Format:   sdl.AUDIO_U8,
// 		Channels: 2,
// 		Samples:  sampleHz,
// 		Callback: sdl.AudioCallback(C.SineWave),
// 	}
// 	if err := sdl.OpenAudio(spec, nil); err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	sdl.PauseAudio(false)
// 	sdl.Delay(5000) // play audio for long enough to understand whether it works
// 	sdl.CloseAudio()
// }

func sdl3_play() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(sdl.Path())
	defer sdl.Quit()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("Hello world", 500, 500, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()
	defer window.Destroy()

	renderer.SetDrawColor(255, 255, 255, 255)

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		renderer.DebugText(50, 50, "Hello world")
		renderer.Present()

		return nil
	})
}