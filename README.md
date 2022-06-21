# disusechan

Don`t use Go channels everywhere!
Use Fetching or Readers (interfaces).

## Benchmarks
### With channels 
```
BenchmarkLinesChannel
BenchmarkLinesChannel-4   	    5586	    209052 ns/op
BenchmarkCSVChannel
BenchmarkCSVChannel-4     	    2066	    559391 ns/op
```
### With fetching
```
BenchmarkFetchLines
BenchmarkFetchLines-4     	  115167	      9855 ns/op
BenchmarkFetchCSV
BenchmarkFetchCSV-4       	   27824	     39816 ns/op
```
