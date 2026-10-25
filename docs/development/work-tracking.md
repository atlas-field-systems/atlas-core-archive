# Work tracking

GitHub Issues hold Atlas's assigned specifications and implementation work. Project documentation owns enduring rules under the [documentation guide](../documentation.md). A PStack coordinator's local queue is execution bookkeeping, not another specification authority.

## Intake and state

Read the whole issue, its comments, labels, dependencies and source documents. Use an authenticated GitHub client or the connected repository tools. Resolve a bare issue/PR number before acting; GitHub shares their number space. Use canonical issue and PR URLs in coordinator records.

| Existing tracker label | Meaning |
| --- | --- |
| `needs-triage` | Scope and relevance need assessment |
| `needs-info` | Required information is missing |
| `ready-for-agent` | Work is ready for an agent within the issue's constraints |
| `ready-for-human` | A human decision or action is required |
| `wontfix` | The work is explicitly declined |

Read labels actually present before applying them. A label does not supply authority to change requirements, publish a PR or merge. Do not rename labels or rewrite unrelated issues merely to match local runtime terminology.

A coordinator ingests only its authorized request source and deduplicates by the canonical tracker URL. PRs are review and delivery surfaces, not an unsolicited feature-request feed. Existing work in progress remains owned by its current branch and assignee unless reassigned.

## Dependencies and specifications

Use native GitHub issue dependencies and sub-issues where available. Otherwise record explicit `Blocked by` and `Part of` links. An implementation unit starts only when its required inputs are available. Closed dependencies are not proof of a technical guarantee; read their disposition and evidence.

A specification records the problem, intended outcome, source requirements, scope, exclusions and acceptance evidence. Link to durable rules instead of copying every topic page. A bounded unit can point directly to an existing issue; do not create a hierarchy of empty tickets for a small edit. Research alone does not require an issue unless assigned as tracked work.

Record a required documentation or implementation baseline by commit or prerequisite PR. If the baseline is not on the target branch, make that dependency explicit. Do not copy an unmerged implementation into a documentation-only branch to satisfy it.

## Worker brief

Give each worker enough context to act without access to another agent's conversation:

```text
Goal: the authorized observable outcome.
Source: canonical issue/request and relevant decisions.
Base: exact starting revision and prerequisite PRs.
Scope: writable paths, exclusions and exclusive worktree.
Contract: governing rules, trust model and important limits.
Acceptance: checkable success and failure conditions.
Verify: actual commands and relevant scenario/skill pointers.
Authority: permitted edits and whether PR publication or merging is authorized.
Report: branch, head, outcome, commands/results and missing evidence.
Standing orders: docs/development/house-rules.md, carried into the brief.
```

Use the installed Brigade brief interface rather than inventing another queue schema. Transfer upstream reports and decisions when a later unit depends on them. A request to change shared documentation belongs to one designated writer.

## Review and completion

Review the exact commit against the specification and [coding standards](../../CODING_STANDARDS.md) independently. Record acceptance evidence in the PR, with links to committed artifacts or accessible CI outputs. Do not rely on local log paths as the reader's only evidence.

A partial implementation references its issue without closing the full specification. File genuine out-of-scope follow-ups separately when authorized; do not expand the current PR to absorb every suggestion. Close work only when its acceptance conditions and applicable checks are satisfied. Promote lasting decisions into documentation before the handoff.
