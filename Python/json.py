import json

student = {
    "name": "Niraj",
    "age": 22,
    "course": "MCA"
}

json_data = json.dumps(student)

print(json_data)