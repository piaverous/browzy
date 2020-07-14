package health

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/piaverous/browzy/cmd/utils"

	"github.com/gizak/termui/v3"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var Colors = []termui.Color{ui.ColorRed, ui.ColorCyan, ui.ColorGreen, ui.ColorBlue, ui.ColorYellow, ui.ColorMagenta}

type MasterMind struct {
	urls          []string
	checkers      []Checker
	resultMasters []ResultsMaster
	wg            *sync.WaitGroup
	rtgroup       []*widgets.Sparkline
	bcgroup       []*widgets.BarChart
	slg           *widgets.SparklineGroup
	grid          *ui.Grid
}

func (mm *MasterMind) getTitle(url string, min float64, max float64) string {
	str := url
	n := 28

	if len(str) > n {
		str = str[:n] + "[...]"
	}
	return fmt.Sprintf("%s - Min: %.0fms - Max: %.0fms", str, min, max)
}

func (mm *MasterMind) Start() {
	for i, _ := range mm.urls {
		go mm.checkers[i].start(mm.wg)
		go mm.resultMasters[i].consumeResults()

		mm.rtgroup[i] = widgets.NewSparkline()
		mm.rtgroup[i].Title = mm.getTitle(mm.checkers[i].url, 0, 0)
		mm.rtgroup[i].Data = mm.resultMasters[i].responseTimes
		mm.rtgroup[i].LineColor = Colors[int(math.Mod(float64(i), float64(len(Colors))))]

		mm.bcgroup[i] = widgets.NewBarChart()
		mm.bcgroup[i].Title = mm.checkers[i].url
		mm.bcgroup[i].Labels, mm.bcgroup[i].Data = mm.resultMasters[i].getStatusCodeResults()
		mm.bcgroup[i].BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
		mm.bcgroup[i].LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
		mm.bcgroup[i].NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
	}

	go mm.UpdateUI()
	mm.CreateUI()
	mm.wg.Wait()
}

func (mm *MasterMind) Stop() {
	for i, _ := range mm.urls {
		mm.checkers[i].stop()
	}
}

func (mm *MasterMind) UpdateUI() {
	tick := time.NewTicker(time.Duration(HealthCheckFrequency) * time.Millisecond)
	for {
		<-tick.C
		for i, _ := range mm.urls {
			min, max := utils.MiniMaxFloat(mm.resultMasters[i].responseTimes)
			mm.rtgroup[i].Data = mm.resultMasters[i].responseTimes
			mm.rtgroup[i].Title = mm.getTitle(mm.checkers[i].url, min, max)

			mm.bcgroup[i].Labels, mm.bcgroup[i].Data = mm.resultMasters[i].getStatusCodeResults()
		}
		ui.Render(mm.grid)
	}
}

func (mm *MasterMind) CreateUI() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// single
	mm.slg = widgets.NewSparklineGroup(mm.rtgroup...)
	mm.slg.Title = "Response Times"

	barcharts := make([]interface{}, 0)
	n := float64(len(mm.urls))
	for i, _ := range mm.urls {
		barcharts = append(barcharts, ui.NewCol(1.0/n, mm.bcgroup[i]))
	}

	termWidth, termHeight := ui.TerminalDimensions()
	mm.grid.SetRect(0, 0, termWidth, termHeight)
	mm.grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0/2, mm.slg),
			ui.NewCol(1.0/2,
				barcharts...,
			),
		),
	)

	ui.Render(mm.grid)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			mm.Stop()
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			mm.grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(mm.grid)
		}
	}
}

func CreateMasterMind(urls []string) MasterMind {
	var wg sync.WaitGroup
	var checkers = make([]Checker, len(urls))
	var resultMasters = make([]ResultsMaster, len(urls))

	for index, url := range urls {
		wg.Add(1)
		resultMasters[index] = createResultsMaster()
		checkers[index] = createChecker(url, resultMasters[index].in)
	}
	return MasterMind{
		urls,
		checkers,
		resultMasters,
		&wg,
		make([]*widgets.Sparkline, len(urls)),
		make([]*widgets.BarChart, len(urls)),
		widgets.NewSparklineGroup(),
		ui.NewGrid(),
	}
}
