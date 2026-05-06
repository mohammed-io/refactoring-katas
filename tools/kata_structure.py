#!/usr/bin/env python3
"""Print a Markdown structure report for one kata/language pair."""

from __future__ import annotations

import argparse
import re
import sys
from collections import Counter
from pathlib import Path

try:
    from tree_sitter_language_pack import get_parser
except ImportError:  # pragma: no cover - exercised by CLI users before install
    get_parser = None


LANGUAGES = {
    "javascript": {"parser": "javascript", "ext": ".js"},
    "js": {"parser": "javascript", "ext": ".js", "dir": "javascript"},
    "python": {"parser": "python", "ext": ".py"},
    "py": {"parser": "python", "ext": ".py", "dir": "python"},
    "ruby": {"parser": "ruby", "ext": ".rb"},
    "rb": {"parser": "ruby", "ext": ".rb", "dir": "ruby"},
    "golang": {"parser": "go", "ext": ".go"},
    "go": {"parser": "go", "ext": ".go", "dir": "golang"},
    "csharp": {"parser": "csharp", "ext": ".cs"},
    "cs": {"parser": "csharp", "ext": ".cs", "dir": "csharp"},
}

CONTROL_TYPES = {
    "if_statement": "if",
    "if_expression": "if",
    "unless": "if",
    "switch_statement": "switch",
    "switch_expression": "switch",
    "case_statement": "case",
    "case": "case",
    "for_statement": "loop",
    "enhanced_for_statement": "loop",
    "for_in_clause": "loop",
    "while_statement": "loop",
    "do_statement": "loop",
    "return_statement": "return",
}

DECLARATION_TYPES = {
    "function_declaration",
    "function_definition",
    "method_definition",
    "method_declaration",
    "method",
    "singleton_method",
    "class_declaration",
    "class_definition",
    "class",
    "module",
    "type_declaration",
    "struct_declaration",
    "interface_declaration",
    "constructor_declaration",
}

ASSERTION_PATTERNS = {
    "javascript": re.compile(r"\bassert\."),
    "python": re.compile(r"(^|\s)assert\s+", re.MULTILINE),
    "ruby": re.compile(r"\b(assert|refute)_?\w*\b"),
    "golang": re.compile(r"\bt\.(Error|Errorf|Fail|Fatal|Fatalf)\b"),
    "csharp": re.compile(r"\bAssert\."),
}


def node_text(source: bytes, node) -> str:
    return source[node.start_byte : node.end_byte].decode("utf-8", errors="replace")


def first_line(text: str, limit: int = 140) -> str:
    line = " ".join(text.strip().splitlines()[0].split())
    return line[: limit - 3] + "..." if len(line) > limit else line


def walk(node):
    yield node
    for child in node.children:
        yield from walk(child)


def discover_kata(root: Path, language: str, kata: str) -> Path:
    config = LANGUAGES[language]
    lang_dir = config.get("dir", language)
    base = root / lang_dir
    if not base.is_dir():
        raise SystemExit(f"Unknown language directory: {lang_dir}")

    prefix = kata if kata[:2].isdigit() else kata.zfill(2)
    matches = sorted(path for path in base.glob(f"{prefix}*") if path.is_dir())
    if not matches:
        raise SystemExit(f"No kata matching {language}/{prefix}*")
    if len(matches) > 1:
        names = ", ".join(path.name for path in matches)
        raise SystemExit(f"Ambiguous kata {prefix}: {names}")
    return matches[0]


def discover_files(kata_dir: Path, language: str) -> tuple[list[Path], list[Path]]:
    ext = LANGUAGES[language]["ext"]
    if language in {"golang", "go"}:
        files = sorted(kata_dir.glob(f"*{ext}"))
        return [p for p in files if not p.name.endswith("_test.go")], [p for p in files if p.name.endswith("_test.go")]

    source_root = kata_dir / "src"
    test_root = kata_dir / "test"
    source_files = sorted(source_root.rglob(f"*{ext}")) if source_root.is_dir() else []
    test_files = sorted(test_root.rglob(f"*{ext}")) if test_root.is_dir() else []
    return source_files, test_files


def parser_for(language: str):
    if get_parser is None:
        raise SystemExit("Missing dependency. Run: make install-tools")
    parser_name = LANGUAGES[language]["parser"]
    try:
        return get_parser(parser_name)
    except Exception as exc:
        raise SystemExit(f"Could not load Tree-sitter parser '{parser_name}': {exc}") from exc


