func canCompleteCircuit(gas []int, cost []int) int {
	index, curGas, sum := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		val := gas[i] - cost[i]
		curGas += val
		sum += val
		if curGas < 0 {
			curGas = 0
			index = i + 1
		}
	}
	if sum < 0 {
		return -1
	}
	return index
}