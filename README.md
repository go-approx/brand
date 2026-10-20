# go-approx — brand assets

Official logos for the **go-approx** organization.

<img src="avatar/go-approx.png" width="160" alt="go-approx">

The mark is approximation drawn as its own definition: a smooth curve through
a handful of sample points. `mark/main.go` samples a function, builds the
Catmull-Rom spline through the samples and puts a dot on every few of them;
nothing is placed by hand, and the curve, its tangents and the dots follow
from the function and the counts.

Its gradient (`#F6AD55` to `#B45309`) is an orange that sets go-approx apart
from the violet of go-numerics, the emerald of go-ndarray, go-fft and
go-notebook, go-plot's blue and go-scikit's slate; its dark end reads on white
at 5.02:1 and its light end on a dark page at 10.14:1 (`landingkit brand`).

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job writes `svg/color/go-approx.svg` again from `mark/main.go` and fails if
a byte differs.

```sh
cd mark && go run . > ../svg/color/go-approx.svg
```

---
*`svg/color/go-approx.svg` is written by `mark/main.go`; every other file is
generated from it by the fleet's Go brand pipeline (brandgen) — rounded square
badge, line glyph, color / white / black variants.*