def summarize_ast(parser, path: Path) -> dict:
    source = path.read_bytes()
    tree = parser.parse(source)
    root = tree.root_node
    controls: Counter[str] = Counter()
    node_types: Counter[str] = Counter()
    identifiers: Counter[str] = Counter()
    declarations: list[str] = []
    parse_errors = 0

    def visit(node, control_depth: int = 0):
        nonlocal parse_errors
        if node.is_named:
            node_types[node.type] += 1
        if node.type in {"identifier", "property_identifier", "field_identifier", "type_identifier", "constant", "simple_symbol"}:
            value = node_text(source, node).strip(":@")
            if value and len(value) <= 60:
                identifiers[value] += 1
        if node.has_error or node.type == "ERROR":
            parse_errors += 1
        if node.type in DECLARATION_TYPES:
            declaration = first_line(node_text(source, node))
            if declaration not in {"class", "module"} and declaration not in declarations:
                declarations.append(declaration)
        label = CONTROL_TYPES.get(node.type)
        if label:
            controls[label] += 1
            control_depth += 1
        max_depth = control_depth
        for child in node.children:
            max_depth = max(max_depth, visit(child, control_depth))
        return max_depth

    max_control_depth = visit(root)
    return {
        "path": path,
        "source": source.decode("utf-8", errors="replace"),
        "declarations": declarations,
        "controls": controls,
        "node_types": node_types,
        "identifiers": identifiers,
        "parse_errors": parse_errors,
        "max_control_depth": max_control_depth,
    }


def extract_literals(text: str) -> tuple[list[str], list[str]]:
    strings = []
    for match in re.finditer(r"""(?:"([^"\\]*(?:\\.[^"\\]*)*)"|'([^'\\]*(?:\\.[^'\\]*)*)')""", text):
        value = match.group(1) if match.group(1) is not None else match.group(2)
        if value and len(value) <= 80:
            strings.append(value)
    numbers = re.findall(r"\b\d+(?:\.\d+)?\b", text)
    return sorted(set(strings)), sorted(set(numbers), key=lambda item: (float(item), item))


def extract_tests(language: str, text: str) -> tuple[list[str], int]:
    names: list[str] = []
    case_count = 0

    if language == "javascript":
        names = re.findall(r"\b(?:test|it)\s*\(\s*['\"`]([^'\"`]+)['\"`]", text)
    elif language == "python":
        names = re.findall(r"^\s*def\s+(test_[A-Za-z0-9_]+)\s*\(", text, re.MULTILINE)
        case_count += sum(max(1, body.count(",")) for body in re.findall(r"@pytest\.mark\.parametrize\([^\n]+,\s*\[([^\]]*)\]", text))
    elif language == "ruby":
        names = re.findall(r"^\s*def\s+(test_[A-Za-z0-9_!?]+)", text, re.MULTILINE)
        names.extend(re.findall(r"\btest\s+['\"]([^'\"]+)['\"]\s+do", text))
    elif language == "golang":
        names = re.findall(r"\bfunc\s+(Test[A-Za-z0-9_]+)\s*\(", text)
    elif language == "csharp":
        lines = text.splitlines()
        pending = False
        for line in lines:
            if re.search(r"\[(Fact|Theory)\b", line):
                pending = True
                continue
            if pending:
                match = re.search(r"\b(?:async\s+)?(?:void|Task|ValueTask)\s+([A-Za-z_][A-Za-z0-9_]*)\s*\(", line)
                if match:
                    names.append(match.group(1))
                    pending = False
        case_count += text.count("[InlineData")

    return names, max(len(names), case_count)


def summarize_group(language: str, parser, files: list[Path], root: Path, is_test: bool) -> dict:
    summaries = [summarize_ast(parser, path) for path in files]
    combined = "\n".join(summary["source"] for summary in summaries)
    strings, numbers = extract_literals(combined)
    tests, test_cases = extract_tests(language, combined) if is_test else ([], 0)
    assertions = len(ASSERTION_PATTERNS[language].findall(combined)) if is_test else 0

    controls: Counter[str] = Counter()
    node_types: Counter[str] = Counter()
    identifiers: Counter[str] = Counter()
    declarations: list[str] = []
    parse_errors = 0
    max_depth = 0
    for summary in summaries:
        controls.update(summary["controls"])
        node_types.update(summary["node_types"])
        identifiers.update(summary["identifiers"])
        declarations.extend(summary["declarations"])
        parse_errors += summary["parse_errors"]
        max_depth = max(max_depth, summary["max_control_depth"])

    return {
        "files": [path.relative_to(root) for path in files],
        "declarations": declarations,
        "controls": controls,
        "node_types": node_types,
        "identifiers": identifiers,
        "strings": strings,
        "numbers": numbers,
        "tests": tests,
        "test_cases": test_cases,
        "assertions": assertions,
        "parse_errors": parse_errors,
        "max_control_depth": max_depth,
    }


