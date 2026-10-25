# Atlas Core documentation

Use this index to understand Atlas, find its contracts, or contribute a change. The [documentation guide](documentation.md) identifies authoritative sources and the rules for maintaining them.

## Understand Atlas

| Question | Source |
| --- | --- |
| Who uses Atlas, at what scale, and in which field conditions? | [Operating model](architecture/operating-model.md) |
| What do Core, Protocol, SDK and external consumers own? | [System outline](architecture/system-outline.md) |
| How do modules, storage, generated contracts and local tools collaborate? | [System design](architecture/system-design.md) |
| What does a domain term mean? | [Glossary](../GLOSSARY.md) |
| What is implemented, and what comes next? | [Implementation sequence](architecture/implementation-sequence.md) |

## Behavioral and interface reference

[Topic pages](topics/README.md) own current behavioral rules, exposed routes, SDK operations and explicitly open questions. They are design authority, not a claim that every rule is already implemented.

| Area | Source |
| --- | --- |
| Enrollment, credentials, retirement and access | [Identity and access](topics/identity-and-access.md) |
| Asset reports, Contact and partial updates | [Asset reporting](topics/asset-reporting.md) |
| Commands, Task lifecycle, queues and reconciliation | [Tasks](topics/tasks.md) |
| Uploads, publication, deletion and required results | [Objects](topics/objects.md) |
| Entity identity, Components, Tracks and Geofeatures | [Entities and spatial subjects](topics/tracks-and-geofeatures.md) |
| Plugin capabilities, lifecycle and storage | [Plugins](topics/plugins.md) |
| Start, Stop, Restart, Reset and Installation state | [Dataset lifecycle](topics/dataset-lifecycle.md) |
| Client modes, Asset client and synchronization | [SDK](topics/sdk.md) |
| Activity and movement records | [History](topics/history.md) |
| Spatial quantities and MVP examples | [Spatial data](topics/spatial-data.md) |
| Public methods, paths, inputs and effects | [Endpoint map](api-endpoints.md) |
| Component inventory and proposed storage mapping | [Component catalog](data-components.md) |

## Decisions, delivery and evidence

| Need | Source |
| --- | --- |
| Accepted tradeoffs and their rationale | [Architectural decisions](adr/README.md) |
| Technology and generation tools | [Stack decision](adr/0016-use-go-sqlite-and-openapi-tooling.md) |
| Containers and deployment boundaries | [Deployment decision](adr/0017-deploy-core-and-plugins-as-docker-containers.md) |
| Deployment clocks and original event times | [Clock decision](adr/0029-use-deployment-clocks-and-preserve-event-times.md) |
| Required workflows, failure scenarios and completion evidence | [Testing strategy](testing-strategy.md) |
| Delivered foundation and per-ticket qualification | [Slice 0 coverage](../tests/contract/README.md) |
| Focused Plugin component and its qualification limits | [Plugin integration](../Atlas%20Core/plugins/README.md) |
| Earlier executable generation research | [Toolchain proof](research/atlas-reassessment/13-protocol-toolchain-proof.md) |

## Develop with PStack

[Contributing](../CONTRIBUTING.md) defines the contribution lifecycle. The [PStack profile](development/README.md) routes tasks through the installed skills, and [verification](development/verification.md) connects them to Atlas's actual checks. A standing coordinator uses the [project brief](development/menu.md) and [standing orders](development/house-rules.md). [Work tracking](development/work-tracking.md) connects GitHub specifications to bounded units of work.

## Research and history

The [decision log](planning-reconciliation.md) records dated authorizations, including superseded decisions. It is historical evidence, not current-rule authority. The [Modernization comparison](architecture/modernization-differences.md) records confirmed differences and points to successor authority. The [source reference](atlas-modernization-reference.md), [reassessment research](research/atlas-reassessment/README.md) and [review retrospective](research/pr-review-retrospective/README.md) preserve evidence and alternatives; a research recommendation is not an implementation commitment.
