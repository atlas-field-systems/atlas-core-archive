# Atlas Core

Atlas Core is a control plane for observing Entities, preserving operational data, and tasking Assets. This repository contains Core, Protocol and SDK work. The operator-facing Command Interface and Asset-side execution systems are separate consumers.

## Start here

| Goal | Read |
| --- | --- |
| Understand the system and find a contract | [Documentation](docs/README.md) |
| Develop or review a change with PStack | [Contributing](CONTRIBUTING.md) and [PStack workflow](docs/development/README.md) |
| Find the current implementation boundary | [Implementation sequence](docs/architecture/implementation-sequence.md) |
| Understand the vocabulary | [Glossary](GLOSSARY.md) |
| Find who owns a rule or decision | [Documentation ownership](docs/documentation.md) |

## Implementation status

Atlas Core is being extracted and simplified from Atlas Modernization. The delivered foundation is the bounded [Slice 0 contract foundation](tests/contract/README.md), with a [focused Plugin bookkeeping component](Atlas%20Core/plugins/README.md). Its public Protocol has no operational routes. Operational Core/SDK workflows remain implementation work; accepted design and an open implementation PR are not evidence that a workflow has shipped.

The documentation includes the accepted S1 scope and [deployment-clock decision](docs/adr/0029-use-deployment-clocks-and-preserve-event-times.md). The [implementation sequence](docs/architecture/implementation-sequence.md) owns the delivery plan and remaining qualifications. The [testing strategy](docs/testing-strategy.md) distinguishes foundation, operational, independent-extension and field evidence.

## Verify the foundation

From the repository root:

```sh
python3 scripts/verify.py --bootstrap
```

This runs the repository's required foundation checks. It does not establish unimplemented operational behavior, published compatibility, production capacity or physical Asset readiness. See [verification](docs/development/verification.md) for documentation-only changes, exact-revision evidence and the checks required for implementation work.

## Repository boundaries

[Core](Atlas%20Core/README.md) owns the server implementation. [Protocol](Atlas%20Protocol/README.md) owns shared external contracts. [SDK](Atlas%20SDK/README.md) owns supported client behavior. The [system outline](docs/architecture/system-outline.md) assigns responsibilities, and the [system design](docs/architecture/system-design.md) explains collaboration across those boundaries.
