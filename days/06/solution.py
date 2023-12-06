import math

with open("days/06/input.txt", "r") as f:
    content = f.readlines()

a, b = content

def parse_nums(s: str):
    return int(s.replace(" ","").replace("\n", "").replace("Time:", "").replace("Distance:", ""))

def distance_for_time(t: int, T: int) -> int:
    return (T - t) * t
    

time, distance = parse_nums(a), parse_nums(b)

front = back = 0

for t in range(time):
    if distance_for_time(t, time) > distance:
        front = t
        break

for t in range(time, 0, -1):
    if distance_for_time(t, time) > distance:
        back = t
        break

print(back - front + 1)
