class Solution:
    def search(self, nums, target):
        # 左闭右闭
        left = 0
        right = len(nums) - 1
        while (left <= right):
            middle = (right + left) // 2
            if nums[middle] == target:
                return middle
            elif nums[middle] > target:
                right = middle - 1
            elif nums[middle] < target:
                left = middle + 1
        return -1


class Solution:
    def search(self, nums, target):
        # 左闭右开
        left = 0
        right = len(nums)
        while left < right:
            mid = (left + right) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
				left = mid + 1
			elif nums[mid] > target:
				right = mid
		return -1



