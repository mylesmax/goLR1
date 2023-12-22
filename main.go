package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

var (
	tinfo = []float64{0, 600}
	stim1  = Stim{
		start:     30,
		end:       30.5,
		intensity: 60,
	}
	//stim2  = Stim{
	//	start:     350,
	//	end:       350.5,
	//	intensity: 60,
	//}
	tspan []float64
	Itot float64
	out string
	stims = []Stim{stim1}//, stim2}
)

func main() {
	cur = Status{
		V:     c.Vrest,
		dt:    0.25,
		stim: 0,
		i: 0,
	}
	cur.out.V = []float64{c.Vrest}

	cur := updateGates(cur, c)

	cur.out = struct {
		V, m, h, j, d, f, X, K1, Xi, Kp []float64
		INa, Isi, Ik, IK1, IKP, Ib   []float64
		Cai                          []float64
	}{V:[]float64{c.Vrest}, m: []float64{cur.gates.m.Ss()},
		h: []float64{cur.gates.h.Ss()},
		j: []float64{cur.gates.j.Ss()},
		d: []float64{cur.gates.d.Ss()},
		f: []float64{cur.gates.f.Ss()},
		X: []float64{cur.gates.X.Ss()},
		K1: []float64{cur.gates.K1.Ss()},
		Xi: []float64{cur.gates.Xi.val},
		Kp: []float64{cur.gates.Kp.val}, INa: []float64{0}, Isi: []float64{0}, Ik: []float64{0}, IK1: []float64{0}, IKP: []float64{0}, Ib: []float64{0},
		Cai: []float64{c.Cai}}


	for t := tinfo[0]; t < tinfo[1]; t += cur.dt {
		tspan = append(tspan, t)
	}

	for i, t := range tspan {
		cur.i = i
		cur.stim = 0

		for _, stim := range stims {
			if t == stim.start {
				fmt.Println("Firing stimulus of", stim.intensity, "mV at time", stim.start, "ms with duration", strconv.FormatFloat(stim.end-stim.start, 'f', 2, 64) +  "ms.")
			}

			if t >= stim.start && t <= stim.end {
				cur.stim = stim.intensity
			}
		}

		cur, Itot = compTotCurrent(cur,c)
		dvdt := 1/c.Cm * (cur.stim - Itot)
		cur.out.V = append(cur.out.V, cur.out.V[cur.i] + dvdt * cur.dt)

		dCaidt := -math.Pow10(-4) * cur.out.Isi[cur.i+1] +0.07 *(math.Pow10(-4) - cur.out.Cai[cur.i])
		cur.out.Cai = append(cur.out.Cai, cur.out.Cai[cur.i]+dCaidt * cur.dt)

		out += strconv.FormatFloat(t, 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.V[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.m[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.h[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.j[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.d[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.f[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.X[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.Xi[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.K1[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.Kp[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.INa[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.Isi[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.Ik[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.IK1[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.IKP[i], 'f', 7, 64)
		out += "   "
		out += strconv.FormatFloat(cur.out.Ib[i], 'f', 7, 64)

		out += "\n"
	}
	err := ioutil.WriteFile("out.txt", []byte(out), os.ModePerm)
	if err != nil {
		panic("a")
	}
	fmt.Println("Simulation complete. Outputted variables: t, V, m, h, j, d, f, X, Xi, K1, Kp, INa, Isi, Ik, IK1, IKP, and Ib.")
}