def bullet_list(items: list[str], empty: str = "none", limit: int = 40) -> list[str]:
    if not items:
        return [f"- {empty}"]
    visible = items[:limit]
    lines = [f"- {item}" for item in visible]
    if len(items) > limit:
        lines.append(f"- ... {len(items) - limit} more")
    return lines


def counter_lines(counter: Counter[str], empty: str = "none") -> list[str]:
    if not counter:
        return [f"- {empty}"]
    return [f"- {name}: {count}" for name, count in sorted(counter.items())]


def top_counter_lines(counter: Counter[str], empty: str = "none", limit: int = 30) -> list[str]:
    if not counter:
        return [f"- {empty}"]
    return [f"- {name}: {count}" for name, count in counter.most_common(limit)]


def render_markdown(language: str, kata_dir: Path, source: dict, tests: dict) -> str:
    lines: list[str] = [
        f"# Kata Structure: {language}/{kata_dir.name}",
        "",
        "## Source Files",
        *bullet_list([str(path) for path in source["files"]]),
        "",
        "## Source Structure",
        f"- declarations: {len(source['declarations'])}",
        *bullet_list(source["declarations"]),
        f"- max control depth: {source['max_control_depth']}",
        f"- parse errors: {source['parse_errors']}",
        "",
        "## Source Control Flow",
        *counter_lines(source["controls"]),
        "",
        "## Source Literals",
        "- strings: " + (", ".join(source["strings"][:50]) if source["strings"] else "none"),
        "- numbers: " + (", ".join(source["numbers"][:50]) if source["numbers"] else "none"),
        "",
        "## Source Identifiers",
        *top_counter_lines(source["identifiers"]),
        "",
        "## Source Node Profile",
        *top_counter_lines(source["node_types"], limit=20),
        "",
        "## Test Files",
        *bullet_list([str(path) for path in tests["files"]]),
        "",
        "## Test Structure",
        f"- test count: {len(tests['tests'])}",
        f"- estimated runtime cases: {tests['test_cases']}",
        f"- assertions: {tests['assertions']}",
        f"- max control depth: {tests['max_control_depth']}",
        f"- parse errors: {tests['parse_errors']}",
        "",
        "## Test Names",
        *bullet_list(tests["tests"]),
        "",
        "## Test Literals",
        "- strings: " + (", ".join(tests["strings"][:80]) if tests["strings"] else "none"),
        "- numbers: " + (", ".join(tests["numbers"][:80]) if tests["numbers"] else "none"),
        "",
        "## Test Identifiers",
        *top_counter_lines(tests["identifiers"]),
        "",
        "## Test Node Profile",
        *top_counter_lines(tests["node_types"], limit=20),
    ]
    return "\n".join(lines) + "\n"


def main(argv: list[str] | None = None) -> int:
    arg_parser = argparse.ArgumentParser(description="Print a Markdown structure report for one refactoring kata.")
    arg_parser.add_argument("--language", "-l", required=True, help="Language directory or alias: javascript, python, ruby, golang, csharp")
    arg_parser.add_argument("--kata", "-k", required=True, help="Kata number or directory prefix, e.g. 07 or 07-switch-factory")
    args = arg_parser.parse_args(argv)

    language = args.language.lower()
    if language not in LANGUAGES:
        choices = ", ".join(sorted(LANGUAGES))
        raise SystemExit(f"Unsupported language '{args.language}'. Choices: {choices}")
    language = LANGUAGES[language].get("dir", language)

    root = Path.cwd()
    kata_dir = discover_kata(root, language, args.kata)
    source_files, test_files = discover_files(kata_dir, language)
    if not source_files:
        raise SystemExit(f"No source files found in {kata_dir}")
    if not test_files:
        raise SystemExit(f"No test files found in {kata_dir}")

    parser = parser_for(language)
    source = summarize_group(language, parser, source_files, root, is_test=False)
    tests = summarize_group(language, parser, test_files, root, is_test=True)
    print(render_markdown(language, kata_dir, source, tests), end="")
    return 0


if __name__ == "__main__":
    sys.exit(main())
