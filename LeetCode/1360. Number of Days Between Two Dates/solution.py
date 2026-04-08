class Solution:
    def daysBetweenDates(self, date1: str, date2: str) -> int:
        def days(date):
            y, m ,d = map(int, date.split('-'))

            res = d

            month_days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
            for i in range(m-1):
                res += month_days[i]
            if m > 2 and (y % 4 == 0 and y % 100 != 0 or y % 400 == 0):
                res += 1
            for i in range(1971, y):
                res += 365
                if i % 4 == 0 and i % 100 != 0 or i % 400 == 0:
                    res +=1 
            return res
        return abs(days(date1) - days(date2))