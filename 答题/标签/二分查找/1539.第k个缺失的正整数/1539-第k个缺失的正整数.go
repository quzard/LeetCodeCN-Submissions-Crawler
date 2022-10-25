func findKthPositive1(arr []int, k int) int {
    if len(arr) == 0{
        return 0
    }
    last := 0
    for i := 0; i < len(arr); i++{
        if i == 0{
            if arr[0] > k{
                return k
            }
            k -= arr[0] - 1
            last = arr[0]
            continue
        }
        if arr[i] == arr[i - 1] + 1{
            last = arr[i]
            continue
        }
        if arr[i] - arr[i - 1] - 1 >= k{
            return arr[i - 1] + k
        }
        last = arr[i]
        k -= arr[i] - arr[i - 1] - 1
    }
    return last + k
}

func findKthPositive(arr []int, k int) int {
    if arr[0] > k{
        return k
    }
    l := 0
    r := len(arr)
    for l<r{
        mid := l + (r-l)/2
        // arr[mid] - mid - 1 Ϊ ȱ�ٵ���
        if arr[mid] - mid - 1<k{
            l = mid+1
        }else{
            r = mid
        }
    }
    // arr[l-1] ȱʧС��k  arr[l]ȱʧ���ڵ���k��ȱʧ��k ����arr[l-1] ��arr[l] ֮��
    missCnt := arr[l-1] - l
    return arr[l-1] + k - missCnt
}