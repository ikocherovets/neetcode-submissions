func search(nums []int, target int) int {
    /*
    nums=[3,4,5,6,1,2]
    target=4
    */
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l + r) / 2
        if target == nums[mid] {
            return mid
        }

        if nums[l] <= nums[mid] { // ліва половина відсортована
            if target > nums[mid] || target < nums[l] {
                l = mid + 1 // target поза [nums[l], nums[mid]] → праворуч
            } else {
                r = mid - 1 // target усередині лівої → ліворуч
            }
        } else { // права половина відсортована
            if target < nums[mid] || target > nums[r] {
                r = mid - 1 // target поза [nums[mid], nums[r]] → ліворуч
            } else {
                l = mid + 1 // target усередині правої → праворуч
            }
        }
    }
    return -1
}