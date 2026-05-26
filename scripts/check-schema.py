#!/usr/bin/env python3
from pathlib import Path
import json
import re
import sys


ROOT = Path(__file__).resolve().parents[1]
SCHEMA_DIR = ROOT / "schema"

EXPECTED_SCHEMAS = {
    "ni.project.v0.json": "ni.project.v0",
    "ni.contract.v0.json": "ni.contract.v0",
    "ni.lock.v0.json": "ni.lock.v0",
    "ni.readiness-rules.v0.json": "ni.readiness-rules.v0",
    "ni.readiness-profiles.v0.json": "ni.readiness-profiles.v0",
    "ni.feedback.v0.json": "ni.feedback.v0",
    "ni.pressure.v0.json": "ni.pressure.v0",
    "ni.amendment.v0.json": "ni.amendment.v0",
    "ni.harness-candidate.v0.json": "ni.harness-candidate.v0",
}


def load_json(path):
    return json.loads(path.read_text(encoding="utf-8"))


def check_type(value, expected):
    if expected == "object":
        return isinstance(value, dict)
    if expected == "array":
        return isinstance(value, list)
    if expected == "string":
        return isinstance(value, str)
    if expected == "boolean":
        return isinstance(value, bool)
    if expected == "integer":
        return isinstance(value, int) and not isinstance(value, bool)
    return True


def path_join(path, part):
    if path == "$":
        return "$." + part
    return path + "." + part


def validate(value, schema, path="$"):
    errors = []

    if "const" in schema and value != schema["const"]:
        errors.append(f"{path}: expected const {schema['const']!r}, got {value!r}")
    if "enum" in schema and value not in schema["enum"]:
        errors.append(f"{path}: expected one of {schema['enum']!r}, got {value!r}")

    expected_type = schema.get("type")
    if expected_type and not check_type(value, expected_type):
        errors.append(f"{path}: expected {expected_type}, got {type(value).__name__}")
        return errors

    if isinstance(value, str):
        if "minLength" in schema and len(value) < schema["minLength"]:
            errors.append(f"{path}: expected length >= {schema['minLength']}")
        if "pattern" in schema and not re.search(schema["pattern"], value):
            errors.append(f"{path}: expected pattern {schema['pattern']!r}, got {value!r}")

    if isinstance(value, int) and not isinstance(value, bool):
        if "minimum" in schema and value < schema["minimum"]:
            errors.append(f"{path}: expected value >= {schema['minimum']}")

    if isinstance(value, list):
        if "minItems" in schema and len(value) < schema["minItems"]:
            errors.append(f"{path}: expected at least {schema['minItems']} item(s)")
        item_schema = schema.get("items")
        if isinstance(item_schema, dict):
            for index, item in enumerate(value):
                errors.extend(validate(item, item_schema, f"{path}[{index}]"))

    if isinstance(value, dict):
        required = schema.get("required", [])
        for key in required:
            if key not in value:
                errors.append(f"{path}: missing required property {key!r}")

        name_schema = schema.get("propertyNames")
        if isinstance(name_schema, dict):
            for key in value:
                errors.extend(validate(key, name_schema, path_join(path, f"<property {key}>")))

        properties = schema.get("properties", {})
        for key, prop_schema in properties.items():
            if key in value:
                errors.extend(validate(value[key], prop_schema, path_join(path, key)))

        additional = schema.get("additionalProperties", True)
        known = set(properties)
        for key, prop_value in value.items():
            if key in known:
                continue
            if additional is False:
                errors.append(f"{path}: unexpected property {key!r}")
            elif isinstance(additional, dict):
                errors.extend(validate(prop_value, additional, path_join(path, key)))

    return errors


def schema_const(schema):
    return schema.get("properties", {}).get("schema", {}).get("const")


def state_targets():
    targets = []
    add = targets.append
    candidates = [
        (ROOT / ".ni/project.json", "ni.project.v0"),
        (ROOT / ".ni/contract.json", "ni.contract.v0"),
        (ROOT / ".ni/readiness.rules.json", "ni.readiness.rules.v0"),
        (ROOT / ".ni/readiness.profiles.json", "ni.readiness.profiles.v0"),
        (ROOT / ".ni/pressure.json", "ni.pressure.v0"),
        (ROOT / ".ni/harness.candidates.json", "ni.harness_candidates.v0"),
        (ROOT / ".ni/plan.lock.json", "ni.lock.v0"),
    ]
    for path, schema_id in candidates:
        if path.exists():
            add((path, schema_id))

    for path in sorted((ROOT / ".ni/locks").glob("*.json")):
        add((path, "ni.lock.v0"))
    for path in sorted((ROOT / ".ni/amendments").glob("*.json")):
        add((path, "ni.amendment.v0"))
    return targets


def main():
    failures = []
    schemas = {}

    for filename, expected_id in EXPECTED_SCHEMAS.items():
        path = SCHEMA_DIR / filename
        if not path.exists():
            failures.append(f"missing schema: schema/{filename}")
            continue
        try:
            schema = load_json(path)
        except Exception as exc:
            failures.append(f"invalid schema JSON: schema/{filename}: {exc}")
            continue
        schemas[schema_const(schema)] = schema
        if schema.get("$schema") != "https://json-schema.org/draft/2020-12/schema":
            failures.append(f"schema/{filename}: missing draft 2020-12 $schema")
        if schema.get("$id") != expected_id:
            failures.append(f"schema/{filename}: expected $id {expected_id!r}, got {schema.get('$id')!r}")
        if schema.get("type") != "object":
            failures.append(f"schema/{filename}: top-level schema must be an object schema")
        if not schema_const(schema):
            failures.append(f"schema/{filename}: must constrain the state file schema field")

    actual_files = {path.name for path in SCHEMA_DIR.glob("*.json")}
    for extra in sorted(actual_files - set(EXPECTED_SCHEMAS)):
        failures.append(f"unexpected schema file: schema/{extra}")

    for path, schema_id in state_targets():
        schema = schemas.get(schema_id)
        if not schema:
            failures.append(f"{path.relative_to(ROOT)}: no published schema for {schema_id}")
            continue
        try:
            data = load_json(path)
        except Exception as exc:
            failures.append(f"{path.relative_to(ROOT)}: invalid JSON: {exc}")
            continue
        for error in validate(data, schema):
            failures.append(f"{path.relative_to(ROOT)}: {error}")

    feedback_path = ROOT / ".ni/feedback.jsonl"
    if feedback_path.exists():
        schema = schemas.get("ni.feedback.v0")
        for line_number, line in enumerate(feedback_path.read_text(encoding="utf-8").splitlines(), start=1):
            if not line.strip():
                continue
            try:
                data = json.loads(line)
            except Exception as exc:
                failures.append(f"{feedback_path.relative_to(ROOT)}:{line_number}: invalid JSON: {exc}")
                continue
            for error in validate(data, schema):
                failures.append(f"{feedback_path.relative_to(ROOT)}:{line_number}: {error}")

    if failures:
        for failure in failures:
            print(f"Schema check failed: {failure}")
        return 1
    print("Schema check passed")
    return 0


if __name__ == "__main__":
    sys.exit(main())
