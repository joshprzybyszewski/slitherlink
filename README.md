# slitherlink
slitherlink Solver - golang

This is the self-proclaimed World's fastest solver for [the slitherlink puzzle](www.puzzle-loop.com). It is similar to [my Masyu Solver](https://github.com/joshprzybyszewski/masyu).

To run, execute `make compete`.

## Results

Check the [Hall of Fame](https://www.puzzle-loop.com/hall.php?hallsize=9) for the results recorded by the puzzle server (which include network duration of submission). Below are the results of the solver as recorded on my machine.

_NOTE: Update this table with `make results`._

<resultsMarker>

_GOOS: linux_

_GOARCH: amd64_

_cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz_

_Solve timeout: 10s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|5x5 medium|214.2µs|309.6µs|348.67µs|383.72µs|442.32µs|691.79µs|84|
|10x10 medium|331.04µs|473.95µs|592.64µs|679.82µs|781.39µs|864.03µs|73|
|15x15 medium|577.65µs|918.18µs|1.03ms|1.19ms|1.34ms|1.44ms|73|
|20x20 medium|932.63µs|1.42ms|1.68ms|1.92ms|2.24ms|2.34ms|70|
|5x5 hard|190.83µs|306µs|371.28µs|927.27µs|1.03ms|1.29ms|78|
|10x10 hard|355.36µs|1.19ms|1.56ms|2.01ms|6.48ms|7.66ms|76|
|15x15 hard|1.22ms|2.21ms|3.18ms|6.31ms|137.6ms|293.15ms|76|
|20x20 hard|3.48ms|7.16ms|19.58ms|51.41ms|1.01s|10.01s|74|
|25x30 medium|1.88ms|2.53ms|3.13ms|3.48ms|4.23ms|4.27ms|65|
|25x30 hard|6.05ms|55.93ms|244.82ms|2.32s|10s|10.02s|76|
|7x7 medium|202.43µs|335.84µs|395.51µs|450.63µs|511.04µs|587.63µs|69|
|7x7 hard|213.28µs|434.67µs|767.21µs|1.09ms|1.71ms|2.21ms|73|

_Last Updated: 10 Mar 23 14:30 CST_
</resultsMarker>
