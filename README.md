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
|5x5 medium|209.84µs|338.69µs|377.34µs|415.48µs|485.54µs|509.48µs|75|
|10x10 medium|315.98µs|613.8µs|682.93µs|755.93µs|850.94µs|943.97µs|70|
|15x15 medium|572.98µs|1.06ms|1.18ms|1.35ms|1.51ms|1.87ms|70|
|20x20 medium|791.72µs|1.57ms|1.79ms|2.2ms|2.44ms|2.93ms|67|
|5x5 hard|196.13µs|313.53µs|391.26µs|1.02ms|1.29ms|1.32ms|75|
|10x10 hard|278.05µs|1.29ms|1.67ms|2.12ms|5.78ms|9.24ms|73|
|15x15 hard|1.6ms|2.36ms|3.36ms|7.04ms|114.34ms|306.91ms|73|
|20x20 hard|2.86ms|7.26ms|16.38ms|57.19ms|1.05s|10s|71|
|25x30 medium|2.24ms|2.79ms|3.04ms|3.69ms|4.12ms|4.24ms|61|
|25x30 hard|8.02ms|53.32ms|217.55ms|2.01s|10.01s|10.01s|73|
|7x7 medium|230.6µs|354.68µs|402.16µs|434.07µs|555.91µs|650.2µs|66|
|7x7 hard|313.87µs|517.24µs|1.13ms|1.38ms|1.85ms|2.54ms|70|

_Last Updated: 04 Mar 23 10:53 CST_
</resultsMarker>
