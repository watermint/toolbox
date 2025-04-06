package eg_geom

import (
	"strings"

	"github.com/watermint/toolbox/essentials/go/es_errors"
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

func ParsePosition(p string) (Position, error) {
	switch strings.ToLower(p) {
	case "center", "centre":
		return PositionCenter, nil
	case "topleft", "top_left", "top-left", "top left":
		return PositionTopLeft, nil
	case "topcenter", "top_center", "top-center", "top center", "topcentre", "top_centre", "top-centre", "top centre":
		return PositionTopCenter, nil
	case "topright", "top_right", "top-right", "top right":
		return PositionTopRight, nil
	case "centerleft", "center_left", "center-left", "center left", "centreleft", "centre_left", "centre-left", "centre left":
		return PositionCenterLeft, nil
	case "centerright", "center_right", "center-right", "center right", "centreright", "centre_right", "centre-right", "centre right":
		return PositionCenterRight, nil
	case "bottomleft", "bottom_left", "bottom-left", "bottom left":
		return PositionBottomLeft, nil
	case "bottomcenter", "bottom_center", "bottom-center", "bottom center", "bottomcentre", "bottom_centre", "bottom-centre", "bottom centre":
		return PositionBottomCenter, nil
	case "bottomright", "bottom_right", "bottom-right", "bottom right":
		return PositionBottomRight, nil
	}
	return PositionUndefined, es_errors.NewInvalidFormatError("invalid position name or format")
}
