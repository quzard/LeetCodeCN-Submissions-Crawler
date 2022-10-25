func numRescueBoats1(people []int, limit int) int {
    sort.Ints(people)
    res := 0
    l, r := 0, len(people) - 1
    for l <= r{
        height := 0
        nums := 0
        for l <= r && height + people[r] <= limit && nums < 2{
            height += people[r]
            r--
            nums++
        }
        for l <= r && height + people[l] <= limit && nums < 2{
            height += people[l]
            l++
            nums++
        }
        res++
    }
    return res
}

func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    l,r := 0,len(people)-1
    out :=0
    for l <= r{
        if people[l]+people[r] <= limit{
            l++
        }
        r--
        out++
    }
    return out
}