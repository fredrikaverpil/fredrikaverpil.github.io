"""Print the pattern of base62 to file."""

s = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
length = len(s)

rows = 200
counter = 1
lines = []
for row in range(rows):
    line = s[counter:] + s*10  + "\n"
    lines.append(line)
    counter += 1

with open("base62.txt", "w+") as outfile:
    outfile.writelines(lines)
