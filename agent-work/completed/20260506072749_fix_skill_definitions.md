# fix_skill_definitions

## Status: completed (20260506073813)

## Context
The repository ships `validate-kata` skills for Codex, OpenCode, and Claude-compatible locations.
Current problems:
- The `.opencode` and `.codex` `SKILL.md` files do not start with required YAML frontmatter.
- The `.claude` skill file is a broken symlink, so the definition cannot be loaded.
- The three copies diverge in behavior and wording, which increases maintenance risk and can make one host valid while another is stale.

## Value Proposition
Make the repository-provided `validate-kata` skill discoverable and valid across supported hosts, so users can reliably invoke kata validation without manual repair.

## Alternatives considered (with trade-offs)
1. Keep separate skill bodies per host.
Trade-off: allows host-specific examples, but duplicates content and increases drift risk.
2. Make one canonical skill file and reuse it from compatible locations.
Trade-off: best consistency and lowest maintenance, but examples must stay generic enough for multiple hosts.
3. Remove Claude-compatible skill and document only `.claude/commands`.
Trade-off: cleaner Anthropic-native setup, but changes current repo structure and user expectations more than necessary.

Chosen approach: keep one canonical skill body aligned to the Agent Skills format, mirror it into `.codex`, and repair the `.claude` compatibility copy.

## Todos
- [x] Inspect current skill files and confirm validity requirements
- [x] Rewrite canonical `validate-kata` skill with valid frontmatter and consistent instructions
- [x] Repair `.claude` compatibility copy so it resolves correctly
- [x] Verify final on-disk state and check for remaining edge cases

## Acceptance Criteria
- `.opencode/skills/validate-kata/SKILL.md` starts with valid YAML frontmatter and a matching skill name
- `.codex/skills/validate-kata/SKILL.md` is also valid and consistent with the canonical skill
- `.claude/skills/validate-kata/SKILL.md` resolves correctly instead of pointing to a broken target
- Final skill content documents how `validate-kata` works without contradictory invocation details

## Notes
Primary references:
- OpenCode skills docs: https://opencode.ai/docs/skills/
- OpenAI skills follow Agent Skills open standard: https://help.openai.com/en/articles/20001066-skills-in-chatgpt

Edge cases found during implementation:
- `.claude/skills/validate-kata/SKILL.md` was a dangling symlink because `../../.opencode/...` resolved under `.claude/` instead of repo root
- Frontmatter-free markdown may still be human-readable, but it is not reliable for modern skill loaders that expect metadata extraction
