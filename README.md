# Disuse chan

Don`t use Go-channels everywhere!
Use Fetching or Reader-interfaces.

An implementation with channels is usually an order of magnitude slower than the implementation with fetching.

```
func DataChannel() (chan Data)
```
vs 
```
func FetchData(fn func(Data))
```



## Benchmarks
### With channels 
```
goos: darwin
goarch: amd64
pkg: github.com/denisskin/disusechan
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz

BenchmarkLinesChannel
BenchmarkLinesChannel-4   	    4752	    215411 ns/op
BenchmarkCSVChannel
BenchmarkCSVChannel-4     	    1468	    728230 ns/op
```
### With fetching (~x20)
```
BenchmarkFetchLines
BenchmarkFetchLines-4     	  115438	      9681 ns/op
BenchmarkFetchCSV
BenchmarkFetchCSV-4       	   29938	     38732 ns/op
```
