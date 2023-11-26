package egeom

import (
	"github.com/watermint/toolbox/essentials/islet/eidiom/eoutcome"
	"strings"
)

type Position int

// Locate defines top-left corner of obj location. Padding ignored if a position is center.
// Returns (0, 0) if undefined position
func (z Position) Locate(base, obj Rectangle, padding Padding) Point {
	pad := padding.Define(obj)
	centerX := func() int {
		return (base.Width() - obj.Width()) / 2
	}
	centerY := func() int {
		return (base.Height() - obj.Height()) / 2
	}
	rightX := func() int {
		return base.Width() - obj.Width() - pad.X()
	}
	bottomY := func() int {
		return base.Height() - obj.Height() - pad.Y()
	}
	switch z {
	case PositionTopLeft:
		return pad
	case PositionTopCenter:
		return NewPoint(
			centerX(),
			pad.Y(),
		)
	case PositionTopRight:
		return NewPoint(
			rightX(),
			pad.Y(),
		)
	case PositionCenterLeft:
		return NewPoint(
			pad.X(),
			centerY(),
		)
	case PositionCenter:
		return NewPoint(
			centerX(),
			centerY(),
		)
	case PositionCenterRight:
		return NewPoint(
			rightX(),
			centerY(),
		)
	case PositionBottomLeft:
		return NewPoint(
			pad.X(),
			bottomY(),
		)
	case PositionBottomCenter:
		return NewPoint(
			centerX(),
			bottomY(),
		)
	case PositionBottomRight:
		return NewPoint(
			rightX(),
			bottomY(),
		)
	}

	return ZeroPoint
}

const (
	PositionUndefined Position = iota
	PositionTopLeft
	PositionTopCenter
	PositionTopRight
	PositionCenterLeft
	PositionCenter
	PositionCenterRight
	PositionBottomLeft
	PositionBottomCenter
	PositionBottomRight
)

func ParsePosition(p string) (Position, eoutcome.ParseOutcome) {
	switch strings.ToLower(p) {
	case "center", "centre":
		return PositionCenter, eoutcome.NewParseSuccess()
	case "topleft", "top_left", "top-left", "top left":
		return PositionTopLeft, eoutcome.NewParseSuccess()
	case "topcenter", "top_center", "top-center", "top center", "topcentre", "top_centre", "top-centre", "top centre":
		return PositionTopCenter, eoutcome.NewParseSuccess()
	case "topright", "top_right", "top-right", "top right":
		return PositionTopRight, eoutcome.NewParseSuccess()
	case "centerleft", "center_left", "center-left", "center left", "centreleft", "centre_left", "centre-left", "centre left":
		return PositionCenterLeft, eoutcome.NewParseSuccess()
	case "centerright", "center_right", "center-right", "center right", "centreright", "centre_right", "centre-right", "centre right":
		return PositionCenterRight, eoutcome.NewParseSuccess()
	case "bottomleft", "bottom_left", "bottom-left", "bottom left":
		return PositionBottomLeft, eoutcome.NewParseSuccess()
	case "bottomcenter", "bottom_center", "bottom-center", "bottom center", "bottomcentre", "bottom_centre", "bottom-centre", "bottom centre":
		return PositionBottomCenter, eoutcome.NewParseSuccess()
	case "bottomright", "bottom_right", "bottom-right", "bottom right":
		return PositionBottomRight, eoutcome.NewParseSuccess()
	}
	return PositionUndefined, eoutcome.NewParseInvalidFormat("invalid position name or format")
}
