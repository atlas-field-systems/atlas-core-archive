# Contributing to Atlas Core

Atlas uses [PStack-T3](docs/development/README.md) to investigate, implement and review work. The workflow serves the project contracts; it does not define a separate version of them. Start with the [documentation index](docs/README.md).

## Pick up a unit of work

Read the request or issue, its comments, dependencies and required baseline. Resolve factual questions by inspecting the repository and its evidence. Escalate only a product choice, an irreversible action outside authorization, or a conflict that the available sources cannot settle. Existing authorization remains valid; do not ask for it again.

An accepted specification may require documentation newer than the target branch. Carry that prerequisite explicitly, or identify the blocking prerequisite PR. Do not silently substitute the target branch's older contract. [Work tracking](docs/development/work-tracking.md) defines the brief and dependency conventions.

## Execute with PStack

Use `poteto-mode` for a bounded task. Use a standing `brigade` coordinator only for an authorized queue or program. Give each writer an isolated worktree and an explicit write scope. Prefer one complete, observable workflow over disconnected scaffolding. A small mechanical edit does not require a coordinator, multiple proposals or a research dossier.

Read the applicable [behavioral references](docs/topics/README.md), [architecture](docs/architecture/system-design.md), [ADRs](docs/adr/README.md), [coding standards](CODING_STANDARDS.md) and [testing strategy](docs/testing-strategy.md). Preserve the [decision-authority rules](docs/documentation.md#decision-authority).

## Review and publication

Review conventions separately from the requested behavior. Use PStack's adversarial review for material changes, validate each finding against a requirement and concrete consequence, and distinguish a defect from a suggestion. Passing style checks is not proof that the right behavior was implemented.

Review the exact candidate commit and report the checks actually run. Any subsequent edit invalidates the evidence it affects. Use the [verification procedure](docs/development/verification.md); never imply independent or cross-model review when it did not occur.

A request to implement does not itself authorize opening a PR. Open a PR when explicitly requested, and leave it unmerged unless merging is explicitly authorized. For authorized publication through PStack's landing queue, use `human` mode unless the user expressly grants another mode. `human` mode still publishes a PR, so do not submit to that queue before publication is authorized.

## Complete the handoff

The PR records the outcome, scope, authoritative requirements, validation commands and results, and unresolved limits. Link its issue or specification when one exists. A partial delivery does not close the full specification. [Documentation completion](docs/documentation.md#documentation-completion) applies even when no application code changes.

Keep durable knowledge in the [assigned documentation home](docs/documentation.md#where-information-belongs). Work logs, model settings, leases and coordinator runtime state are not another source of product requirements.
