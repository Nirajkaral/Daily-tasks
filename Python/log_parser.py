file = open("log.txt", "r")

error_count = 0

for line in file:
    if "ERROR" in line:
        print(line.strip())
        error_count += 1

print("Total Errors:", error_count)

file.close()