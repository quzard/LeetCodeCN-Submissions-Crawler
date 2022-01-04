func numBusesToDestination(routes [][]int, source int, target int) int {
    if source == target{
        return 0
    }
    // key 为站台 value 为buses， 记录经过站台的公交车
    station2bus := map[int][]int{}
    
    for bus, route := range routes{
        if len(route) == 1{
            continue
        }
        for _, station := range route{
            // 去重
            visited := false
            for _, b := range station2bus[station]{
                if b == bus{
                    visited = true
                }
            }
            if visited == false{
                station2bus[station] = append(station2bus[station], bus)
            }
        }
    }
    // 访问过的公交车
    visited := map[int]bool{}
    // 经过source的公交车
    buses := station2bus[source]
    res := 0
    for len(buses) > 0{
        newBuses := []int{}
        for _, bus := range(buses){
            visited[bus] = true
            for _, station := range routes[bus]{
                if station == target{
                    return res + 1
                }else{
                    // 该站点经过的公交车
                    for _, b := range station2bus[station]{
                        if visited[b] == false{
                            visited[b] = true
                            newBuses = append(newBuses, b)
                        }
                    }
                }
            }
        }
        buses = newBuses
        res++
    }
    return -1
}