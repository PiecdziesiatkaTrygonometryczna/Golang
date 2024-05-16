import json
import random

with open("global-shark-attack.json", "r") as file:
    data = json.load(file)
selected_records = random.sample(data, 10)

output_data = []
for record in selected_records:
    output_record = {
        "date": record.get("date"),
        "country": record.get("country"),
        "injury": record.get("injury"),
        "fatal_y_n": record.get("fatal_y_n"),
        "type": record.get("type"),
        "activity": record.get("activity")
    }
    output_data.append(output_record)

with open("data.json", "w") as outfile:
    json.dump(output_data, outfile, indent=1)
