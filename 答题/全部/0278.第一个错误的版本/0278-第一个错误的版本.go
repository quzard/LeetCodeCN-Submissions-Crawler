import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion1(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + int((r-l)/2)
		if isBadVersion(mid) && isBadVersion(mid-1) {
			r = mid
		} else if isBadVersion(mid) && isBadVersion(mid-1) == false {
			return mid
		} else if isBadVersion(mid) == false && isBadVersion(mid+1) {
			return mid + 1
		} else if isBadVersion(mid) == false && isBadVersion(mid+1) == false {
			l = mid
		}

	}
	return l
}

func firstBadVersion(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
