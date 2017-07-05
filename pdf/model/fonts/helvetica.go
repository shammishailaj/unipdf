/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */
/*
 * The embedded character metrics specified in this file are distributed under the terms listed in
 * ./afms/MustRead.html.
 */

package fonts

import (
	"github.com/unidoc/unidoc/pdf/core"
	"github.com/unidoc/unidoc/pdf/model/textencoding"
)

// Font Helvetica.  Implements Font interface.
// This is a built-in font and it is assumed that every reader has access to it.
type fontHelvetica struct {
	encoder textencoding.TextEncoder
}

func NewFontHelvetica() fontHelvetica {
	font := fontHelvetica{}
	font.encoder = textencoding.NewWinAnsiTextEncoder() // Default
	return font
}

func (font fontHelvetica) SetEncoder(encoder textencoding.TextEncoder) {
	font.encoder = encoder
}

func (font fontHelvetica) GetGlyphCharMetrics(glyph string) (CharMetrics, bool) {
	metrics, has := helveticaCharMetrics[glyph]
	if !has {
		return metrics, false
	}

	return metrics, true
}

func (font fontHelvetica) ToPdfObject() core.PdfObject {
	obj := &core.PdfIndirectObject{}

	fontDict := &core.PdfObjectDictionary{
		"Type":     core.MakeName("Font"),
		"Subtype":  core.MakeName("Type1"),
		"BaseFont": core.MakeName("Helvetica"),
		"Encoding": font.encoder.ToPdfObject(),
	}

	obj.PdfObject = fontDict
	return obj
}

