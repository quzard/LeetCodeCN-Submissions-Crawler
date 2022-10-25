// 公式套用：num = (randN-1)*N + randN，可得出每个num的概率均相同，再拒绝采样即可。
func rand10() int {
    for {
        num := (rand7()-1)*7 + rand7()
        if num <= 40 {
            return num%10+1
        }
    }
}