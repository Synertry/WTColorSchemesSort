/*
 *           Copyright Synertry 2022.
 *  Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE_1_0.txt or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package WTColorSchemesSort

import (
	"encoding/json"
)

// colorSchemes is a helper struct for parsing and sorting
type colorSchemes struct {
	Schemes []colorScheme `json:"schemes,required"`
}

// colorScheme ordered after ColorScheme.cpp from TerminalSettingsModel
// https://github.com/microsoft/terminal/blob/094273b995352ff9d80f36279c297488b756df1d/src/cascadia/TerminalSettingsModel/ColorScheme.cpp#L151
type colorScheme struct {
	Name                string `json:"name,omitempty"`
	Foreground          string `json:"foreground,omitempty"`
	Background          string `json:"background,omitempty"`
	CursorColor         string `json:"cursorColor,omitempty"`
	SelectionBackground string `json:"selectionBackground,omitempty"`
	Black               string `json:"black,omitempty"`
	Red                 string `json:"red,omitempty"`
	Green               string `json:"green,omitempty"`
	Yellow              string `json:"yellow,omitempty"`
	Blue                string `json:"blue,omitempty"`
	Purple              string `json:"purple,omitempty"`
	Cyan                string `json:"cyan,omitempty"`
	White               string `json:"white,omitempty"`
	BrightBlack         string `json:"brightBlack,omitempty"`
	BrightRed           string `json:"brightRed,omitempty"`
	BrightGreen         string `json:"brightGreen,omitempty"`
	BrightYellow        string `json:"brightYellow,omitempty"`
	BrightBlue          string `json:"brightBlue,omitempty"`
	BrightPurple        string `json:"brightPurple,omitempty"`
	BrightCyan          string `json:"brightCyan,omitempty"`
	BrightWhite         string `json:"brightWhite,omitempty"`
}

// parseSchemes core sorting mechanism
func parseSchemes(str string) (colorSchemes colorSchemes, err error) {
	err = json.Unmarshal([]byte(str), &colorSchemes.Schemes)
	return
}