// Helvetica font metics loaded from afms/Helvetica.afm.  See afms/MustRead.html for license information.
var helveticaCharMetrics map[string]CharMetrics = map[string]CharMetrics{
	"A":              CharMetrics{GlyphName: "A", Wx: 667.000000, Wy: 0.000000},
	"AE":             CharMetrics{GlyphName: "AE", Wx: 1000.000000, Wy: 0.000000},
	"Aacute":         CharMetrics{GlyphName: "Aacute", Wx: 667.000000, Wy: 0.000000},
	"Abreve":         CharMetrics{GlyphName: "Abreve", Wx: 667.000000, Wy: 0.000000},
	"Acircumflex":    CharMetrics{GlyphName: "Acircumflex", Wx: 667.000000, Wy: 0.000000},
	"Adieresis":      CharMetrics{GlyphName: "Adieresis", Wx: 667.000000, Wy: 0.000000},
	"Agrave":         CharMetrics{GlyphName: "Agrave", Wx: 667.000000, Wy: 0.000000},
	"Amacron":        CharMetrics{GlyphName: "Amacron", Wx: 667.000000, Wy: 0.000000},
	"Aogonek":        CharMetrics{GlyphName: "Aogonek", Wx: 667.000000, Wy: 0.000000},
	"Aring":          CharMetrics{GlyphName: "Aring", Wx: 667.000000, Wy: 0.000000},
	"Atilde":         CharMetrics{GlyphName: "Atilde", Wx: 667.000000, Wy: 0.000000},
	"B":              CharMetrics{GlyphName: "B", Wx: 667.000000, Wy: 0.000000},
	"C":              CharMetrics{GlyphName: "C", Wx: 722.000000, Wy: 0.000000},
	"Cacute":         CharMetrics{GlyphName: "Cacute", Wx: 722.000000, Wy: 0.000000},
	"Ccaron":         CharMetrics{GlyphName: "Ccaron", Wx: 722.000000, Wy: 0.000000},
	"Ccedilla":       CharMetrics{GlyphName: "Ccedilla", Wx: 722.000000, Wy: 0.000000},
	"D":              CharMetrics{GlyphName: "D", Wx: 722.000000, Wy: 0.000000},
	"Dcaron":         CharMetrics{GlyphName: "Dcaron", Wx: 722.000000, Wy: 0.000000},
	"Dcroat":         CharMetrics{GlyphName: "Dcroat", Wx: 722.000000, Wy: 0.000000},
	"Delta":          CharMetrics{GlyphName: "Delta", Wx: 612.000000, Wy: 0.000000},
	"E":              CharMetrics{GlyphName: "E", Wx: 667.000000, Wy: 0.000000},
	"Eacute":         CharMetrics{GlyphName: "Eacute", Wx: 667.000000, Wy: 0.000000},
	"Ecaron":         CharMetrics{GlyphName: "Ecaron", Wx: 667.000000, Wy: 0.000000},
	"Ecircumflex":    CharMetrics{GlyphName: "Ecircumflex", Wx: 667.000000, Wy: 0.000000},
	"Edieresis":      CharMetrics{GlyphName: "Edieresis", Wx: 667.000000, Wy: 0.000000},
	"Edotaccent":     CharMetrics{GlyphName: "Edotaccent", Wx: 667.000000, Wy: 0.000000},
	"Egrave":         CharMetrics{GlyphName: "Egrave", Wx: 667.000000, Wy: 0.000000},
	"Emacron":        CharMetrics{GlyphName: "Emacron", Wx: 667.000000, Wy: 0.000000},
	"Eogonek":        CharMetrics{GlyphName: "Eogonek", Wx: 667.000000, Wy: 0.000000},
	"Eth":            CharMetrics{GlyphName: "Eth", Wx: 722.000000, Wy: 0.000000},
	"Euro":           CharMetrics{GlyphName: "Euro", Wx: 556.000000, Wy: 0.000000},
	"F":              CharMetrics{GlyphName: "F", Wx: 611.000000, Wy: 0.000000},
	"G":              CharMetrics{GlyphName: "G", Wx: 778.000000, Wy: 0.000000},
	"Gbreve":         CharMetrics{GlyphName: "Gbreve", Wx: 778.000000, Wy: 0.000000},
	"Gcommaaccent":   CharMetrics{GlyphName: "Gcommaaccent", Wx: 778.000000, Wy: 0.000000},
	"H":              CharMetrics{GlyphName: "H", Wx: 722.000000, Wy: 0.000000},
	"I":              CharMetrics{GlyphName: "I", Wx: 278.000000, Wy: 0.000000},
	"Iacute":         CharMetrics{GlyphName: "Iacute", Wx: 278.000000, Wy: 0.000000},
	"Icircumflex":    CharMetrics{GlyphName: "Icircumflex", Wx: 278.000000, Wy: 0.000000},
	"Idieresis":      CharMetrics{GlyphName: "Idieresis", Wx: 278.000000, Wy: 0.000000},
	"Idotaccent":     CharMetrics{GlyphName: "Idotaccent", Wx: 278.000000, Wy: 0.000000},
	"Igrave":         CharMetrics{GlyphName: "Igrave", Wx: 278.000000, Wy: 0.000000},
	"Imacron":        CharMetrics{GlyphName: "Imacron", Wx: 278.000000, Wy: 0.000000},
	"Iogonek":        CharMetrics{GlyphName: "Iogonek", Wx: 278.000000, Wy: 0.000000},
	"J":              CharMetrics{GlyphName: "J", Wx: 500.000000, Wy: 0.000000},
	"K":              CharMetrics{GlyphName: "K", Wx: 667.000000, Wy: 0.000000},
	"Kcommaaccent":   CharMetrics{GlyphName: "Kcommaaccent", Wx: 667.000000, Wy: 0.000000},
	"L":              CharMetrics{GlyphName: "L", Wx: 556.000000, Wy: 0.000000},
	"Lacute":         CharMetrics{GlyphName: "Lacute", Wx: 556.000000, Wy: 0.000000},
	"Lcaron":         CharMetrics{GlyphName: "Lcaron", Wx: 556.000000, Wy: 0.000000},
	"Lcommaaccent":   CharMetrics{GlyphName: "Lcommaaccent", Wx: 556.000000, Wy: 0.000000},
	"Lslash":         CharMetrics{GlyphName: "Lslash", Wx: 556.000000, Wy: 0.000000},
	"M":              CharMetrics{GlyphName: "M", Wx: 833.000000, Wy: 0.000000},
	"N":              CharMetrics{GlyphName: "N", Wx: 722.000000, Wy: 0.000000},
	"Nacute":         CharMetrics{GlyphName: "Nacute", Wx: 722.000000, Wy: 0.000000},
	"Ncaron":         CharMetrics{GlyphName: "Ncaron", Wx: 722.000000, Wy: 0.000000},
	"Ncommaaccent":   CharMetrics{GlyphName: "Ncommaaccent", Wx: 722.000000, Wy: 0.000000},
	"Ntilde":         CharMetrics{GlyphName: "Ntilde", Wx: 722.000000, Wy: 0.000000},
	"O":              CharMetrics{GlyphName: "O", Wx: 778.000000, Wy: 0.000000},
	"OE":             CharMetrics{GlyphName: "OE", Wx: 1000.000000, Wy: 0.000000},
	"Oacute":         CharMetrics{GlyphName: "Oacute", Wx: 778.000000, Wy: 0.000000},
	"Ocircumflex":    CharMetrics{GlyphName: "Ocircumflex", Wx: 778.000000, Wy: 0.000000},
	"Odieresis":      CharMetrics{GlyphName: "Odieresis", Wx: 778.000000, Wy: 0.000000},
	"Ograve":         CharMetrics{GlyphName: "Ograve", Wx: 778.000000, Wy: 0.000000},
	"Ohungarumlaut":  CharMetrics{GlyphName: "Ohungarumlaut", Wx: 778.000000, Wy: 0.000000},
	"Omacron":        CharMetrics{GlyphName: "Omacron", Wx: 778.000000, Wy: 0.000000},
	"Oslash":         CharMetrics{GlyphName: "Oslash", Wx: 778.000000, Wy: 0.000000},
	"Otilde":         CharMetrics{GlyphName: "Otilde", Wx: 778.000000, Wy: 0.000000},
	"P":              CharMetrics{GlyphName: "P", Wx: 667.000000, Wy: 0.000000},
	"Q":              CharMetrics{GlyphName: "Q", Wx: 778.000000, Wy: 0.000000},
	"R":              CharMetrics{GlyphName: "R", Wx: 722.000000, Wy: 0.000000},
	"Racute":         CharMetrics{GlyphName: "Racute", Wx: 722.000000, Wy: 0.000000},
	"Rcaron":         CharMetrics{GlyphName: "Rcaron", Wx: 722.000000, Wy: 0.000000},
	"Rcommaaccent":   CharMetrics{GlyphName: "Rcommaaccent", Wx: 722.000000, Wy: 0.000000},
	"S":              CharMetrics{GlyphName: "S", Wx: 667.000000, Wy: 0.000000},
	"Sacute":         CharMetrics{GlyphName: "Sacute", Wx: 667.000000, Wy: 0.000000},
	"Scaron":         CharMetrics{GlyphName: "Scaron", Wx: 667.000000, Wy: 0.000000},
	"Scedilla":       CharMetrics{GlyphName: "Scedilla", Wx: 667.000000, Wy: 0.000000},
	"Scommaaccent":   CharMetrics{GlyphName: "Scommaaccent", Wx: 667.000000, Wy: 0.000000},
	"T":              CharMetrics{GlyphName: "T", Wx: 611.000000, Wy: 0.000000},
	"Tcaron":         CharMetrics{GlyphName: "Tcaron", Wx: 611.000000, Wy: 0.000000},
	"Tcommaaccent":   CharMetrics{GlyphName: "Tcommaaccent", Wx: 611.000000, Wy: 0.000000},
	"Thorn":          CharMetrics{GlyphName: "Thorn", Wx: 667.000000, Wy: 0.000000},
	"U":              CharMetrics{GlyphName: "U", Wx: 722.000000, Wy: 0.000000},
	"Uacute":         CharMetrics{GlyphName: "Uacute", Wx: 722.000000, Wy: 0.000000},
	"Ucircumflex":    CharMetrics{GlyphName: "Ucircumflex", Wx: 722.000000, Wy: 0.000000},
	"Udieresis":      CharMetrics{GlyphName: "Udieresis", Wx: 722.000000, Wy: 0.000000},
	"Ugrave":         CharMetrics{GlyphName: "Ugrave", Wx: 722.000000, Wy: 0.000000},
	"Uhungarumlaut":  CharMetrics{GlyphName: "Uhungarumlaut", Wx: 722.000000, Wy: 0.000000},
	"Umacron":        CharMetrics{GlyphName: "Umacron", Wx: 722.000000, Wy: 0.000000},
	"Uogonek":        CharMetrics{GlyphName: "Uogonek", Wx: 722.000000, Wy: 0.000000},
	"Uring":          CharMetrics{GlyphName: "Uring", Wx: 722.000000, Wy: 0.000000},
	"V":              CharMetrics{GlyphName: "V", Wx: 667.000000, Wy: 0.000000},
	"W":              CharMetrics{GlyphName: "W", Wx: 944.000000, Wy: 0.000000},
	"X":              CharMetrics{GlyphName: "X", Wx: 667.000000, Wy: 0.000000},
	"Y":              CharMetrics{GlyphName: "Y", Wx: 667.000000, Wy: 0.000000},
	"Yacute":         CharMetrics{GlyphName: "Yacute", Wx: 667.000000, Wy: 0.000000},
	"Ydieresis":      CharMetrics{GlyphName: "Ydieresis", Wx: 667.000000, Wy: 0.000000},
	"Z":              CharMetrics{GlyphName: "Z", Wx: 611.000000, Wy: 0.000000},
	"Zacute":         CharMetrics{GlyphName: "Zacute", Wx: 611.000000, Wy: 0.000000},
	"Zcaron":         CharMetrics{GlyphName: "Zcaron", Wx: 611.000000, Wy: 0.000000},
	"Zdotaccent":     CharMetrics{GlyphName: "Zdotaccent", Wx: 611.000000, Wy: 0.000000},
	"a":              CharMetrics{GlyphName: "a", Wx: 556.000000, Wy: 0.000000},
	"aacute":         CharMetrics{GlyphName: "aacute", Wx: 556.000000, Wy: 0.000000},
	"abreve":         CharMetrics{GlyphName: "abreve", Wx: 556.000000, Wy: 0.000000},
	"acircumflex":    CharMetrics{GlyphName: "acircumflex", Wx: 556.000000, Wy: 0.000000},
	"acute":          CharMetrics{GlyphName: "acute", Wx: 333.000000, Wy: 0.000000},
	"adieresis":      CharMetrics{GlyphName: "adieresis", Wx: 556.000000, Wy: 0.000000},
	"ae":             CharMetrics{GlyphName: "ae", Wx: 889.000000, Wy: 0.000000},
	"agrave":         CharMetrics{GlyphName: "agrave", Wx: 556.000000, Wy: 0.000000},
	"amacron":        CharMetrics{GlyphName: "amacron", Wx: 556.000000, Wy: 0.000000},
	"ampersand":      CharMetrics{GlyphName: "ampersand", Wx: 667.000000, Wy: 0.000000},
	"aogonek":        CharMetrics{GlyphName: "aogonek", Wx: 556.000000, Wy: 0.000000},
	"aring":          CharMetrics{GlyphName: "aring", Wx: 556.000000, Wy: 0.000000},
	"asciicircum":    CharMetrics{GlyphName: "asciicircum", Wx: 469.000000, Wy: 0.000000},
	"asciitilde":     CharMetrics{GlyphName: "asciitilde", Wx: 584.000000, Wy: 0.000000},
	"asterisk":       CharMetrics{GlyphName: "asterisk", Wx: 389.000000, Wy: 0.000000},
	"at":             CharMetrics{GlyphName: "at", Wx: 1015.000000, Wy: 0.000000},
	"atilde":         CharMetrics{GlyphName: "atilde", Wx: 556.000000, Wy: 0.000000},
	"b":              CharMetrics{GlyphName: "b", Wx: 556.000000, Wy: 0.000000},
	"backslash":      CharMetrics{GlyphName: "backslash", Wx: 278.000000, Wy: 0.000000},
	"bar":            CharMetrics{GlyphName: "bar", Wx: 260.000000, Wy: 0.000000},
	"braceleft":      CharMetrics{GlyphName: "braceleft", Wx: 334.000000, Wy: 0.000000},
	"braceright":     CharMetrics{GlyphName: "braceright", Wx: 334.000000, Wy: 0.000000},
	"bracketleft":    CharMetrics{GlyphName: "bracketleft", Wx: 278.000000, Wy: 0.000000},
	"bracketright":   CharMetrics{GlyphName: "bracketright", Wx: 278.000000, Wy: 0.000000},
	"breve":          CharMetrics{GlyphName: "breve", Wx: 333.000000, Wy: 0.000000},
	"brokenbar":      CharMetrics{GlyphName: "brokenbar", Wx: 260.000000, Wy: 0.000000},
	"bullet":         CharMetrics{GlyphName: "bullet", Wx: 350.000000, Wy: 0.000000},
	"c":              CharMetrics{GlyphName: "c", Wx: 500.000000, Wy: 0.000000},
	"cacute":         CharMetrics{GlyphName: "cacute", Wx: 500.000000, Wy: 0.000000},
	"caron":          CharMetrics{GlyphName: "caron", Wx: 333.000000, Wy: 0.000000},
	"ccaron":         CharMetrics{GlyphName: "ccaron", Wx: 500.000000, Wy: 0.000000},
	"ccedilla":       CharMetrics{GlyphName: "ccedilla", Wx: 500.000000, Wy: 0.000000},
	"cedilla":        CharMetrics{GlyphName: "cedilla", Wx: 333.000000, Wy: 0.000000},
	"cent":           CharMetrics{GlyphName: "cent", Wx: 556.000000, Wy: 0.000000},
	"circumflex":     CharMetrics{GlyphName: "circumflex", Wx: 333.000000, Wy: 0.000000},
	"colon":          CharMetrics{GlyphName: "colon", Wx: 278.000000, Wy: 0.000000},
	"comma":          CharMetrics{GlyphName: "comma", Wx: 278.000000, Wy: 0.000000},
	"commaaccent":    CharMetrics{GlyphName: "commaaccent", Wx: 250.000000, Wy: 0.000000},
	"copyright":      CharMetrics{GlyphName: "copyright", Wx: 737.000000, Wy: 0.000000},
	"currency":       CharMetrics{GlyphName: "currency", Wx: 556.000000, Wy: 0.000000},
	"d":              CharMetrics{GlyphName: "d", Wx: 556.000000, Wy: 0.000000},
	"dagger":         CharMetrics{GlyphName: "dagger", Wx: 556.000000, Wy: 0.000000},
	"daggerdbl":      CharMetrics{GlyphName: "daggerdbl", Wx: 556.000000, Wy: 0.000000},
	"dcaron":         CharMetrics{GlyphName: "dcaron", Wx: 643.000000, Wy: 0.000000},
	"dcroat":         CharMetrics{GlyphName: "dcroat", Wx: 556.000000, Wy: 0.000000},
	"degree":         CharMetrics{GlyphName: "degree", Wx: 400.000000, Wy: 0.000000},
	"dieresis":       CharMetrics{GlyphName: "dieresis", Wx: 333.000000, Wy: 0.000000},
	"divide":         CharMetrics{GlyphName: "divide", Wx: 584.000000, Wy: 0.000000},
	"dollar":         CharMetrics{GlyphName: "dollar", Wx: 556.000000, Wy: 0.000000},
	"dotaccent":      CharMetrics{GlyphName: "dotaccent", Wx: 333.000000, Wy: 0.000000},
	"dotlessi":       CharMetrics{GlyphName: "dotlessi", Wx: 278.000000, Wy: 0.000000},
	"e":              CharMetrics{GlyphName: "e", Wx: 556.000000, Wy: 0.000000},
	"eacute":         CharMetrics{GlyphName: "eacute", Wx: 556.000000, Wy: 0.000000},
	"ecaron":         CharMetrics{GlyphName: "ecaron", Wx: 556.000000, Wy: 0.000000},
	"ecircumflex":    CharMetrics{GlyphName: "ecircumflex", Wx: 556.000000, Wy: 0.000000},
	"edieresis":      CharMetrics{GlyphName: "edieresis", Wx: 556.000000, Wy: 0.000000},
	"edotaccent":     CharMetrics{GlyphName: "edotaccent", Wx: 556.000000, Wy: 0.000000},
	"egrave":         CharMetrics{GlyphName: "egrave", Wx: 556.000000, Wy: 0.000000},
	"eight":          CharMetrics{GlyphName: "eight", Wx: 556.000000, Wy: 0.000000},
	"ellipsis":       CharMetrics{GlyphName: "ellipsis", Wx: 1000.000000, Wy: 0.000000},
	"emacron":        CharMetrics{GlyphName: "emacron", Wx: 556.000000, Wy: 0.000000},
	"emdash":         CharMetrics{GlyphName: "emdash", Wx: 1000.000000, Wy: 0.000000},
	"endash":         CharMetrics{GlyphName: "endash", Wx: 556.000000, Wy: 0.000000},
	"eogonek":        CharMetrics{GlyphName: "eogonek", Wx: 556.000000, Wy: 0.000000},
	"equal":          CharMetrics{GlyphName: "equal", Wx: 584.000000, Wy: 0.000000},
	"eth":            CharMetrics{GlyphName: "eth", Wx: 556.000000, Wy: 0.000000},
	"exclam":         CharMetrics{GlyphName: "exclam", Wx: 278.000000, Wy: 0.000000},
	"exclamdown":     CharMetrics{GlyphName: "exclamdown", Wx: 333.000000, Wy: 0.000000},
	"f":              CharMetrics{GlyphName: "f", Wx: 278.000000, Wy: 0.000000},
	"fi":             CharMetrics{GlyphName: "fi", Wx: 500.000000, Wy: 0.000000},
	"five":           CharMetrics{GlyphName: "five", Wx: 556.000000, Wy: 0.000000},
	"fl":             CharMetrics{GlyphName: "fl", Wx: 500.000000, Wy: 0.000000},
	"florin":         CharMetrics{GlyphName: "florin", Wx: 556.000000, Wy: 0.000000},
	"four":           CharMetrics{GlyphName: "four", Wx: 556.000000, Wy: 0.000000},
	"fraction":       CharMetrics{GlyphName: "fraction", Wx: 167.000000, Wy: 0.000000},
	"g":              CharMetrics{GlyphName: "g", Wx: 556.000000, Wy: 0.000000},
	"gbreve":         CharMetrics{GlyphName: "gbreve", Wx: 556.000000, Wy: 0.000000},
	"gcommaaccent":   CharMetrics{GlyphName: "gcommaaccent", Wx: 556.000000, Wy: 0.000000},
	"germandbls":     CharMetrics{GlyphName: "germandbls", Wx: 611.000000, Wy: 0.000000},
	"grave":          CharMetrics{GlyphName: "grave", Wx: 333.000000, Wy: 0.000000},
	"greater":        CharMetrics{GlyphName: "greater", Wx: 584.000000, Wy: 0.000000},
	"greaterequal":   CharMetrics{GlyphName: "greaterequal", Wx: 549.000000, Wy: 0.000000},
	"guillemotleft":  CharMetrics{GlyphName: "guillemotleft", Wx: 556.000000, Wy: 0.000000},
	"guillemotright": CharMetrics{GlyphName: "guillemotright", Wx: 556.000000, Wy: 0.000000},
	"guilsinglleft":  CharMetrics{GlyphName: "guilsinglleft", Wx: 333.000000, Wy: 0.000000},
	"guilsinglright": CharMetrics{GlyphName: "guilsinglright", Wx: 333.000000, Wy: 0.000000},
	"h":              CharMetrics{GlyphName: "h", Wx: 556.000000, Wy: 0.000000},
	"hungarumlaut":   CharMetrics{GlyphName: "hungarumlaut", Wx: 333.000000, Wy: 0.000000},
	"hyphen":         CharMetrics{GlyphName: "hyphen", Wx: 333.000000, Wy: 0.000000},
	"i":              CharMetrics{GlyphName: "i", Wx: 222.000000, Wy: 0.000000},
	"iacute":         CharMetrics{GlyphName: "iacute", Wx: 278.000000, Wy: 0.000000},
	"icircumflex":    CharMetrics{GlyphName: "icircumflex", Wx: 278.000000, Wy: 0.000000},
	"idieresis":      CharMetrics{GlyphName: "idieresis", Wx: 278.000000, Wy: 0.000000},
	"igrave":         CharMetrics{GlyphName: "igrave", Wx: 278.000000, Wy: 0.000000},
	"imacron":        CharMetrics{GlyphName: "imacron", Wx: 278.000000, Wy: 0.000000},
	"iogonek":        CharMetrics{GlyphName: "iogonek", Wx: 222.000000, Wy: 0.000000},
	"j":              CharMetrics{GlyphName: "j", Wx: 222.000000, Wy: 0.000000},
	"k":              CharMetrics{GlyphName: "k", Wx: 500.000000, Wy: 0.000000},
	"kcommaaccent":   CharMetrics{GlyphName: "kcommaaccent", Wx: 500.000000, Wy: 0.000000},
	"l":              CharMetrics{GlyphName: "l", Wx: 222.000000, Wy: 0.000000},
	"lacute":         CharMetrics{GlyphName: "lacute", Wx: 222.000000, Wy: 0.000000},
	"lcaron":         CharMetrics{GlyphName: "lcaron", Wx: 299.000000, Wy: 0.000000},
	"lcommaaccent":   CharMetrics{GlyphName: "lcommaaccent", Wx: 222.000000, Wy: 0.000000},
	"less":           CharMetrics{GlyphName: "less", Wx: 584.000000, Wy: 0.000000},
	"lessequal":      CharMetrics{GlyphName: "lessequal", Wx: 549.000000, Wy: 0.000000},
	"logicalnot":     CharMetrics{GlyphName: "logicalnot", Wx: 584.000000, Wy: 0.000000},
	"lozenge":        CharMetrics{GlyphName: "lozenge", Wx: 471.000000, Wy: 0.000000},
	"lslash":         CharMetrics{GlyphName: "lslash", Wx: 222.000000, Wy: 0.000000},
	"m":              CharMetrics{GlyphName: "m", Wx: 833.000000, Wy: 0.000000},
	"macron":         CharMetrics{GlyphName: "macron", Wx: 333.000000, Wy: 0.000000},
	"minus":          CharMetrics{GlyphName: "minus", Wx: 584.000000, Wy: 0.000000},
	"mu":             CharMetrics{GlyphName: "mu", Wx: 556.000000, Wy: 0.000000},
	"multiply":       CharMetrics{GlyphName: "multiply", Wx: 584.000000, Wy: 0.000000},
	"n":              CharMetrics{GlyphName: "n", Wx: 556.000000, Wy: 0.000000},
	"nacute":         CharMetrics{GlyphName: "nacute", Wx: 556.000000, Wy: 0.000000},
	"ncaron":         CharMetrics{GlyphName: "ncaron", Wx: 556.000000, Wy: 0.000000},
	"ncommaaccent":   CharMetrics{GlyphName: "ncommaaccent", Wx: 556.000000, Wy: 0.000000},
	"nine":           CharMetrics{GlyphName: "nine", Wx: 556.000000, Wy: 0.000000},
	"notequal":       CharMetrics{GlyphName: "notequal", Wx: 549.000000, Wy: 0.000000},
	"ntilde":         CharMetrics{GlyphName: "ntilde", Wx: 556.000000, Wy: 0.000000},
	"numbersign":     CharMetrics{GlyphName: "numbersign", Wx: 556.000000, Wy: 0.000000},
	"o":              CharMetrics{GlyphName: "o", Wx: 556.000000, Wy: 0.000000},
	"oacute":         CharMetrics{GlyphName: "oacute", Wx: 556.000000, Wy: 0.000000},
	"ocircumflex":    CharMetrics{GlyphName: "ocircumflex", Wx: 556.000000, Wy: 0.000000},
	"odieresis":      CharMetrics{GlyphName: "odieresis", Wx: 556.000000, Wy: 0.000000},
	"oe":             CharMetrics{GlyphName: "oe", Wx: 944.000000, Wy: 0.000000},
	"ogonek":         CharMetrics{GlyphName: "ogonek", Wx: 333.000000, Wy: 0.000000},
	"ograve":         CharMetrics{GlyphName: "ograve", Wx: 556.000000, Wy: 0.000000},
	"ohungarumlaut":  CharMetrics{GlyphName: "ohungarumlaut", Wx: 556.000000, Wy: 0.000000},
	"omacron":        CharMetrics{GlyphName: "omacron", Wx: 556.000000, Wy: 0.000000},
	"one":            CharMetrics{GlyphName: "one", Wx: 556.000000, Wy: 0.000000},
	"onehalf":        CharMetrics{GlyphName: "onehalf", Wx: 834.000000, Wy: 0.000000},
	"onequarter":     CharMetrics{GlyphName: "onequarter", Wx: 834.000000, Wy: 0.000000},
	"onesuperior":    CharMetrics{GlyphName: "onesuperior", Wx: 333.000000, Wy: 0.000000},
	"ordfeminine":    CharMetrics{GlyphName: "ordfeminine", Wx: 370.000000, Wy: 0.000000},
	"ordmasculine":   CharMetrics{GlyphName: "ordmasculine", Wx: 365.000000, Wy: 0.000000},
	"oslash":         CharMetrics{GlyphName: "oslash", Wx: 611.000000, Wy: 0.000000},
	"otilde":         CharMetrics{GlyphName: "otilde", Wx: 556.000000, Wy: 0.000000},
	"p":              CharMetrics{GlyphName: "p", Wx: 556.000000, Wy: 0.000000},
	"paragraph":      CharMetrics{GlyphName: "paragraph", Wx: 537.000000, Wy: 0.000000},
	"parenleft":      CharMetrics{GlyphName: "parenleft", Wx: 333.000000, Wy: 0.000000},
	"parenright":     CharMetrics{GlyphName: "parenright", Wx: 333.000000, Wy: 0.000000},
	"partialdiff":    CharMetrics{GlyphName: "partialdiff", Wx: 476.000000, Wy: 0.000000},
	"percent":        CharMetrics{GlyphName: "percent", Wx: 889.000000, Wy: 0.000000},
	"period":         CharMetrics{GlyphName: "period", Wx: 278.000000, Wy: 0.000000},
	"periodcentered": CharMetrics{GlyphName: "periodcentered", Wx: 278.000000, Wy: 0.000000},
	"perthousand":    CharMetrics{GlyphName: "perthousand", Wx: 1000.000000, Wy: 0.000000},
	"plus":           CharMetrics{GlyphName: "plus", Wx: 584.000000, Wy: 0.000000},
	"plusminus":      CharMetrics{GlyphName: "plusminus", Wx: 584.000000, Wy: 0.000000},
	"q":              CharMetrics{GlyphName: "q", Wx: 556.000000, Wy: 0.000000},
	"question":       CharMetrics{GlyphName: "question", Wx: 556.000000, Wy: 0.000000},
	"questiondown":   CharMetrics{GlyphName: "questiondown", Wx: 611.000000, Wy: 0.000000},
	"quotedbl":       CharMetrics{GlyphName: "quotedbl", Wx: 355.000000, Wy: 0.000000},
	"quotedblbase":   CharMetrics{GlyphName: "quotedblbase", Wx: 333.000000, Wy: 0.000000},
	"quotedblleft":   CharMetrics{GlyphName: "quotedblleft", Wx: 333.000000, Wy: 0.000000},
	"quotedblright":  CharMetrics{GlyphName: "quotedblright", Wx: 333.000000, Wy: 0.000000},
	"quoteleft":      CharMetrics{GlyphName: "quoteleft", Wx: 222.000000, Wy: 0.000000},
	"quoteright":     CharMetrics{GlyphName: "quoteright", Wx: 222.000000, Wy: 0.000000},
	"quotesinglbase": CharMetrics{GlyphName: "quotesinglbase", Wx: 222.000000, Wy: 0.000000},
	"quotesingle":    CharMetrics{GlyphName: "quotesingle", Wx: 191.000000, Wy: 0.000000},
	"r":              CharMetrics{GlyphName: "r", Wx: 333.000000, Wy: 0.000000},
	"racute":         CharMetrics{GlyphName: "racute", Wx: 333.000000, Wy: 0.000000},
	"radical":        CharMetrics{GlyphName: "radical", Wx: 453.000000, Wy: 0.000000},
	"rcaron":         CharMetrics{GlyphName: "rcaron", Wx: 333.000000, Wy: 0.000000},
	"rcommaaccent":   CharMetrics{GlyphName: "rcommaaccent", Wx: 333.000000, Wy: 0.000000},
	"registered":     CharMetrics{GlyphName: "registered", Wx: 737.000000, Wy: 0.000000},
	"ring":           CharMetrics{GlyphName: "ring", Wx: 333.000000, Wy: 0.000000},
	"s":              CharMetrics{GlyphName: "s", Wx: 500.000000, Wy: 0.000000},
	"sacute":         CharMetrics{GlyphName: "sacute", Wx: 500.000000, Wy: 0.000000},
	"scaron":         CharMetrics{GlyphName: "scaron", Wx: 500.000000, Wy: 0.000000},
	"scedilla":       CharMetrics{GlyphName: "scedilla", Wx: 500.000000, Wy: 0.000000},
	"scommaaccent":   CharMetrics{GlyphName: "scommaaccent", Wx: 500.000000, Wy: 0.000000},
	"section":        CharMetrics{GlyphName: "section", Wx: 556.000000, Wy: 0.000000},
	"semicolon":      CharMetrics{GlyphName: "semicolon", Wx: 278.000000, Wy: 0.000000},
	"seven":          CharMetrics{GlyphName: "seven", Wx: 556.000000, Wy: 0.000000},
	"six":            CharMetrics{GlyphName: "six", Wx: 556.000000, Wy: 0.000000},
	"slash":          CharMetrics{GlyphName: "slash", Wx: 278.000000, Wy: 0.000000},
	"space":          CharMetrics{GlyphName: "space", Wx: 278.000000, Wy: 0.000000},
	"sterling":       CharMetrics{GlyphName: "sterling", Wx: 556.000000, Wy: 0.000000},
	"summation":      CharMetrics{GlyphName: "summation", Wx: 600.000000, Wy: 0.000000},
	"t":              CharMetrics{GlyphName: "t", Wx: 278.000000, Wy: 0.000000},
	"tcaron":         CharMetrics{GlyphName: "tcaron", Wx: 317.000000, Wy: 0.000000},
	"tcommaaccent":   CharMetrics{GlyphName: "tcommaaccent", Wx: 278.000000, Wy: 0.000000},
	"thorn":          CharMetrics{GlyphName: "thorn", Wx: 556.000000, Wy: 0.000000},
	"three":          CharMetrics{GlyphName: "three", Wx: 556.000000, Wy: 0.000000},
	"threequarters":  CharMetrics{GlyphName: "threequarters", Wx: 834.000000, Wy: 0.000000},
	"threesuperior":  CharMetrics{GlyphName: "threesuperior", Wx: 333.000000, Wy: 0.000000},
	"tilde":          CharMetrics{GlyphName: "tilde", Wx: 333.000000, Wy: 0.000000},
	"trademark":      CharMetrics{GlyphName: "trademark", Wx: 1000.000000, Wy: 0.000000},
	"two":            CharMetrics{GlyphName: "two", Wx: 556.000000, Wy: 0.000000},
	"twosuperior":    CharMetrics{GlyphName: "twosuperior", Wx: 333.000000, Wy: 0.000000},
	"u":              CharMetrics{GlyphName: "u", Wx: 556.000000, Wy: 0.000000},
	"uacute":         CharMetrics{GlyphName: "uacute", Wx: 556.000000, Wy: 0.000000},
	"ucircumflex":    CharMetrics{GlyphName: "ucircumflex", Wx: 556.000000, Wy: 0.000000},
	"udieresis":      CharMetrics{GlyphName: "udieresis", Wx: 556.000000, Wy: 0.000000},
	"ugrave":         CharMetrics{GlyphName: "ugrave", Wx: 556.000000, Wy: 0.000000},
	"uhungarumlaut":  CharMetrics{GlyphName: "uhungarumlaut", Wx: 556.000000, Wy: 0.000000},
	"umacron":        CharMetrics{GlyphName: "umacron", Wx: 556.000000, Wy: 0.000000},
	"underscore":     CharMetrics{GlyphName: "underscore", Wx: 556.000000, Wy: 0.000000},
	"uogonek":        CharMetrics{GlyphName: "uogonek", Wx: 556.000000, Wy: 0.000000},
	"uring":          CharMetrics{GlyphName: "uring", Wx: 556.000000, Wy: 0.000000},
	"v":              CharMetrics{GlyphName: "v", Wx: 500.000000, Wy: 0.000000},
	"w":              CharMetrics{GlyphName: "w", Wx: 722.000000, Wy: 0.000000},
	"x":              CharMetrics{GlyphName: "x", Wx: 500.000000, Wy: 0.000000},
	"y":              CharMetrics{GlyphName: "y", Wx: 500.000000, Wy: 0.000000},
	"yacute":         CharMetrics{GlyphName: "yacute", Wx: 500.000000, Wy: 0.000000},
	"ydieresis":      CharMetrics{GlyphName: "ydieresis", Wx: 500.000000, Wy: 0.000000},
	"yen":            CharMetrics{GlyphName: "yen", Wx: 556.000000, Wy: 0.000000},
	"z":              CharMetrics{GlyphName: "z", Wx: 500.000000, Wy: 0.000000},
	"zacute":         CharMetrics{GlyphName: "zacute", Wx: 500.000000, Wy: 0.000000},
	"zcaron":         CharMetrics{GlyphName: "zcaron", Wx: 500.000000, Wy: 0.000000},
	"zdotaccent":     CharMetrics{GlyphName: "zdotaccent", Wx: 500.000000, Wy: 0.000000},
	"zero":           CharMetrics{GlyphName: "zero", Wx: 556.000000, Wy: 0.000000},
}
