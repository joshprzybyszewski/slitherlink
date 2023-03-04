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
|5x5 medium|250.36µs|366.52µs|385.72µs|426.1µs|498.63µs|638.16µs|75|
|10x10 medium|333.93µs|619.08µs|684.72µs|766.02µs|859.54µs|901.75µs|70|
|15x15 medium|636.5µs|1.01ms|1.11ms|1.34ms|1.53ms|2.02ms|70|
|20x20 medium|835.51µs|1.67ms|1.89ms|2.18ms|2.42ms|2.86ms|67|
|5x5 hard|243.48µs|326.78µs|422.56µs|1.06ms|1.2ms|1.36ms|75|
|10x10 hard|432.14µs|1.43ms|1.72ms|2.05ms|5.35ms|8.8ms|73|
|15x15 hard|1.66ms|2.29ms|3.15ms|6.85ms|108.03ms|292.63ms|73|
|20x20 hard|2.83ms|6.72ms|18.83ms|55.96ms|1.04s|5s|71|
|25x30 medium|1.39ms|2.61ms|2.98ms|3.67ms|4.12ms|4.3ms|61|
|25x30 hard|7.06ms|57.41ms|209.42ms|2.06s|5.01s|5.02s|73|
|7x7 medium|208.99µs|342.52µs|373.4µs|430.37µs|497.09µs|606µs|66|
|7x7 hard|277.61µs|480.1µs|1.08ms|1.23ms|1.69ms|2.56ms|70|

_Last Updated: 04 Mar 23 09:07 CST_
</resultsMarker>
