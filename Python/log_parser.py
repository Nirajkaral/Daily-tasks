import os

# Get the directory where the script is located
script_dir = os.path.dirname(os.path.abspath(__file__))
log_file_path = os.path.join(script_dir, "log.txt")

file = open(log_file_path, "r")

error_count = 0

for line in file:
    if "ERROR" in line:
        print(line.strip())
        error_count += 1

print("Total Errors:", error_count)

file.close()