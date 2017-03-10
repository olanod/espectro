[![Coverage Status](https://coveralls.io/repos/github/olanod/espectro/badge.svg?branch=master)](https://coveralls.io/github/olanod/espectro?branch=master)

espectro
========

Raw audio spectrum analyzer.
The idea is to have a simple unix like utility that you can pipe raw audio data to and get a frequency spectrum out so it can be further chained with other tools that can make use of that data. E.g. create a nice render of the spectrum or use it with hardware to light up things ^_^

Dependencies
------------
Use glide to install the dependencies: `glide install`

Usage
-----
Reads the standard input and writes to the standard output
```
$ espectro -h
Usage of espectro:
  -i duration
    	Window size (default 50ms)
  -n uint
    	Distribute output spectrum in n (default 10)
  -r uint
    	Sample rate (default 44100)
```
__-i__ is the interval of time used to cut the signal so it can be analyzed and written in chunks.  
__-n__ is used to average the spectrum into _n_ parts. 0 means no averaging takes place.  
__-r__ the sampling rate that the signal comes with. Since we are handling raw audio with no headers it has to be known in advance.
```
$ espectro -n 0 -i 3s < test_audio.raw > test
```
The sample data was created with [sox](http://sox.sourceforge.net). Sox is like a swiss army knife of audio processing, I haven't got that point yet but I want to use it along with _espectro_ to generate spectum data in realtime from an audio source that is playing something.
