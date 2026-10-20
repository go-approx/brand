// Command mark writes the go-approx mark, svg/color/go-approx.svg:
//
//	cd mark && go run . > ../svg/color/go-approx.svg
//
// go-approx is approximation drawn as its own definition: a smooth curve
// through a handful of sample points. A function is sampled at `knots` evenly
// spaced abscissae, the curve is the Catmull-Rom spline through those knots
// (converted to cubic Beziers), and a dot marks every few knots: the samples
// the curve is said to pass through.
//
// Nothing is placed by hand. The knots follow from the function and their
// count; the tangents follow from the neighbouring knots; the dots sit on
// knots; the glyph is centred on the square from the extent of the samples.
// Change the function or a count and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. The sample dots are the one
// filled shape: a dot wider than the line is what tells a sample from the
// curve. logos/brandgen renders the format matrix from the file this writes;
// the kit workflow runs this again and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#F6AD55" // the gradient's light stop
	dark  = "#B45309" // its dark stop (5.0:1 on white)

	stroke = 13.0 // the family's line width

	left, right = 44.0, 212.0 // the glyph's horizontal extent (knot centres)
	amplitude   = 46.0        // half the vertical extent of the knots
	dot         = 14.0        // radius of a sample dot
	knots       = 13          // points the spline is built through
	every       = 3           // a dot on every this-many-th knot
)

// f is the function being approximated, on [0,1].
func f(u float64) float64 { return math.Cos(1.5 * math.Pi * u) }

type pt struct{ x, y float64 }

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	// Sample f; y grows downwards, so a larger f is drawn higher.
	lo, hi := math.Inf(1), math.Inf(-1)
	vals := make([]float64, knots)
	for i := range vals {
		vals[i] = f(float64(i) / float64(knots-1))
		lo, hi = math.Min(lo, vals[i]), math.Max(hi, vals[i])
	}
	// Scale so the extremes sit at +-amplitude about the centre of the square.
	scale := amplitude / ((hi - lo) / 2)
	mid := (hi + lo) / 2
	p := make([]pt, knots)
	for i, v := range vals {
		p[i] = pt{
			x: left + (right-left)*float64(i)/float64(knots-1),
			y: 128 - (v-mid)*scale,
		}
	}

	// Catmull-Rom through the knots, as cubic Beziers; the end knots are
	// repeated so the end tangents point at the neighbouring knot.
	at := func(i int) pt { return p[min(max(i, 0), knots-1)] }
	d := fmt.Sprintf("M%s %s", num(p[0].x), num(p[0].y))
	for i := 0; i < knots-1; i++ {
		p0, p1, p2, p3 := at(i-1), at(i), at(i+1), at(i+2)
		c1 := pt{p1.x + (p2.x-p0.x)/6, p1.y + (p2.y-p0.y)/6}
		c2 := pt{p2.x - (p3.x-p1.x)/6, p2.y - (p3.y-p1.y)/6}
		d += fmt.Sprintf("C%s %s %s %s %s %s", num(c1.x), num(c1.y), num(c2.x), num(c2.y), num(p2.x), num(p2.y))
	}
	lines := []string{fmt.Sprintf(`<path d="%s"/>`, d)}
	var dots []string
	for i := 0; i < knots; i += every {
		dots = append(dots, fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`, num(p[i].x), num(p[i].y), num(dot)))
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-approx">
  <defs><linearGradient id="apgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#apgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
  <g fill="#FFFFFF">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(lines, "\n    "), strings.Join(dots, "\n    "))
}
