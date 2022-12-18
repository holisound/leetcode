class Solution:
    def daysBetweenDates(self, date1: str, date2: str) -> int:
        if date1 > date2:
            date1, date2 = date2, date1
        y1, y2 = int(date1[:4]), int(date2[:4])
        days2 = 0
        while y1 < y2:
            days2 += 366 if isleap(y1) else 365
            y1 += 1
        days1 = dayOfYear(date1)
        days2 += dayOfYear(date2)
        return days2-days1

def isleap(year):
    return year % 400 == 0 or year % 4 == 0 and year % 100 != 0

def dayOfYear(date):
    days=[31,28,31,30,31,30,31,31,30,31,30,31]
    year, month, date = map(int, date.split('-'))
    if isleap(year):
        days[1]+=1
    res=0
    for i in range(month-1):
        res+=days[i]
    return res+date
