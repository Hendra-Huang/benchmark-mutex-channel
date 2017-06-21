# benchmark-mutex-channel
Benchmarking simple cache implementation using channel or mutex

cachechannel is implemented using [CSP](https://golang.org/doc/effective_go.html#sharing) concept.
cachemutex is implemented using RWLock

Benchmark result shows that for this case, mutex is having better performance compare to channel.

Any PR is welcome. Cheers
