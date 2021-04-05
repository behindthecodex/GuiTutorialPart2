// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func main() {
	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
		walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
	})

	var mw *walk.MainWindow
	var urlTE *walk.LineEdit
	var outTE *walk.TextEdit

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "My Awesome Scraper",
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Title:  "Options",
				Layout: Grid{Columns: 2},
				Children: []Widget{
					LineEdit{AssignTo: &urlTE},
					PushButton{
						Text: "Exe",
						OnClicked: func() {

							outTE.SetText(urlTE.Text())
						},
					},
				},
			},
			Label{
				Text: "animal:",
			},
			TextEdit{
				AssignTo: &outTE,
				ReadOnly: true,
				//Text:     fmt.Sprintf("%+v", animal),
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
