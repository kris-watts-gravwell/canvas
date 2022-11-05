package canvas

import "image/color"

var namedColors = map[string]color.RGBA{
	"aliceblue":            color.RGBA{R: 0xF0, G: 0xF8, B: 0xFF, A: 0xFF},
	"antiquewhite":         color.RGBA{R: 0xFA, G: 0xEB, B: 0xD7, A: 0xFF},
	"aqua":                 color.RGBA{R: 0x00, G: 0xFF, B: 0xFF, A: 0xFF},
	"aquamarine":           color.RGBA{R: 0x7F, G: 0xFF, B: 0xD4, A: 0xFF},
	"azure":                color.RGBA{R: 0xF0, G: 0xFF, B: 0xFF, A: 0xFF},
	"beige":                color.RGBA{R: 0xF5, G: 0xF5, B: 0xDC, A: 0xFF},
	"bisque":               color.RGBA{R: 0xFF, G: 0xE4, B: 0xC4, A: 0xFF},
	"black":                color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
	"blanchedalmond":       color.RGBA{R: 0xFF, G: 0xEB, B: 0xCD, A: 0xFF},
	"blue":                 color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF},
	"blueviolet":           color.RGBA{R: 0x8A, G: 0x2B, B: 0xE2, A: 0xFF},
	"brown":                color.RGBA{R: 0xA5, G: 0x2A, B: 0x2A, A: 0xFF},
	"burlywood":            color.RGBA{R: 0xDE, G: 0xB8, B: 0x87, A: 0xFF},
	"cadetblue":            color.RGBA{R: 0x5F, G: 0x9E, B: 0xA0, A: 0xFF},
	"chocolate":            color.RGBA{R: 0xD2, G: 0x69, B: 0x1E, A: 0xFF},
	"coral":                color.RGBA{R: 0xFF, G: 0x7F, B: 0x50, A: 0xFF},
	"cornflowerblue":       color.RGBA{R: 0x64, G: 0x95, B: 0xED, A: 0xFF},
	"cornsilk":             color.RGBA{R: 0xFF, G: 0xF8, B: 0xDC, A: 0xFF},
	"crimson":              color.RGBA{R: 0xDC, G: 0x14, B: 0x3C, A: 0xFF},
	"cyan":                 color.RGBA{R: 0x00, G: 0xFF, B: 0xFF, A: 0xFF},
	"darkblue":             color.RGBA{R: 0x00, G: 0x00, B: 0x8B, A: 0xFF},
	"darkcyan":             color.RGBA{R: 0x00, G: 0x8B, B: 0x8B, A: 0xFF},
	"darkgoldenrod":        color.RGBA{R: 0xB8, G: 0x86, B: 0x0B, A: 0xFF},
	"darkgray":             color.RGBA{R: 0xA9, G: 0xA9, B: 0xA9, A: 0xFF},
	"darkgreen":            color.RGBA{R: 0x00, G: 0x64, B: 0x00, A: 0xFF},
	"darkkhaki":            color.RGBA{R: 0xBD, G: 0xB7, B: 0x6B, A: 0xFF},
	"darkmagenta":          color.RGBA{R: 0x8B, G: 0x00, B: 0x8B, A: 0xFF},
	"darkolivegreen":       color.RGBA{R: 0x55, G: 0x6B, B: 0x2F, A: 0xFF},
	"darkorange":           color.RGBA{R: 0xFF, G: 0x8C, B: 0x00, A: 0xFF},
	"darkorchid":           color.RGBA{R: 0x99, G: 0x32, B: 0xCC, A: 0xFF},
	"darkred":              color.RGBA{R: 0x8B, G: 0x00, B: 0x00, A: 0xFF},
	"darksalmon":           color.RGBA{R: 0xE9, G: 0x96, B: 0x7A, A: 0xFF},
	"darkseagreen":         color.RGBA{R: 0x8F, G: 0xBC, B: 0x8F, A: 0xFF},
	"darkslateblue":        color.RGBA{R: 0x48, G: 0x3D, B: 0x8B, A: 0xFF},
	"darkslategray":        color.RGBA{R: 0x2F, G: 0x4F, B: 0x4F, A: 0xFF},
	"darkturquoise":        color.RGBA{R: 0x00, G: 0xCE, B: 0xD1, A: 0xFF},
	"darkviolet":           color.RGBA{R: 0x94, G: 0x00, B: 0xD3, A: 0xFF},
	"deeppink":             color.RGBA{R: 0xFF, G: 0x14, B: 0x93, A: 0xFF},
	"deepskyblue":          color.RGBA{R: 0x00, G: 0xBF, B: 0xFF, A: 0xFF},
	"dimgray":              color.RGBA{R: 0x69, G: 0x69, B: 0x69, A: 0xFF},
	"dodgerblue":           color.RGBA{R: 0x1E, G: 0x90, B: 0xFF, A: 0xFF},
	"firebrick":            color.RGBA{R: 0xB2, G: 0x22, B: 0x22, A: 0xFF},
	"floralwhite":          color.RGBA{R: 0xFF, G: 0xFA, B: 0xF0, A: 0xFF},
	"forestgreen":          color.RGBA{R: 0x22, G: 0x8B, B: 0x22, A: 0xFF},
	"fuchsia":              color.RGBA{R: 0xFF, G: 0x00, B: 0xFF, A: 0xFF},
	"gainsboro":            color.RGBA{R: 0xDC, G: 0xDC, B: 0xDC, A: 0xFF},
	"ghostwhite":           color.RGBA{R: 0xF8, G: 0xF8, B: 0xFF, A: 0xFF},
	"gold":                 color.RGBA{R: 0xFF, G: 0xD7, B: 0x00, A: 0xFF},
	"goldenrod":            color.RGBA{R: 0xDA, G: 0xA5, B: 0x20, A: 0xFF},
	"gray":                 color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xFF},
	"green":                color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xFF},
	"greenyellow":          color.RGBA{R: 0xAD, G: 0xFF, B: 0x2F, A: 0xFF},
	"honeydew":             color.RGBA{R: 0xF0, G: 0xFF, B: 0xF0, A: 0xFF},
	"hotpink":              color.RGBA{R: 0xFF, G: 0x69, B: 0xB4, A: 0xFF},
	"indianred":            color.RGBA{R: 0xCD, G: 0x5C, B: 0x5C, A: 0xFF},
	"indigo":               color.RGBA{R: 0x4B, G: 0x00, B: 0x82, A: 0xFF},
	"ivory":                color.RGBA{R: 0xFF, G: 0xFF, B: 0xF0, A: 0xFF},
	"khaki":                color.RGBA{R: 0xF0, G: 0xE6, B: 0x8C, A: 0xFF},
	"lavender":             color.RGBA{R: 0xE6, G: 0xE6, B: 0xFA, A: 0xFF},
	"lavenderblush":        color.RGBA{R: 0xFF, G: 0xF0, B: 0xF5, A: 0xFF},
	"lawngreen":            color.RGBA{R: 0x7C, G: 0xFC, B: 0x00, A: 0xFF},
	"lemonchiffon":         color.RGBA{R: 0xFF, G: 0xFA, B: 0xCD, A: 0xFF},
	"lightblue":            color.RGBA{R: 0xAD, G: 0xD8, B: 0xE6, A: 0xFF},
	"lightcoral":           color.RGBA{R: 0xF0, G: 0x80, B: 0x80, A: 0xFF},
	"lightcyan":            color.RGBA{R: 0xE0, G: 0xFF, B: 0xFF, A: 0xFF},
	"lightgoldenrodyellow": color.RGBA{R: 0xFA, G: 0xFA, B: 0xD2, A: 0xFF},
	"lightgray":            color.RGBA{R: 0xD3, G: 0xD3, B: 0xD3, A: 0xFF},
	"lightgreen":           color.RGBA{R: 0x90, G: 0xEE, B: 0x90, A: 0xFF},
	"lightpink":            color.RGBA{R: 0xFF, G: 0xB6, B: 0xC1, A: 0xFF},
	"lightsalmon":          color.RGBA{R: 0xFF, G: 0xA0, B: 0x7A, A: 0xFF},
	"lightseagreen":        color.RGBA{R: 0x20, G: 0xB2, B: 0xAA, A: 0xFF},
	"lightskyblue":         color.RGBA{R: 0x87, G: 0xCE, B: 0xFA, A: 0xFF},
	"lightslategray":       color.RGBA{R: 0x77, G: 0x88, B: 0x99, A: 0xFF},
	"lightsteelblue":       color.RGBA{R: 0xB0, G: 0xC4, B: 0xDE, A: 0xFF},
	"lightyellow":          color.RGBA{R: 0xFF, G: 0xFF, B: 0xE0, A: 0xFF},
	"lime":                 color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF},
	"limegreen":            color.RGBA{R: 0x32, G: 0xCD, B: 0x32, A: 0xFF},
	"linen":                color.RGBA{R: 0xFA, G: 0xF0, B: 0xE6, A: 0xFF},
	"magenta":              color.RGBA{R: 0xFF, G: 0x00, B: 0xFF, A: 0xFF},
	"maroon":               color.RGBA{R: 0x80, G: 0x00, B: 0x00, A: 0xFF},
	"mediumaquamarine":     color.RGBA{R: 0x66, G: 0xCD, B: 0xAA, A: 0xFF},
	"mediumblue":           color.RGBA{R: 0x00, G: 0x00, B: 0xCD, A: 0xFF},
	"mediumorchid":         color.RGBA{R: 0xBA, G: 0x55, B: 0xD3, A: 0xFF},
	"mediumpurple":         color.RGBA{R: 0x93, G: 0x70, B: 0xDB, A: 0xFF},
	"mediumseagreen":       color.RGBA{R: 0x3C, G: 0xB3, B: 0x71, A: 0xFF},
	"mediumslateblue":      color.RGBA{R: 0x7B, G: 0x68, B: 0xEE, A: 0xFF},
	"mediumspringgreen":    color.RGBA{R: 0x00, G: 0xFA, B: 0x9A, A: 0xFF},
	"mediumturquoise":      color.RGBA{R: 0x48, G: 0xD1, B: 0xCC, A: 0xFF},
	"mediumvioletred":      color.RGBA{R: 0xC7, G: 0x15, B: 0x85, A: 0xFF},
	"midnightblue":         color.RGBA{R: 0x19, G: 0x19, B: 0x70, A: 0xFF},
	"mintcream":            color.RGBA{R: 0xF5, G: 0xFF, B: 0xFA, A: 0xFF},
	"mistyrose":            color.RGBA{R: 0xFF, G: 0xE4, B: 0xE1, A: 0xFF},
	"moccasin":             color.RGBA{R: 0xFF, G: 0xE4, B: 0xB5, A: 0xFF},
	"navajowhite":          color.RGBA{R: 0xFF, G: 0xDE, B: 0xAD, A: 0xFF},
	"navy":                 color.RGBA{R: 0x00, G: 0x00, B: 0x80, A: 0xFF},
	"oldlace":              color.RGBA{R: 0xFD, G: 0xF5, B: 0xE6, A: 0xFF},
	"olive":                color.RGBA{R: 0x80, G: 0x80, B: 0x00, A: 0xFF},
	"olivedrab":            color.RGBA{R: 0x6B, G: 0x8E, B: 0x23, A: 0xFF},
	"orange":               color.RGBA{R: 0xFF, G: 0xA5, B: 0x00, A: 0xFF},
	"orangered":            color.RGBA{R: 0xFF, G: 0x45, B: 0x00, A: 0xFF},
	"orchid":               color.RGBA{R: 0xDA, G: 0x70, B: 0xD6, A: 0xFF},
	"palegoldenrod":        color.RGBA{R: 0xEE, G: 0xE8, B: 0xAA, A: 0xFF},
	"palegreen":            color.RGBA{R: 0x98, G: 0xFB, B: 0x98, A: 0xFF},
	"paleturquoise":        color.RGBA{R: 0xAF, G: 0xEE, B: 0xEE, A: 0xFF},
	"palevioletred":        color.RGBA{R: 0xDB, G: 0x70, B: 0x93, A: 0xFF},
	"papayawhip":           color.RGBA{R: 0xFF, G: 0xEF, B: 0xD5, A: 0xFF},
	"peachpuff":            color.RGBA{R: 0xFF, G: 0xDA, B: 0xB9, A: 0xFF},
	"peru":                 color.RGBA{R: 0xCD, G: 0x85, B: 0x3F, A: 0xFF},
	"pink":                 color.RGBA{R: 0xFF, G: 0xC0, B: 0xCB, A: 0xFF},
	"plum":                 color.RGBA{R: 0xDD, G: 0xA0, B: 0xDD, A: 0xFF},
	"powderblue":           color.RGBA{R: 0xB0, G: 0xE0, B: 0xE6, A: 0xFF},
	"purple":               color.RGBA{R: 0x80, G: 0x00, B: 0x80, A: 0xFF},
	"red":                  color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF},
	"rosybrown":            color.RGBA{R: 0xBC, G: 0x8F, B: 0x8F, A: 0xFF},
	"royalblue":            color.RGBA{R: 0x41, G: 0x69, B: 0xE1, A: 0xFF},
	"saddlebrown":          color.RGBA{R: 0x8B, G: 0x45, B: 0x13, A: 0xFF},
	"salmon":               color.RGBA{R: 0xFA, G: 0x80, B: 0x72, A: 0xFF},
	"sandybrown":           color.RGBA{R: 0xF4, G: 0xA4, B: 0x60, A: 0xFF},
	"seagreen":             color.RGBA{R: 0x2E, G: 0x8B, B: 0x57, A: 0xFF},
	"seashell":             color.RGBA{R: 0xFF, G: 0xF5, B: 0xEE, A: 0xFF},
	"sienna":               color.RGBA{R: 0xA0, G: 0x52, B: 0x2D, A: 0xFF},
	"silver":               color.RGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF},
	"skyblue":              color.RGBA{R: 0x87, G: 0xCE, B: 0xEB, A: 0xFF},
	"slateblue":            color.RGBA{R: 0x6A, G: 0x5A, B: 0xCD, A: 0xFF},
	"slategray":            color.RGBA{R: 0x70, G: 0x80, B: 0x90, A: 0xFF},
	"snow":                 color.RGBA{R: 0xFF, G: 0xFA, B: 0xFA, A: 0xFF},
	"springgreen":          color.RGBA{R: 0x00, G: 0xFF, B: 0x7F, A: 0xFF},
	"steelblue":            color.RGBA{R: 0x46, G: 0x82, B: 0xB4, A: 0xFF},
	"tan":                  color.RGBA{R: 0xD2, G: 0xB4, B: 0x8C, A: 0xFF},
	"teal":                 color.RGBA{R: 0x00, G: 0x80, B: 0x80, A: 0xFF},
	"thistle":              color.RGBA{R: 0xD8, G: 0xBF, B: 0xD8, A: 0xFF},
	"tomato":               color.RGBA{R: 0xFF, G: 0x63, B: 0x47, A: 0xFF},
	"turquoise":            color.RGBA{R: 0x40, G: 0xE0, B: 0xD0, A: 0xFF},
	"violet":               color.RGBA{R: 0xEE, G: 0x82, B: 0xEE, A: 0xFF},
	"wheat":                color.RGBA{R: 0xF5, G: 0xDE, B: 0xB3, A: 0xFF},
	"white":                color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	"whitesmoke":           color.RGBA{R: 0xF5, G: 0xF5, B: 0xF5, A: 0xFF},
	"yellow":               color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF},
	"yellowgreen":          color.RGBA{R: 0x9A, G: 0xCD, B: 0x32, A: 0xFF},
}
