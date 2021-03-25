```shell
$ go test -bench .
BenchmarkAtomicCounterAdd-8       	195539016	         5.827 ns/op
BenchmarkAtomicCounterRead-8      	1000000000	         0.3174 ns/op
BenchmarkAtomicCounterAddRead-8   	50456290	        23.75 ns/op
BenchmarkCounterAdd-8             	42463191	        26.73 ns/op
BenchmarkCounterRead-8            	70519336	        16.21 ns/op
BenchmarkCounterAddRead-8         	12227102	        98.35 ns/op
PASS
ok  	bench	7.852s
```