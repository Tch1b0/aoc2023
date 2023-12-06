import math

with open("days/06/input.txt", "r") as f:
    content = f.readlines()

a, b = content

def parse_nums(s: str):
    nums = []
    buf = ""
    for c in s:
        if c == " " and len(buf) != 0:
            nums.append(int(buf))
            buf = ""
        elif c.isdigit():
            buf += c
    if len(buf) != 0:
        nums.append(int(buf))
    return nums

def distance_for_time(t: int, T: int) -> int:
    return (T - t) * t
    

a, b = parse_nums(a), parse_nums(b)
races = [x for x in zip(a, b)]

r = []
for time, distance in races:
    m = 0
    for t in range(time):
        d = distance_for_time(t, time)
        if d > distance:
            m += 1
    r += [m]

print(r)
print(math.prod(r))
