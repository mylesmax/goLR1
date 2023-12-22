package main

import (
	"math"
)

var (
	c Parameters
	cur Status
)

func init() {
	c = Parameters{
		Vrest:  -85,                //mV
		Cm:     1,                  //mS/cm^2
		Ko:     5.4,                //mM
		Ki:     145,                //mM
		Nai:    18,                 //mM
		Nao:    140,                //mM
		Cao:    1.8,                //mM
		Cai:    2 * math.Pow10(-4), //mM
		PRNaK:  0.01833,            //permeability
		gbarNa: 23,
		gbarKP: 0.0183,
		gbarB:  0.03921,
		RTF:    (8.3145) * (37 + 273.15) / (96.485),
		Eb:     -59.87,
	}

	c.gbarK = 0.282 * math.Sqrt(c.Ko/5.4)
	c.gbarK1 = 0.6047 * math.Sqrt(c.Ko/5.4)
	c.ENa = c.RTF * math.Log(c.Nao/c.Nai)
	c.Ek = c.RTF * math.Log((c.Ko+c.PRNaK*c.Nao)/(c.Ki+c.PRNaK*c.Nai))
	c.Ek1 = c.RTF * math.Log(c.Ko/c.Ki)
	c.Ekp = c.RTF * math.Log(c.Ko/c.Ki)
}