// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone"
)

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlightDark).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)

	tabGap = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
)

type tabs struct {
	id     string
	height int
	width  int

	active string
	items  []string
}

func (m tabs) Init() tea.Cmd {
	return nil
}

func (m tabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.MouseClickMsg:
		if msg.Button != tea.MouseLeft {
			return m, nil
		}

		m.handleMouse(msg)
	case tea.MouseWheelMsg:
		m.handleMouse(msg)
	}
	return m, nil
}

func (m *tabs) handleMouse(msg tea.MouseMsg) {
	for _, item := range m.items {
		if zone.Get(m.id + item).InBounds(msg) {
			m.active = item
			break
		}
	}
}

func (m tabs) View() string {
	out := []string{}

	for _, item := range m.items {
		// Make sure to mark each tab when rendering.
		if item == m.active {
			out = append(out, zone.Mark(m.id+item, activeTab.Render(item)))
		} else {
			out = append(out, zone.Mark(m.id+item, tab.Render(item)))
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, out...)
	gap := tabGap.Render(strings.Repeat(" ", max(0, m.width-lipgloss.Width(row)-2)))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
	return row
}
