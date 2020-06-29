package src

import (
	"github.com/gdamore/tcell"
	"os"
)

// renderer is a system
type RendererSys struct {
	*SystemBase
	RenderComps []*RenderComp
	Screen Screen
}

type RenderComp struct {
	// store data about a renderable here
	Depth int
	Pos   Vec
	View  uint8
	Buffer tcell.CellBuffer
}

func (R *RendererSys) Init(ECS *ECS){
	screen, err := tcell.NewScreen()
	if err != nil{
		LogErr(err)
	}
	R.Screen.Screen = screen
	if err = screen.Init(); err != nil{
		LogErr(err)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
	screen.EnableMouse()
	screen.Clear()

	R.Screen.Width = DEFAULT_WIDTH
	R.Screen.Height = DEFAULT_HEIGHT
	R.Screen.Resize()
	R.Screen.ZBuffer = make([]tcell.CellBuffer, 3)
	for i:=0; i<Z_DEPTH; i++ {
		R.Screen.ZBuffer[i] = tcell.CellBuffer{}
		R.Screen.ZBuffer[i].Resize(R.Screen.Width, R.Screen.Height)
	}

	go R.Screen.Poll()
}

func (R *RendererSys) Close(ECS *ECS){

}

func (R *RendererSys) AddEntity(Entity *Entity, RenderComp *RenderComp){
	R.Entities = append(R.Entities, Entity)
	R.RenderComps = append(R.RenderComps, RenderComp)
	R.Size++
}

func (R *RendererSys) Update(ECS *ECS){

	// process camera movement
	if InputBuffer.KeyPressed == 'a'{
		R.Screen.Cam = R.Screen.Cam.Add(V2(1,0))
	}else if InputBuffer.KeyPressed == 'd'{
		R.Screen.Cam = R.Screen.Cam.Add(V2(-1,0))
	}else if InputBuffer.KeyPressed == 'w'{
		R.Screen.Cam = R.Screen.Cam.Add(V2(0,1))
	}else if InputBuffer.KeyPressed == 's'{
		R.Screen.Cam = R.Screen.Cam.Add(V2(0,-1))
	}else if InputBuffer.CtrlKeyPressed == tcell.KeyEscape{
		Running = false
	}

	// render each RenderComp
	for _, r := range R.RenderComps{
		width, height := r.Buffer.Size()
		// iterate over each cell in the cell buffer
		for x:=0; x<width;x ++{
			for y:=0; y<height;y ++{
				rune, _, style, _ := r.Buffer.GetContent(x,y)
				// draw at the offset from the RenderComp position
				R.Screen.Char(rune, r.Pos.Add(V2i(x,y)), style, r.View, r.Depth)
			}
		}
	}
	R.Screen.Draw()
}

func (R *RendererSys) Remove(){
}