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

Atlas Core is being extracted and simplified from Atlas Modernization. It includes the [Slice 0 contract foundation](tests/contract/README.md), a [focused Plugin bookkeeping component](Atlas%20Core/plugins/README.md), and the [S1 direct-IP Move To workflow](tests/s1/README.md). S1 implements 19 operational routes, a TLS Core container, the HTTP-mode SDK, local lifecycle management and a retained simulator. [Deployment](deployment/README.md) describes the runnable local setup.

The documentation includes the accepted S1 scope and [deployment-clock decision](docs/adr/0029-use-deployment-clocks-and-preserve-event-times.md). The [implementation sequence](docs/architecture/implementation-sequence.md) owns the delivery plan and remaining qualifications. The [testing strategy](docs/testing-strategy.md) distinguishes foundation, operational, independent-extension and field evidence.

## Verify the implementation

From the repository root:

```sh
python3 scripts/verify.py --bootstrap
```

This runs the required foundation, Plugin component and S1 checks, including real container lifecycle, SDK/direct Protocol parity and the offline CLI workflow. [S1 qualification](tests/s1/README.md) lists prerequisites and executable evidence. These checks establish no published compatibility, production capacity or physical Asset readiness. See [verification](docs/development/verification.md) for documentation-only changes, exact-revision evidence and the checks required for implementation work.

## Repository boundaries

[Core](Atlas%20Core/README.md) owns the server implementation. [Protocol](Atlas%20Protocol/README.md) owns shared external contracts. [SDK](Atlas%20SDK/README.md) owns supported client behavior. The [system outline](docs/architecture/system-outline.md) assigns responsibilities, and the [system design](docs/architecture/system-design.md) explains collaboration across those boundaries.
