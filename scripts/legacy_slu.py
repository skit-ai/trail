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
        values = entity.get("values", [{}])
        value = values[0].get("value")

        entity_list.append(dict(
            text=entity.get("body"),
            type=entity.get("type"),
            score=entity.get("score", 0),
            value=entity.get("value", value),
        ))

    return json.dumps(entity_list)


def get_labeled_items(data, type_):
    with open(f"{BASE_PATH}/predicted.{type_}.csv", "w") as fp:
        print(f"{type_}")
        writer = csv.writer(fp)
        writer.writerow(columns_map[type_])
        for record in data:
            label_func = label_func_map.get(type_)
            print(record["uuid"])
            if record["slu_response"]["response"].get(type_):
                label = label_func(record["slu_response"]["response"][type_])
            else:
                label = ""

            writer.writerow([record["uuid"], label])


if __name__ == "__main__":
    data = json.loads(sys.stdin.read())

    BASE_PATH = "."
    label_func_map = {
        "intents": _generate_intent_labels,
        "aux_entities": _generate_entity_labels
    }
    columns_map = {
        "intents": ["id", "intent"],
        "aux_entities": ["id", "entity"],
    }

    get_labeled_items(data, type_="intents")
    get_labeled_items(data, type_="aux_entities")
