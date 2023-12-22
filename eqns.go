package main

import (
	"math"
)

func compTotCurrent(cur Status, c Parameters) (Status, float64) {
	cur = updateGates(cur, c)

	cur.out.m = append(cur.out.m, cur.gates.m.Ss()+(cur.out.m[cur.i]-cur.gates.m.Ss())*math.Exp(-cur.dt/cur.gates.m.Tau()))
	cur.out.h = append(cur.out.h, cur.gates.h.Ss()+(cur.out.h[cur.i]-cur.gates.h.Ss())*math.Exp(-cur.dt/cur.gates.h.Tau()))
	cur.out.j = append(cur.out.j, cur.gates.j.Ss()+(cur.out.j[cur.i]-cur.gates.j.Ss())*math.Exp(-cur.dt/cur.gates.j.Tau()))
	cur.out.d = append(cur.out.d, cur.gates.d.Ss()+(cur.out.d[cur.i]-cur.gates.d.Ss())*math.Exp(-cur.dt/cur.gates.d.Tau()))
	cur.out.f = append(cur.out.f, cur.gates.f.Ss()+(cur.out.f[cur.i]-cur.gates.f.Ss())*math.Exp(-cur.dt/cur.gates.f.Tau()))
	cur.out.X = append(cur.out.X, cur.gates.X.Ss()+(cur.out.X[cur.i]-cur.gates.X.Ss())*math.Exp(-cur.dt/cur.gates.X.Tau()))
	cur.out.Xi = append(cur.out.Xi, cur.gates.Xi.val)
	cur.out.Kp = append(cur.out.Kp, cur.gates.Kp.val)
	cur.out.K1 = append(cur.out.K1, cur.gates.K1.Ss())

	//handle currents
	Esi := 7.7 - 13.0287*math.Log(cur.out.Cai[cur.i]/c.Cao)

	Vm := cur.out.V[cur.i]

	cur.out.INa = append(cur.out.INa, c.gbarNa*math.Pow(cur.out.m[cur.i+1], 3)*cur.out.h[cur.i+1]*cur.out.j[cur.i+1]*(Vm-c.ENa))
	cur.out.Isi = append(cur.out.Isi, 0.09*cur.out.d[cur.i+1]*cur.out.f[cur.i+1]*(Vm-Esi))
	cur.out.Ik = append(cur.out.Ik, c.gbarK*cur.out.X[cur.i+1]*cur.out.Xi[cur.i+1]*(Vm-c.Ek))
	cur.out.IK1 = append(cur.out.IK1, c.gbarK1*cur.out.K1[cur.i+1]*(Vm-c.Ek1))
	cur.out.IKP = append(cur.out.IKP, c.gbarKP*cur.out.Kp[cur.i+1]*(Vm-c.Ekp))
	cur.out.Ib = append(cur.out.Ib, c.gbarB*(Vm-c.Eb))
	Itot := cur.out.INa[cur.i+1] + cur.out.Isi[cur.i+1] + cur.out.Ik[cur.i+1] + cur.out.IK1[cur.i+1] + cur.out.IKP[cur.i+1] + cur.out.Ib[cur.i+1]

	return cur, Itot
}

//Gate handling
func updateGates(cur Status, c Parameters) Status {
	cur.V = cur.out.V[cur.i]

	//Xi (analytical)
	if cur.V > -100 {
		cur.gates.Xi.val = 2.837 * (math.Exp(0.04*(cur.V+77)) - 1) / ((cur.V + 77) * (math.Exp(0.04 * (cur.V + 35))))
	} else {
		cur.gates.Xi.val = 1
	}

	//X
	cur.gates.X.alpha = 0.0005 * math.Exp(0.083*(cur.V+50)) / (1 + math.Exp(0.057*(cur.V+50)))
	cur.gates.X.beta = 0.0013 * math.Exp(-0.06*(cur.V+20)) / (1 + math.Exp(-0.04*(cur.V+20)))

	//K1
	cur.gates.K1.alpha = 1.02 / (1 + math.Exp(0.2385*(cur.V-c.Ek1-59.215)))
	cur.gates.K1.beta = (0.49124*math.Exp(0.08032*(cur.V-c.Ek1+5.476)) + math.Exp(0.06175*(cur.V-c.Ek1-594.31))) / (1 + math.Exp(-0.5143*(cur.V-c.Ek1+4.753)))

	//Kp (analytical)
	cur.gates.Kp.val = 1 / (1 + math.Exp((7.488-cur.V)/5.98))

	//m
	cur.gates.m.alpha = 0.32 * (cur.V + 47.13) / (1 - math.Exp(-0.1*(cur.V+47.13)))
	cur.gates.m.beta = 0.08 * math.Exp(-cur.V/11)

	//h
	//j
	if cur.V >= -40 {
		cur.gates.h.alpha = 0
		cur.gates.h.beta = 1 / (0.13 * (1 + math.Exp((cur.V+10.66)/-11.1)))

		cur.gates.j.alpha = 0
		cur.gates.j.beta = 0.3 * math.Exp(-2.535*math.Pow10(-7)*cur.V) / (1 + math.Exp(-0.1*(cur.V+32)))
	} else {
		cur.gates.h.alpha = 0.135 * math.Exp((80+cur.V)/-6.8)
		cur.gates.h.beta = 3.56*math.Exp(0.079*cur.V) + 3.1*math.Pow10(5)*math.Exp(0.35*cur.V)

		cur.gates.j.alpha = (-1.2714*math.Pow10(5)*math.Exp(0.2444*cur.V) - 3.474*math.Pow10(-5)*math.Exp(-0.04391*cur.V)) * (cur.V + 37.78) / (1 + math.Exp(0.311*(cur.V+79.23)))
		cur.gates.j.beta = 0.1212 * math.Exp(-0.01052*cur.V) / (1 + math.Exp(-0.1378*(cur.V+40.14)))
	}

	//d
	cur.gates.d.alpha = 0.095 * math.Exp(-0.01*(cur.V-5)) / (1 + math.Exp(-0.072*(cur.V-5)))
	cur.gates.d.beta = 0.07 * math.Exp(-0.017*(cur.V+44)) / (1 + math.Exp(0.05*(cur.V+44)))

	//f
	cur.gates.f.alpha = 0.012 * math.Exp(-0.008*(cur.V+28)) / (1 + math.Exp(0.15*(cur.V+28)))
	cur.gates.f.beta = 0.0065 * math.Exp(-0.02*(cur.V+30)) / (1 + math.Exp(-0.2*(cur.V+30)))

	return cur
}
