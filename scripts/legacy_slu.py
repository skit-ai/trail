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
        entity_value = values[0]
        entity_type = entity_value.get("type")

        if entity_type == "interval":
            value = {
                "from": {"value": entity_value["value"].get("from")},
                "to": {"value": entity_value["value"].get("to")}
            }
        else:
            value = entity_value

        entity_list.append(dict(
            text=entity.get("body"),
            type=entity_type,
            score=entity.get("score", 0),
            value=value,
        ))

    return json.dumps(entity_list)


def get_labeled_items(data, type_):
    with open(f"{BASE_PATH}/predicted.{type_}.csv", "w") as fp:
        writer = csv.writer(fp)
        writer.writerow(columns_map[type_])
        for record in data:
            label_func = label_func_map.get(type_)
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
        "aux_entities": ["id", "entities"],
    }

    get_labeled_items(data, type_="intents")
    get_labeled_items(data, type_="aux_entities")
