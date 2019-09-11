package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/pkg/errors"
)

//MD: 590002150112233445566778899AABBCCDDEEFF0DEADC0DEC3
/*
 Keep getting these errors: 2019/09/10 11:23:49 can't accept: skt: unsupported sco packet: 40 0C 01 53
 Worked better after changing go-ble source, specificly removing the return statements in
 go-ble/ble/linux/hci/hci.go lines 292-322 (function sktLoop). This basically means that it will ignore the errors and keep scannign

*/
var (
	device = flag.String("device", "default", "implementation of ble")
	dup    = flag.Bool("dup", true, "disallow duplicate reported")
	c      = make(chan packet)
)

type packet struct {
	md   string
	rssi int
}

func main() {
	go firestorePusher(c, 0)

	d, err := dev.NewDevice(*device)
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	// Scan for specified durantion, or until interrupted by user.
	chkErr(ble.Scan(context.Background(), *dup, advHandler, nil))
}

func advHandler(a ble.Advertisement) {
	md := a.ManufacturerData()
	mdl := len(md)
	if len(md) == 25 && uint8Equals(md[0:2], []uint8{0x59, 0x00}) && uint8Equals(md[mdl-5:mdl-1], []uint8{0xDE, 0xAD, 0xC0, 0xDE}) {
		p := packet{hex.EncodeToString(md), a.RSSI()}
		fmt.Printf("%d, %s\n", p.rssi, p.md)
		c <- p
	}
}

func chkErr(err error) {
	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		fmt.Printf("done\n")
	case context.Canceled:
		fmt.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}

func firestorePusher(c chan packet, id int) {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("raspi-admin.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	for {
		p := <-c
		fmt.Printf("ID: %d got packet %d\n", id, p.rssi)
		// events0 is the collection that the data will be pushed to.
		// The idea is to use one collection per gateway to differentiate between locations
		_, _, err = client.Collection("events0").Add(ctx, map[string]interface{}{
			"raw":       p.md,
			"rssi":      p.rssi,
			"uuid":      p.md[8:40],
			"timestamp": time.Now().UnixNano() / 1000000,
		})
		if err != nil {
			log.Fatalf("Failed adding data: %v", err)
		}
	}

}

func uint8Equals(x []uint8, y []uint8) bool {
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
