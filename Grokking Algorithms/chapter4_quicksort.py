import random

# This was my solution to a leetcode exercise, used it to understand better

class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        # Kickstart the recursive quicksort on the entire array
        self.quickSort(nums, 0, len(nums) - 1)
        return nums

    def quickSort(self, nums, relativeStart, relativeEnd):

        if relativeStart < relativeEnd:
            pivotIndex = self.partition(nums, relativeStart, relativeEnd)

            self.quickSort(nums, relativeStart, pivotIndex -1)
            self.quickSort(nums, pivotIndex + 1, relativeEnd)

    def partition(self, nums, relativeStart, relativeEnd):
        # Partition has to centralize pivot and move things around. It returns pivot's index
        # 7 9 3 2 5

        # 2 9 3 7 5
        i = relativeStart - 1
        randomIndex = random.randint(relativeStart, relativeEnd)
        nums[randomIndex], nums[relativeEnd] = nums[relativeEnd], nums[randomIndex]
        pivot = nums[relativeEnd] # pivot is last value
        for j in range(relativeStart, relativeEnd):
            if nums[j] < pivot:
                # we swap here. i at first cycle becomes 0 
                i += 1
                temp = nums[i]
                nums[i] = nums[j]
                nums[j] = temp
                # swapped 

        # at the end, we swap central value with pivot
        # i has the index of the latest minor value, in this case 0
        # 5 subs 9

        nums[i+1], nums[relativeEnd] = nums[relativeEnd], nums[i+1]

        return i+1






    
