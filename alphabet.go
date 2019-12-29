package cupfmc

import "time"

func GetAlphabet() map[string][]time.Duration {
  var alphabet map[string][]time.Duration
  alphabet = make(map[string][]time.Duration)
  alphabet["A"] = []time.Duration{Dot,Dash}
  alphabet["B"] = []time.Duration{Dash,Dot,Dot,Dot}
  alphabet["W"] = []time.Duration{Dash,Dash}
  alphabet["G"] = []time.Duration{Dash,Dash,Dot}
  alphabet["D"] = []time.Duration{Dash,Dot,Dot}
  alphabet["E"] = []time.Duration{Dot}
  alphabet["V"] = []time.Duration{Dot,Dot,Dot,Dash}
  alphabet["Z"] = []time.Duration{Dash,Dash,Dot,Dot}
  alphabet["I"] = []time.Duration{Dot,Dot}
  alphabet["J"] = []time.Duration{Dot,Dash,Dash,Dash}
  alphabet["K"] = []time.Duration{Dash,Dot,Dash}
  alphabet["L"] = []time.Duration{Dot,Dash,Dot,Dot}
  alphabet["M"] = []time.Duration{Dash,Dash}
  alphabet["N"] = []time.Duration{Dash,Dot}
  alphabet["O"] = []time.Duration{Dash,Dash,Dash}
  alphabet["P"] = []time.Duration{Dot,Dash,Dash,Dot}
  alphabet["R"] = []time.Duration{Dot,Dash,Dot}
  alphabet["S"] = []time.Duration{Dot,Dot,Dot}
  alphabet["T"] = []time.Duration{Dash}
  alphabet["U"] = []time.Duration{Dot,Dot,Dash}
  alphabet["F"] = []time.Duration{Dot,Dot,Dash,Dot}
  alphabet["H"] = []time.Duration{Dot,Dot,Dot,Dot}
  alphabet["C"] = []time.Duration{Dot,Dash,Dot,Dash}
  alphabet["Q"] = []time.Duration{Dash,Dash,Dot,Dash}
  alphabet["Y"] = []time.Duration{Dash,Dot,Dash,Dash}
  alphabet["X"] = []time.Duration{Dash,Dot,Dot,Dash}
  alphabet["1"] = []time.Duration{Dot,Dash,Dash,Dash,Dash}
  alphabet["2"] = []time.Duration{Dot,Dot,Dash,Dash,Dash}
  alphabet["3"] = []time.Duration{Dot,Dot,Dot,Dash,Dash}
  alphabet["4"] = []time.Duration{Dot,Dot,Dot,Dot,Dash}
  alphabet["5"] = []time.Duration{Dot,Dot,Dot,Dot,Dot}
  alphabet["6"] = []time.Duration{Dash,Dot,Dot,Dot,Dot}
  alphabet["7"] = []time.Duration{Dash,Dash,Dot,Dot,Dot}
  alphabet["8"] = []time.Duration{Dash,Dash,Dash,Dot,Dot}
  alphabet["9"] = []time.Duration{Dash,Dash,Dash,Dash,Dot}
  alphabet["0"] = []time.Duration{Dash,Dash,Dash,Dash,Dash}
  alphabet["."] = []time.Duration{Dot,Dot,Dot,Dot,Dot,Dot}
  alphabet[","] = []time.Duration{Dot,Dash,Dot,Dash,Dot,Dash}
  alphabet[":"] = []time.Duration{Dash,Dash,Dash,Dot,Dot,Dot}
  alphabet[";"] = []time.Duration{Dash,Dot,Dash,Dot,Dash,Dot}
  alphabet["'"] = []time.Duration{Dot,Dash,Dash,Dash,Dash,Dot}
  alphabet["\""] = []time.Duration{Dot,Dash,Dot,Dot,Dash,Dot}
  alphabet["-"] = []time.Duration{Dash,Dot,Dot,Dot,Dot,Dash}
  alphabet["/"] = []time.Duration{Dash,Dot,Dot,Dash,Dot}
  alphabet["_"] = []time.Duration{Dot,Dot,Dash,Dash,Dot,Dash}
  alphabet["?"] = []time.Duration{Dot,Dot,Dash,Dash,Dot,Dot}
  alphabet["!"] = []time.Duration{Dash,Dash,Dot,Dot,Dash,Dash}
  alphabet["+"] = []time.Duration{Dot,Dash,Dot,Dash,Dot}
  alphabet["@"] = []time.Duration{Dot,Dash,Dash,Dot,Dash,Dot}
  return alphabet
}
