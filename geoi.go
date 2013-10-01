package geoi

// #cgo CFLAGS: -I/local/libgeoi/include -I/home/anight/projects/misc.git
// #cgo LDFLAGS: /local/libgeoi/lib/libgeoi.a -lm
// #include <geoi.h>
// #include <array.h>
import "C"
import "unsafe"

type GeoiType struct {
	geoi      unsafe.Pointer
	max_coord uint
}

func Geoi(tri_height_m float64, max_tri_n uint) *GeoiType {
	g := &GeoiType{}
	g.geoi = C.geoi_init(C.double(tri_height_m), C.uint(max_tri_n))
	g.max_coord = uint(C.geoi_max_coord(g.geoi))
	return g
}

func (g *GeoiType) MaxCoors() uint {
	return g.max_coord
}

func (g *GeoiType) Cx2ll(tri2_id uint, c [3]uint32) *[2]float64 {
	var ll [2]float64
	C.geoi_cx2ll(g.geoi, C.uint(tri2_id), C.uint(c[0]), C.uint(c[1]), C.uint(c[2]), (*_Ctype_double)(unsafe.Pointer(&ll[0])))
	return &ll
}

func (g *GeoiType) Cs2ll(tri2_id uint, c [3]uint32, sw uint) *[2]float64 {
	var ll [2]float64
	C.geoi_cs2ll(g.geoi, C.uint(tri2_id), C.uint(c[0]), C.uint(c[1]), C.uint(c[2]), C.uint(sw), (*_Ctype_double)(unsafe.Pointer(&ll[0])))
	return &ll
}

func (g *GeoiType) TriCoords(tri2_id uint, ll [2]float64) *[3]uint32 {
	var c [3]uint32
	res := C.geoi_tri_coords(g.geoi, C.uint(tri2_id), C.double(ll[0]), C.double(ll[1]), (*_Ctype_uint)(unsafe.Pointer(&c[0])))
	if uint(res) != 1 {
		return nil
	}
	return &c
}

func (g *GeoiType) Tri2Data(ll *[2]float64) uint {
	ret := C.geoi_tri2_data(g.geoi, C.double(ll[0]), C.double(ll[1]))
	return uint(ret)
}

func (g *GeoiType) Triangles(tri2_data uint, offset uint) []uint32 {
	var t [6]uint32
	res := C.struct_array_s{data: unsafe.Pointer(&t[0]), sz: 4, used: 0, allocated: 6}
	C.geoi_triangles(g.geoi, C.uint(tri2_data), C.uint(offset), &res)
	return t[:res.used]
}
