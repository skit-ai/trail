#!/usr/local/bin/python3

import csv
import sys
import json


def _generate_intent_labels(intents):
    intent = max(intents, key=lambda x: x["score"])
    return intent["name"]


def _generate_entity_labels(entities):
    entity_list = []
    for entity in entities:
        entity_list.append(dict(
            text=entity.get("body"),
            type=entity.get("type"),
            score=entity.get("score"),
            value=entity.get("value"),
        ))

    return json.dumps(entity_list)


def get_labeled_items(data, type_):
    records = filter(lambda d: d["response"].get(type_), data)

    with open(f"{type_}_label.csv", "w") as fp:
        writer = csv.writer(fp)
        for record in records:
            label_func = label_func_map.get(type_)
            writer.writerow([record["uuid"], label_func(record["response"][type_])])


if __name__ == "__main__":
    data = json.loads(sys.stdin.read())
    label_func_map = {
        "intents": _generate_intent_labels,
        "entities": _generate_entity_labels
    }

    get_labeled_items(data, type_="intents")
    get_labeled_items(data, type_="entities")
