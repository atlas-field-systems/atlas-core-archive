# PStack engineering workflow

PStack-T3 is Atlas's engineering workflow. `poteto-mode` drives a task from investigation to a verified result. `brigade` coordinates an authorized stream of work. Atlas's [documentation ownership](../documentation.md), [contribution policy](../../CONTRIBUTING.md), [coding standards](../../CODING_STANDARDS.md) and [testing strategy](../testing-strategy.md) remain binding in both modes.

## Set up the runtime

Install the skills from [creedants/pstack-t3](https://github.com/creedants/pstack-t3). Use its [setup instructions](https://github.com/creedants/pstack-t3/blob/4ee9e3c73dcca16278ace6879164f0314e6b161d/README.md#quick-start) and check the installed revision's T3 prerequisites. This profile was checked against revision `4ee9e3c73dcca16278ace6879164f0314e6b161d`.

The upstream installer links the skills into the supported providers. Keep that checkout available, run its installer doctor, and open a new T3 thread. Run `$setup-pstack` to resolve model roles against the live catalog and your available providers. Do not copy a model name out of a guide and assume it is available. Skill installation and local model configuration are machine setup, not part of the Atlas source tree.

The local [verify-atlas skill](../../.agents/skills/verify-atlas/SKILL.md) connects PStack to Atlas's verification procedure. `AGENTS.md` points to it explicitly; read it even when the harness does not automatically discover repository-local skills. Do not claim that the runtime, coordinator or a review panel is installed or running merely because these Markdown files exist.

## One task

Start with a goal and a checkable outcome:

```text
$poteto-mode Implement the authorized Atlas issue in this checkout.
Read its comments, dependencies and required baseline. Follow AGENTS.md.
Keep the change within the issue's scope and prove the required workflow.
Open a PR only if this request explicitly authorizes publication. Do not merge.
```

Let the playbook choose the supporting skills. Use `how` to establish what exists, `architect` for consequential interface design, `interrogate` for adversarial review, and `swarm` or `arena` when parallel investigation or competing designs earn their cost. For documents and PR prose, use the installed writing skills. Read a skill before claiming to apply it.

Keep the visible task list tied to the selected playbook. Mark a skipped step with its reason. For a mechanical documentation edit, use proportional inspection and verification rather than manufacturing a multi-agent program.

## A standing coordinator

A coordinator uses the committed [project brief](menu.md) and [standing orders](house-rules.md) as inputs to its local `menu.md` and `house-rules.md`. They describe Atlas's purpose and constraints; the installed Brigade skill owns the live store and command syntax. Resolve the task's specification and baseline before creating a worker brief.

```text
$brigade Open an Atlas Core coordinator for the work queue I have authorized.
Use docs/development/menu.md and docs/development/house-rules.md.
Use milestone reporting and at most two writing or reviewing units in flight.
Do not enable unsolicited issue intake or schedules without authorization.
Publish only explicitly authorized PRs, using human landing mode. Do not merge.
```

Opening this guide does not start that coordinator or grant standing authorization. Existing PRs are not a new source of feature requests. Do not take over an active branch or duplicate an in-flight implementation without an explicit assignment.

The coordinator owns intake, dependency order, briefs and evidence review. Workers own implementation in isolated worktrees. A verifier checks the exact candidate revision, preferably on a different model family. The landing queue owns authorized integration; workers never write trunk. Check path leases before parallel work, and route shared documentation edits through one owner rather than allowing competing policy changes.

Use the [work-tracking brief](work-tracking.md#worker-brief) to carry the goal, source requirements, write scope, acceptance criteria, verification commands and publication authority. After a rejected review, consolidate the original request and new findings for a fresh worker rather than accumulating an unbounded fix-and-test loop. Three send-backs trigger the installed Brigade decision gate; they do not authorize weakening a requirement.

## Review and completion

Verify the artifact, not the confidence of the report. A green structural suite does not prove an operational workflow. Use [verification](verification.md), and assess correctness, scope and maintainability against the same exact revision. Prove or dismiss automated review findings instead of treating every comment as a mandatory new abstraction or test.

A report identifies the branch and head, implemented outcome, requirements covered, commands and results, reviewer independence, and remaining limits. A changed head invalidates any affected review or test evidence. A partial result stays partial even when its subset passes.

Use PStack's `human` landing mode only after PR publication is authorized. It opens a PR but does not merge it. A request to review, execute a task or run a coordinator is not blanket permission to merge, deploy, force-push or delete data. The [contribution policy](../../CONTRIBUTING.md#review-and-publication) owns these gates.

## Working records and durable knowledge

| Record | Home | Git policy |
| --- | --- | --- |
| Product contracts, accepted tradeoffs and terminology | Topic pages, ADRs and glossary | Committed and reviewed |
| Requested implementation and acceptance criteria | GitHub issue, linked to governing docs | Tracker record; do not maintain a second editable copy |
| Coordinator state, leases, reports and queue state | Installed Brigade/landing runtime store | Local runtime state, not authoritative docs |
| Task decision trail | `decisions.tsv` or `.audit/<task-slug>.tsv` | Local by default; publish only deliberately reviewed evidence |
| Skill installation and model-role settings | Installed PStack runtime and its user configuration | Machine-local |

Before closing work, promote lasting authorized decisions to their [documentation home](../documentation.md#recording-changes). Keep the working trail short and evidence-linked. Do not dump transcripts, model deliberations or per-action narration into the repository.

The repository does not automatically ignore every local PStack artifact. Keep runtime stores outside the checkout. If a task uses `.audit/` or root `decisions.tsv`, add those paths to the checkout's local Git exclude file before generating them, and inspect the staged diff before committing.

## Runtime references

The pinned upstream [task workflow](https://github.com/creedants/pstack-t3/blob/4ee9e3c73dcca16278ace6879164f0314e6b161d/skills/poteto-mode/SKILL.md), [Brigade coordinator](https://github.com/creedants/pstack-t3/blob/4ee9e3c73dcca16278ace6879164f0314e6b161d/skills/brigade/SKILL.md) and [decision-trail skill](https://github.com/creedants/pstack-t3/blob/4ee9e3c73dcca16278ace6879164f0314e6b161d/skills/show-me-your-work/SKILL.md) define runtime behavior. Keep this repository's profile small instead of vendoring another copy of the skill pack. Runtime upgrades do not silently change Atlas's acceptance or publication policy.
