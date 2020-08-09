# audiodsplib
A Go library for Audio Digital Signal Processing operations.

This library provides different packages with methods to process signals.
For the moment provides methods to:
- Distort 
- Clipping
- Normalize 
- Change gain
- Sine synthesization

This library provides a custom type Signal to do all the operations. The `audiodsputils` package provides a method to construct a Signal safely, checking that all parameters are valid.

```
type Signal struct {
	Data       []float64
	Channels   int
	Samplerate float64
	NumSamples int
}
```