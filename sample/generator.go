package sample

import (
	"math/rand"
	"time"

	"github.com/bugg123/grpc-intro/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//NewKeyboard(returns a new sample keyboard)
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	cpu := &pb.CPU{
		Brand:         randomCPUBrand(),
		Name:          randomCpuName(),
		NumberCoreos:  0,
		NumberThreads: 0,
		MinGhz:        0,
		MaxGhz:        0,
	}
	return cpu
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: 0,
		Memory: &pb.Memory{},
	}
}

func NewScreen() *pb.Screen {
	return &pb.Screen{
		SizeInch:   0,
		Resolution: &pb.Screen_Resolution{},
		Panel:      0,
		Multitouch: false,
	}
}

func NewLaptop() *pb.Laptop {
	return &pb.Laptop{
		Id:          "",
		Branch:      "",
		Name:        "",
		Cpu:         NewCPU(),
		Ram:         &pb.Memory{},
		Gpus:        []*pb.GPU{},
		Storages:    []*pb.Storage{},
		Screen:      &pb.Screen{},
		Keyboard:    &pb.Keyboard{},
		Weight:      nil,
		PriceUsd:    0,
		ReleaseYear: 0,
		Timestamep:  &timestamppb.Timestamp{},
	}
}

func randomCpuName() string {
	return "CpuName"
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(in ...string) string {
	size := len(in)
	if size == 0 {
		return ""
	}
	return in[rand.Intn(size)]
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
