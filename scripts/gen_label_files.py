#!/usr/local/bin/python3

import csv
import sys
import json


def get_labeled_items(data, type_):
    records = filter(lambda d: d["response"].get(type_), data)
    with open(f"{type_}_label.csv", "w") as fp:
        writer = csv.writer(fp)
        for record in records:
            writer.writerow([record["uuid"], json.dumps(record["response"][type_])])


if __name__ == "__main__":
    data = json.loads(sys.stdin.read())
    get_labeled_items(data, type_="intents")
    get_labeled_items(data, type_="entities")
