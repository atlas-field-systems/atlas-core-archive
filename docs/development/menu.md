# Atlas Core project brief

Use this brief as the project-purpose input to a PStack Brigade coordinator. Read current source documents when opening or recovering a program; this page is not a second specification or status ledger.

## Purpose

Deliver a small, dependable Core, Protocol and SDK for observing Entities, preserving operational data and tasking Assets. Serve the users, workloads and disconnected field conditions in the [operating model](../architecture/operating-model.md). Follow the ownership boundaries in the [system outline](../architecture/system-outline.md).

## What good looks like

Each unit delivers an authorized, observable result with evidence from the relevant real boundaries. Prefer a complete workflow, a small explicit interface and honest failure semantics over scaffolding, speculative frameworks or a large test count. Use the [implementation sequence](../architecture/implementation-sequence.md) to find the current slice and remaining qualifications.

The initial MVP and its acceptance evidence live in the operating model and [testing strategy](../testing-strategy.md). Do not reinterpret a documentation plan, fixture or open PR as delivered operational capability. In particular, the accepted [deployment-clock decision](../adr/0029-use-deployment-clocks-and-preserve-event-times.md) is part of the S1 documentation baseline, not proof that operational S1 schemas or workflows have shipped.

## Work source and boundaries

Take work from the user's explicit request or authorized queue. GitHub issues carry assigned specifications under [work tracking](work-tracking.md). Preserve their dependencies and required documentation baseline. Do not autonomously turn every research idea, competitor feature or reviewer suggestion into implementation scope.

The Command Interface, Asset OS and physical radio implementations are outside this repository's implementation responsibility. Preserve their integration contracts without importing those products into Core. Internal engineering choices may proceed where they preserve accepted behavior. Product changes follow [decision authority](../documentation.md#decision-authority).

## Operating defaults

Use the [standing orders](house-rules.md) in every brief. Start with at most two units writing or under review at once, further limited by path conflicts and the host's capacity. Use milestone reporting. Do not start recurring intake or schedules without authorization. These are coordination defaults, not runtime configuration already installed by this document.

Publish a PR only when explicitly requested. Use human landing for authorized publication and stop before merging unless the user explicitly authorizes the merge. Follow [verification](verification.md) and record the exact candidate revision and missing evidence.
