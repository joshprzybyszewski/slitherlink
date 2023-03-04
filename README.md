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
|5x5 medium|185.01µs|336.03µs|372.12µs|418.37µs|495.22µs|583.83µs|74|
|10x10 medium|365.77µs|616.46µs|692.78µs|741.79µs|855.88µs|897.04µs|69|
|15x15 medium|550.77µs|1.07ms|1.16ms|1.38ms|1.53ms|2.04ms|69|
|20x20 medium|885.63µs|1.67ms|1.9ms|2.22ms|2.42ms|3.42ms|66|
|5x5 hard|255.79µs|350.58µs|410.75µs|1.05ms|1.32ms|1.45ms|74|
|10x10 hard|549.59µs|1.42ms|1.78ms|2.43ms|4.44ms|8.24ms|72|
|15x15 hard|1.38ms|2.29ms|2.93ms|11.69ms|140.96ms|444.97ms|72|
|20x20 hard|3.02ms|8.22ms|20.88ms|65.19ms|4.27s|5s|70|
|25x30 medium|2.18ms|2.85ms|3.27ms|3.85ms|4.25ms|4.49ms|60|
|25x30 hard|7.74ms|50.24ms|210.1ms|1.4s|5.01s|5.01s|72|
|7x7 medium|258.2µs|333.74µs|367.72µs|414.27µs|514.62µs|605.11µs|65|
|7x7 hard|186.3µs|480.55µs|1.12ms|1.36ms|1.73ms|2.48ms|69|

_Last Updated: 04 Mar 23 08:59 CST_
</resultsMarker>
