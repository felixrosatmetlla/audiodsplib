package types

// Signal struct defines the type used to represent the signals
// and its propertys in the library
//
// Fields:
//  Data: samples of the signal
//  	  represented with a 1D slice where all channels data is put consecutively in order
//  Channels: number of the channels the signal has
//  Samplerate: signal samplerate in samples/s
type Signal struct {
	Data       []float64
	Channels   int
	Samplerate float64
}
