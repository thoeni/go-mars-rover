###Mars Rover TDD Kata [![Build Status](https://travis-ci.org/thoeni/go-mars-rover.svg?branch=master)](https://travis-ci.org/thoeni/go-mars-rover)   [![Coverage Status](http://coveralls.io/repos/github/thoeni/go-mars-rover/badge.svg?branch=master)](https://coveralls.io/github/thoeni/go-mars-rover?branch=master)

###Problem:

Mars Rover is an example of a nice Kata to be implemented with TDD style

[x] Develop an api that moves a rover around on a grid.
[ ] You are given the initial starting point (x,y) of a rover and the direction (N,S,E,W) it is facing.
[x] The rover receives a character array of commands.
[x] Implement commands that move the rover forward/backward (f,b).
[x] Implement commands that turn the rover left/right (l,r).
[x] Implement wrapping from one edge of the grid to another. (planets are spheres after all).
[ ] Implement obstacle detection before each move to a new square. If a given sequence of commands encounters an obstacle, the rover moves up to the last possible point and reports the obstacle.
[x] Example: The rover is on a 100x100 grid at location (0, 0) and facing NORTH. The rover is given the commands "ffrff" and should end up at (2, 2).

###Solution
I've worked on this while learning TDD with Golang, since I've started a course on[Pluralsight](https://app.pluralsight.com/library/courses/go-testing-applications/table-of-contents)
The solution is not meant to be an example to follow :-)