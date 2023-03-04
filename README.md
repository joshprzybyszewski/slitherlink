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
|5x5 medium|201.7µs|337.18µs|381.39µs|428.12µs|546.15µs|617.75µs|75|
|10x10 medium|343.86µs|604.63µs|669.82µs|744.99µs|862.85µs|882.83µs|70|
|15x15 medium|546.89µs|1.02ms|1.13ms|1.3ms|1.49ms|1.72ms|70|
|20x20 medium|924.13µs|1.44ms|1.7ms|2.16ms|2.34ms|2.78ms|67|
|5x5 hard|257.23µs|349.67µs|424.74µs|1.03ms|1.23ms|1.4ms|75|
|10x10 hard|517.68µs|1.45ms|1.74ms|2.21ms|5.53ms|8.65ms|73|
|15x15 hard|1.65ms|2.34ms|3.21ms|6.44ms|116.37ms|295.38ms|73|
|20x20 hard|2.17ms|6.54ms|16.46ms|54.97ms|1.03s|10s|71|
|25x30 medium|2.07ms|2.78ms|3.47ms|3.79ms|4.02ms|4.26ms|61|
|25x30 hard|9ms|54.36ms|209.89ms|2.11s|10.01s|10.01s|73|
|7x7 medium|188.54µs|341.03µs|400.25µs|434.32µs|528.62µs|603.42µs|66|
|7x7 hard|364.37µs|517.67µs|1.13ms|1.36ms|1.95ms|2.39ms|70|

_Last Updated: 04 Mar 23 09:49 CST_
</resultsMarker>
