package uc

import "fmt"

func MouseEvent() {

	mx, my := App.Ms.Px, App.Ms.Py
	dx, dy := App.mMenu.X, App.mMenu.Y

	for i, b := range btn {

		if dx+b.dispRc.Bounds().Min.X < mx && mx < dx+b.dispRc.Bounds().Max.X &&
			dy+b.dispRc.Bounds().Min.Y < my && my < dy+b.dispRc.Bounds().Max.Y {

			App.selMMenu = i
			fmt.Println("MouseEvent App.selMMenu:", App.selMMenu)
		}
	}
}
