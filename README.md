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
|5x5 medium|195.87µs|313.73µs|354.47µs|387.28µs|421.19µs|532.43µs|69|
|10x10 medium|426.98µs|548.47µs|587.09µs|656.6µs|740.93µs|769.63µs|64|
|15x15 medium|488.38µs|888.04µs|1ms|1.18ms|1.28ms|1.34ms|64|
|20x20 medium|955.57µs|1.39ms|1.56ms|1.79ms|1.97ms|2.19ms|61|
|5x5 hard|179.9µs|399.15µs|701.37µs|942.48µs|1.14ms|1.35ms|69|
|10x10 hard|749.66µs|1.38ms|1.8ms|2.88ms|13.9ms|27.59ms|68|
|15x15 hard|1.2ms|2.92ms|6.46ms|43.4ms|3.28s|10s|68|
|20x20 hard|4.17ms|22.73ms|305.81ms|10s|10s|10s|66|
|25x30 medium|1.44ms|2.34ms|2.68ms|3.08ms|3.42ms|3.54ms|56|
|25x30 hard|12.68ms|392.66ms|10s|10s|10s|10s|68|
|7x7 medium|248.73µs|365.16µs|404.56µs|429.06µs|463.98µs|545.25µs|61|
|7x7 hard|264.29µs|868.39µs|974.86µs|1.16ms|2.42ms|18.64ms|65|

_Last Updated: 03 Mar 23 20:03 CST_
</resultsMarker>
