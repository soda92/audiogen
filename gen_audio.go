package main

// typedef unsigned char Uint8;
// void SineWave(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"log"
	"math"
	"time"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	toneHz   = 440
	sampleHz = 48000
	dPhase   = 2 * math.Pi * toneHz / sampleHz
)

//export SineWave
func SineWave(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := unsafe.Slice(stream, n)
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	var phase float64
	for i := 0; i < n; i += 2 {
		phase += dPhase
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128)
		buf[i] = sample
		buf[i+1] = sample
	}
}

func sdl2_play(play chan int) {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Println(err)
		return
	}
	defer sdl.Quit()

	spec := &sdl.AudioSpec{
		Freq:     sampleHz,
		Format:   sdl.AUDIO_U8,
		Channels: 2,
		Samples:  sampleHz,
		Callback: sdl.AudioCallback(C.SineWave),
	}

	if err := sdl.OpenAudio(spec, nil); err != nil {
		log.Println(err)
		return
	}
	timer1 := time.NewTimer(3 * time.Second)

	for {
		select {
		case i := <-play:
			if i == 0 {
				sdl.PauseAudio(true)
			}
			if i == 1 {
				timer1 = time.NewTimer(3 * time.Second)
				sdl.PauseAudio(false)
			}
		case <-timer1.C:
			sdl.PauseAudio(true)
		}
	}
}
