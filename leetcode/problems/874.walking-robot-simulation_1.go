package problems

/*
 * @lc app=leetcode id=874 lang=golang
 *
 * [874] Walking Robot Simulation
 *
 * https://leetcode.com/problems/walking-robot-simulation/description/
 *
 * algorithms
 * Easy (31.01%)
 * Total Accepted:    8.3K
 * Total Submissions: 26.7K
 * Testcase Example:  '[4,-1,3]\n[]'
 *
 * A robot on an infinite grid starts at point (0, 0) and faces north.  The
 * robot can receive one of three possible types of commands:
 *
 *
 * -2: turn left 90 degrees
 * -1: turn right 90 degrees
 * 1 <= x <= 9: move forward x units
 *
 *
 * Some of the grid squares are obstacles.
 *
 * The i-th obstacle is at grid point (obstacles[i][0], obstacles[i][1])
 *
 * If the robot would try to move onto them, the robot stays on the previous
 * grid square instead (but still continues following the rest of the route.)
 *
 * Return the square of the maximum Euclidean distance that the robot will be
 * from the origin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: commands = [4,-1,3], obstacles = []
 * Output: 25
 * Explanation: robot will go to (3, 4)
 *
 *
 *
 * Example 2:
 *
 *
 * Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
 * Output: 65
 * Explanation: robot will be stuck at (1, 4) before turning left and going to
 * (1, 8)
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= commands.length <= 10000
 * 0 <= obstacles.length <= 10000
 * -30000 <= obstacle[i][0] <= 30000
 * -30000 <= obstacle[i][1] <= 30000
 * The answer is guaranteed to be less than 2 ^ 31.
 *
 *
 */
func robotSim(commands []int, obstacles [][]int) int {
	x, y := 0, 0

	xMap, yMap := packageObstacle(obstacles)

	maxSquare := 0
	direction := 2 //1:wast 2:north 3:east 4:south
	for _, command := range commands {
		if command < 0 { //turn left
			if command == -2 {
				direction--
				if direction == 0 {
					direction = 4
				}
			} else { //turn right
				direction++
				if direction == 5 {
					direction = 1
				}
			}

			continue
		}

		switch direction {
		case 1: //y坐标不变，x坐标减少
			yObstacle, ok := yMap[y]
			if !ok {
				x -= command
			} else {
				for command > 0 {
					if yObstacle[x-1] {
						break
					}

					x--
					command--
				}
			}
		case 3: //y坐标不变，x坐标增加
			yObstacle, ok := yMap[y]
			if !ok {
				x += command
			} else {
				for command > 0 {
					if yObstacle[x+1] {
						break
					}

					x++
					command--
				}
			}
		case 2: //x坐标不变，y坐标增加
			xObstacle, ok := xMap[x]
			if !ok {
				y += command
			} else {
				for command > 0 {
					if xObstacle[y+1] {
						break
					}

					y++
					command--
				}
			}
		case 4: //x坐标不变，y坐标减少
			xObstacle, ok := xMap[x]
			if !ok {
				y -= command
			} else {
				for command > 0 {
					if xObstacle[y-1] {
						break
					}

					y--
					command--
				}
			}
		}

		seq := x*x + y*y

		if seq > maxSquare {
			maxSquare = seq
		}
	}

	return maxSquare
}

func packageObstacle(obstacles [][]int) (map[int]map[int]bool, map[int]map[int]bool) {
	xMap := make(map[int]map[int]bool, len(obstacles))
	yMap := make(map[int]map[int]bool, len(obstacles))

	for _, obstacle := range obstacles {
		ox, oy := obstacle[0], obstacle[1]

		xm, ok := xMap[ox]
		if ok {
			xm[oy] = true
		} else {
			xm = map[int]bool{oy: true}
			xMap[ox] = xm
		}

		ym, ok := yMap[oy]
		if ok {
			ym[ox] = true
		} else {
			ym = map[int]bool{ox: true}
			yMap[oy] = ym
		}
	}

	return xMap, yMap
}
