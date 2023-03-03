# slitherlink
slitherlink Solver - golang

This is the self-proclaimed World's fastest solver for [the slitherlink puzzle](www.puzzle-loop.com). It is similar to [my Masyu Solver](https://github.com/joshprzybyszewski/masyu).

To run, execute `make compete`.

## Results

Check the [Hall of Fame](www.puzzle-loop.com/hall.php?hallsize=9) for the results recorded by the puzzle server (which include network duration of submission). Below are the results of the solver as recorded on my machine.

_NOTE: Update this table with `make results`._

<resultsMarker>

_GOOS: linux_

_GOARCH: amd64_

_cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz_

_Solve timeout: 1m30s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|5x5 medium|174.78µs|350.61µs|368.93µs|398.69µs|513.47µs|513.47µs|15|
|10x10 medium|300.12µs|454.08µs|496.29µs|555.11µs|590.59µs|590.59µs|10|
|15x15 medium|624.29µs|735.09µs|756.03µs|768.87µs|886.87µs|886.87µs|10|
|20x20 medium|1.07ms|1.2ms|1.26ms|1.38ms|1.57ms|1.57ms|7|
|5x5 hard|868.12µs|1.06ms|1.15ms|1.25ms|1.65ms|1.65ms|15|
|10x10 hard|1.07ms|2.79ms|4.67ms|12.16ms|69.36ms|69.36ms|14|
|15x15 hard|2.64ms|28.09ms|291.56ms|2.63s|5.01s|5.01s|14|
|20x20 hard|170.03ms|1.36s|3.22s|5.01s|5.01s|5.01s|12|
|25x30 medium|815.76µs|815.76µs|1.44ms|1.44ms|1.44ms|1.44ms|2|
|25x30 hard|3.67s|5s|5s|5s|5s|5s|6|
|7x7 medium|109.17µs|177.79µs|222.43µs|249.14µs|251.07µs|251.07µs|7|
|7x7 hard|960.27µs|1.19ms|1.9ms|2.36ms|5.1ms|5.1ms|11|

_Last Updated: 03 Mar 23 07:32 CST_
</resultsMarker>
