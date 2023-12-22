<img src="https://i.imgur.com/8omBgvc.png" width="300" height="300">

# goLR1

a GoLang-based implementation of the 1991 Luo-Rudy model of the ventricular action potential. integrated with the first order runge-kutta method (euler).

original paper : https://pubmed.ncbi.nlm.nih.gov/1709839/

## Usage

alter parameters (resting membrane voltage, stimuli, etc.) in ``inits.go`` or ``main.go`` (or keep default). then run with
```go run goLR1```
or ``go build goLR1`` and run the executable. output ``out.txt`` can be plotted in MATLAB with ``compile/outRunner.m``.

## Output

the expected output for a 60µA stimulus from 30ms to 30.5ms (tspan from 0ms to 600ms) would be the following three plots

![4.png](https://i.imgur.com/Rk0RyC0.png)

## Authors

- [@mylesmax](https://www.github.com/mylesmax)

## License

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
