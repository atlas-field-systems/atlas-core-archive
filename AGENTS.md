# Atlas Core agent guide

Atlas uses PStack-T3 for engineering work. For a nontrivial task, load the installed `poteto-mode` skill, select its matching playbook, and follow the [Atlas PStack profile](docs/development/README.md). Use the installed runtime for model selection, delegation and worktree isolation. Do not invent tools or claim that a skill ran when it was unavailable.

## Read before acting

- **Behavior, design or documentation:** read `GLOSSARY.md`, then every applicable source in [documentation ownership](docs/documentation.md#where-information-belongs). Check the task's required documentation baseline before treating the checkout as current authority.
- **Code, generators or bindings:** read all of `CODING_STANDARDS.md`.
- **Behavior or tests:** read `docs/testing-strategy.md` and identify the required scenarios. Load `.agents/skills/verify-atlas/SKILL.md` for the verification procedure.
- **Tracked or delegated work:** read [work tracking](docs/development/work-tracking.md) and the [standing orders](docs/development/house-rules.md). A standing coordinator also uses the [project brief](docs/development/menu.md).

## Authority

Agents cannot waive requirements or authorize their own exceptions. [Decision authority](docs/documentation.md#decision-authority) and [documentation completion](docs/documentation.md#documentation-completion) apply to every playbook and delegate.

Answer exploratory questions without changing files. Complete clear requests for action within their authorized scope. Open a pull request only when explicitly requested, ready for review. Merge only when explicitly authorized. PStack's autonomy and landing defaults do not override these boundaries.

Before implementing a substantial new visual direction, present distinct static mocks and wait for the user's choice. After approval, implement it without reopening routine layout, copy or consistency decisions.

## Working loop

Understand the real system, settle the relevant interfaces, implement one verifiable unit, prove its behavior, and review the exact diff. Use `how`, `architect`, `interrogate`, `swarm` or `arena` when the selected playbook calls for them, not as a compulsory ritual for every edit. Keep durable decisions in project documentation, not only in an audit trail or coordinator store.
