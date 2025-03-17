package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MatthiasKunnen/go-wayland/wayland/client"
)

func main() {
	display, err := client.Connect("")
	if err != nil {
		panic(err)
	}

	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		log.Fatal("env XDG_RUNTIME_DIR not set")
	}

	display.SetErrorHandler(func(dee client.DisplayErrorEvent) {
		log.Fatalf("display error event: %v", dee)
	})

	registry, err := display.GetRegistry()
	if err != nil {
		log.Fatalf("error getting registry: %v", err)
	}

	ddmCh := make(chan *client.DataDeviceManager, 1)
	seatCh := make(chan *client.Seat, 1)

	registry.SetGlobalHandler(func(rge client.RegistryGlobalEvent) {
		switch rge.Interface {
		case client.DataDeviceManagerInterfaceName:
			ddm := client.NewDataDeviceManager(display.Context())
			err := registry.Bind(rge.Name, rge.Interface, rge.Version, ddm)
			if err != nil {
				log.Fatalf("unable to bind ddm interface: %v", err)
			}
			ddmCh <- ddm
		case client.SeatInterfaceName:
			seat := client.NewSeat(display.Context())

			seat.SetNameHandler(func(sne client.SeatNameEvent) {
				fmt.Println(sne.Name)
			})

			err := registry.Bind(rge.Name, rge.Interface, rge.Version, seat)
			if err != nil {
				log.Fatalf("unable to bind seat interface: %v", err)
			}

			seatCh <- seat
		}
	})

	// Wait for interface to register
	display.Roundtrip()
	// Wait for handler events
	display.Roundtrip()

	ddm := <-ddmCh
	close(ddmCh)

	seat := <-seatCh
	close(seatCh)

	dev, err := ddm.GetDataDevice(seat)
	if err != nil {
		log.Fatalf("error getting data device: %v", err)
	}

	dev.SetDataOfferHandler(func(dddoe client.DataDeviceDataOfferEvent) {
		println("offer")
		dddoe.Id.SetOfferHandler(func(dooe client.DataOfferOfferEvent) {
			fmt.Println(dooe.MimeType)
		})
		display.Roundtrip()
	})

	dev.SetSelectionHandler(func(ddse client.DataDeviceSelectionEvent) {
		println("selection")
		ddse.Id.SetOfferHandler(func(dooe client.DataOfferOfferEvent) {
			fmt.Println(dooe.MimeType)
		})
		display.Roundtrip()
	})

	display.Roundtrip()

	println("idle")

	for {
		dev.Context().Dispatch()
	}
}
