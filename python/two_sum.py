from typing import List, Dict

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        number_index: Dict[int, int] = {}
        for index, number in enumerate(nums):
            remainder = target - number
            
            if remainder in number_index:
                return [number_index[remainder], index]
            number_index[number] = index

        return nums