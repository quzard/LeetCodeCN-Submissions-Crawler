func restoreIpAddresses(s string) []string {
    res := []string{}
    if len(s) < 4{
        return res
    }

    var dfs func(ip string, i int, pos int)

    dfs = func(ip string, i int, pos int){
        if pos == 4 && i == len(s){
            res = append(res, ip)
            return
        }
        if len(s) - i >= 4 - pos && len(s) - i <= (4 - pos) * 3{
            dfs(ip + "." + string(s[i:i+1]), i + 1, pos + 1)
            if s[i] > '0' && s[i] <= '9'{
                if i + 1 < len(s){
                    dfs(ip + "." + string(s[i:i+2]), i + 2, pos + 1)
                }
                if i + 2 < len(s) && string(s[i:i+3]) <= "255"{
                    dfs(ip + "." + string(s[i:i+3]), i + 3, pos + 1)
                }
                
            }
        }
    }
    dfs(string(s[:1]), 1, 1)
    if s[0] > '0' && s[0] <= '9'{
        if 1 < len(s){
            dfs(string(s[:2]), 2, 1)
        }
        if 2 < len(s) && string(s[:3]) <= "255"{
            dfs(string(s[:3]), 3, 1)
        }
    }
    return res
}