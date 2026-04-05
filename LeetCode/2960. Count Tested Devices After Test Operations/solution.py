class Solution:
    def countTestedDevices(self, batteryPercentages: List[int]) -> int:
        count = 0
        p = 0
        for pin in batteryPercentages:
            if pin - p > 0:
                count += 1
                p += 1 
        return count