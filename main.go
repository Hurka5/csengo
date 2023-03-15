package gui

import (
  "time"
  "math"
  "strings"
  "strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/dialog"
  "fyne.io/fyne/v2/layout"
  "fyne.io/fyne/v2/container"
)

var (

  becsengok [][]int = [][]int{
    []int{8,00},
    []int{8,55},
    []int{9,50},
    []int{10,45},
    []int{11,40},
    []int{12,35},
    []int{13,25},
    []int{14,15},
  }
  kicsengok [][]int = [][]int{
    []int{8,45},
    []int{9,40},
    []int{10,35},
    []int{11,30},
    []int{12,25},
    []int{13,20},
    []int{14,10},
    []int{15,00},
  }

  // Gui
  be        *widget.Label
  beHangero *widget.Label
  beSlider  *widget.Slider
  ki        *widget.Label
  kiHangero *widget.Label
  kiSlider  *widget.Slider
)

func chooseFile(w fyne.Window, h *widget.Label) {
  dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
    path := ""
      if err != nil {
        dialog.ShowError(err, w)
        return
      }
      if file != nil {
        rawpath := file.URI().String()
        path = strings.TrimPrefix(rawpath, "file://")
      }
      h.SetText(path)
  }, w)
}

func Run() {

  // Init
  a := app.New()
	w := a.NewWindow("Csengő")

  // Systray logic
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Csengő",
			fyne.NewMenuItem("Megjelenítés", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
  

  // Left layout
  var leftLayout *fyne.Container
  {
    title := widget.NewLabel("BE")
    space1 := widget.NewLabel("")
	  be = widget.NewLabel("/path/path/path")
    button := widget.NewButton("Csere", func() {
      chooseFile(w, be)
	  })
    hangero := widget.NewLabel("Hangerő:")
    space2 := widget.NewLabel("")
    beHangero = widget.NewLabel("0%")

    beSlider := widget.NewSlider(0.0, 100.0)
    beSlider.Value = 100.0
    beHangero.SetText(strconv.Itoa(int(math.Floor(beSlider.Value))) + "%")
    beSlider.OnChanged = func(f float64) {
      beHangero.SetText(strconv.Itoa(int(math.Floor(f))) + "%")
    }

	  leftLayout = container.New(layout.NewFormLayout(),title,space1,button,be,hangero,space2,beHangero,beSlider)
  }

  // Right Layout
  var rightLayout *fyne.Container
  {
    title := widget.NewLabel("KI")
    space1 := widget.NewLabel("")
	  ki = widget.NewLabel("/path/path/path")
    button := widget.NewButton("Csere", func() {
      chooseFile(w, ki)
	  })
    hangero := widget.NewLabel("Hangerő:")
    space2 := widget.NewLabel("")
    kiHangero := widget.NewLabel("0%")

    kiSlider := widget.NewSlider(0.0, 100.0)
    kiSlider.Value = 100.0
    kiHangero.SetText(strconv.Itoa(int(math.Floor(kiSlider.Value))) + "%")
    kiSlider.OnChanged = func(f float64) {
      kiHangero.SetText(strconv.Itoa(int(math.Floor(f))) + "%")
    }

	  rightLayout = container.New(layout.NewFormLayout(),title,space1,button,ki,hangero,space2,kiHangero,kiSlider)
  }

  // Layout
	content := container.New(layout.NewGridLayout(2), leftLayout, rightLayout)
	w.SetContent(content)

  // Csenget
	go func() {
		for range time.Tick(time.Second) {
      hours, minutes, seconds := time.Now().Clock()
      for _, idopont := range becsengok {
        if hours == idopont[0] && minutes == idopont[1] && seconds == 0 {
          PlaySong(be.Text)
        }
      }
      for _, idopont := range kicsengok {
        if hours == idopont[0] && minutes == idopont[1] && seconds == 0 {
          PlaySong(be.Text)
        }
      }

      // Debug code
      /*if hours == 21 && minutes == 33 && seconds == 0 {
        PlaySong(be.Text)
      }
      if hours == 21 && minutes == 34 && seconds == 0 {
        PlaySong(be.Text)
      }*/
		}
	}()

  // Change what closing the window does
	w.SetCloseIntercept(func() {
		w.Hide()
	})

  // 
	w.ShowAndRun()
}




