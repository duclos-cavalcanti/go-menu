package term

import (
    "errors"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-org/cmd/menu/util"
)

func DrawText(tc *TerminalContext, style tcell.Style, text string) error {
	row := tc.Row
    w, _ := tc.Size()
    length := util.LengthString(text)

    if length + row > w {
        return errors.New("Unable to draw to screen, text length would exceed width.")
    } else {
	    for _, c := range []rune(text) {
	    	tc.Screen.SetContent(tc.Row, tc.Col, c, nil, style)
            tc.Row++
	    }
        return nil
    }
}

func DrawTextNewLine(tc *TerminalContext, style tcell.Style, text string) error {
    err := DrawText(tc, style, text)
    if err != nil {
        return err
    } else {
        tc.NewLine()
        return nil
    }
}
