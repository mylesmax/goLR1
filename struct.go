package main

type Gate struct {
	alpha, beta float64
}

type GateA struct {
	val float64
}

func (g Gate) Ss() float64 {
	return g.alpha / (g.alpha + g.beta)
}

func (g Gate) Tau() float64 {
	return 1 / (g.alpha + g.beta)
}

type Status struct {
	V, dt, stim float64
	gates       struct {
		m, h, j, d, f, X, K1 Gate
		Xi, Kp               GateA
	}
	out struct {
		V                                  []float64
		m, h, j, d, f, X, K1, Xi, Kp, stim []float64
		INa, Isi, Ik, IK1, IKP, Ib         []float64
		Cai                                []float64
	}
	i int
}

type Parameters struct {
	Vrest,
	Cm,
	Ko,
	Ki,
	Nai,
	Nao,
	Cao,
	Cai,
	PRNaK,
	gbarK,
	gbarK1,
	gbarNa,
	gbarKP,
	gbarB,
	R,
	T,
	F,
	RTF,
	ENa,
	Ek,
	Ek1,
	Ekp,
	Eb float64
}

type Stim struct {
	start, end, intensity float64
}
