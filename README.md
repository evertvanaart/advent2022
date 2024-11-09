# Advent of Code 2022

Fourth year participating, first year using Go, first year publishing my code.

Not in it to win it, but I'll try to comment my code to explain my approach.

Plain Go, zero dependencies, no parallel processing.

```plain
Usage: go run . <task> <input>
 <task>     Day number (two digits) plus part ('a' or 'b')
 <input>    Input file base name, e.g. 'input' or 'sample'
 --profile  Run solution multiple times and compute average duration
Example: go run . 01a sample
```

# Results

The table below shows the average core runtime of each solution, recorded over an average of 20 runs. These times were recorded on a 2021 MacBook Pro using Go version `1.19.3`. The core runtime does not include the time it takes to read the input file and split it into lines, but does include any additional input parsing.

| Day  | Part A (μs) | Part B (μs) |
| :--: | ----------: | ----------: |
|  01  |          37 |          13 |
|  02  |          28 |          31 |
|  03  |           9 |           9 |
|  04  |         295 |         170 |
|  05  |          66 |          63 |
|  06  |          11 |          22 |
|  07  |         135 |         123 |
|  08  |          57 |         519 |
|  09  |         771 |         717 |
|  10  |           4 |           8 |
|  11  |          46 |      21,429 |
|  12  |       1,064 |         682 |
|  13  |         982 |       2,560 |
|  14  |         981 |      15,307 |
|  15  |           7 |          34 |
|  16  |      15,784 |      42,748 |
|  17  |         361 |         969 |
|  18  |         618 |       1,429 |
|  19  |     175,195 |     535,196 |
|  20  |      12,025 |     122,153 |
|  21  |       1,736 |       1,715 |
|  22  |         168 |         424 |
|  23  |       9,935 |     909,846 |
|  24  |      45,494 |     139,658 |
|  25  |           4 |           - |

All solutions run in less than one second, although the B part of Day 23 (the Game of Life variant) comes close. The B part of Day 16 (opening valves in the volcano) also comes with the caveat that this is an inexact solution which just so happens to produce the correct answer for the given input; changing this to an exact solution would result in a far longer runtime.
