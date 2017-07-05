/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package creator

import (
	"fmt"

	"math"

	"github.com/unidoc/unidoc/common"
	"github.com/unidoc/unidoc/pdf/model/fonts"
)

// A subchapter simply represents a subchapter pertaining to a specific chapter.  It can contain multiple
// drawables, just like a chapter.
type subchapter struct {
	chapterNum    int
	subchapterNum int
	title         string
	heading       *paragraph

	contents []Drawable

	// Positioning: relative / absolute.
	positioning positioning

	// Absolute coordinates (when in absolute mode).
	xPos, yPos float64

	// Margins to be applied around the block when drawing on page.
	margins margins

	// Chapter sizing is set to occupy available space.
	sizing Sizing
}

func (c *Creator) NewSubchapter(ch *chapter, title string) *subchapter {
	subchap := &subchapter{}

	ch.subchapters++
	subchap.subchapterNum = ch.subchapters

	subchap.chapterNum = ch.number
	subchap.title = title

	heading := fmt.Sprintf("%d.%d %s", subchap.chapterNum, subchap.subchapterNum, title)
	p := NewParagraph(heading)

	p.SetFontSize(14)
	p.SetFont(fonts.NewFontHelvetica()) // bold?

	subchap.heading = p
	subchap.contents = []Drawable{}

	// Chapter sizing is fixed to occupy available size.
	subchap.sizing = SizingOccupyAvailableSpace

	// Add subchapter to ch.
	ch.Add(subchap)

	return subchap
}

// Chapter sizing is fixed to occupy available space in the drawing context.
func (subchap *subchapter) GetSizingMechanism() Sizing {
	return subchap.sizing
}

// Chapter height is a sum of the content heights.
func (subchap *subchapter) Height() float64 {
	h := float64(0)
	for _, d := range subchap.contents {
		h += d.Height()
	}
	return h
}

// Chapter width is the maximum of the content widths.
func (subchap *subchapter) Width() float64 {
	maxW := float64(0)
	for _, d := range subchap.contents {
		maxW = math.Max(maxW, d.Width())
	}
	return maxW
}

// Set absolute coordinates.
func (subchap *subchapter) SetPos(x, y float64) {
	subchap.positioning = positionAbsolute
	subchap.xPos = x
	subchap.yPos = y
}

// Set chapter margins.  Typically not needed as the page margins are used.
func (subchap *subchapter) SetMargins(left, right, top, bottom float64) {
	subchap.margins.left = left
	subchap.margins.right = right
	subchap.margins.top = top
	subchap.margins.bottom = bottom
}

// Get the subchapter margins: left, right, top, bototm.
func (subchap *subchapter) GetMargins() (float64, float64, float64, float64) {
	return subchap.margins.left, subchap.margins.right, subchap.margins.top, subchap.margins.bottom
}

// Add a new drawable to the chapter.
func (subchap *subchapter) Add(d Drawable) {
	switch d.(type) {
	case *chapter, *subchapter:
		common.Log.Debug("Error: Cannot add chapter or subchapter to a subchapter")
	case *paragraph, *image, *block:
		subchap.contents = append(subchap.contents, d)
	default:
		common.Log.Debug("Unsupported: %T", d)
	}
}

// XXX/FIXME: Need to know actual page numbers to keep track for TOC.
// TODO: Add page number to context... ?
func (subchap *subchapter) GeneratePageBlocks(ctx drawContext) ([]*block, drawContext, error) {
	ctx.Y += subchap.margins.top
	blocks, ctx, err := subchap.heading.GeneratePageBlocks(ctx)
	if err != nil {
		return blocks, ctx, err
	}

	for _, d := range subchap.contents {
		newBlocks, c, err := d.GeneratePageBlocks(ctx)
		if err != nil {
			return blocks, ctx, err
		}
		if len(newBlocks) < 1 {
			continue
		}

		// The first block is always appended to the last..
		blocks[len(blocks)-1].mergeBlocks(newBlocks[0])
		blocks = append(blocks, newBlocks[1:]...)

		ctx = c
	}

	return blocks, ctx, nil
}
