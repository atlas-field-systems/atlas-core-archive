# Atlas Core standing orders

Carry these numbered orders into each coordinator worker or verifier brief, together with the request, required baseline and applicable source links. They bind every PStack playbook. They do not replace the linked policies.

1. Read [documentation ownership](../documentation.md), the glossary and each applicable contract before changing behavior or design. A task's required accepted baseline can be newer than trunk; never silently substitute older rules.
2. Preserve [decision authority](../documentation.md#decision-authority). The user owns observable behavior. Engineering selects internal mechanisms that preserve it. Existing authorization remains valid; missing authorization is not supplied by a review comment or implementation difficulty.
3. Work only on the assigned goal and paths. One writer owns each worktree. The coordinator manages work and evidence, not application code. Use the runtime's path leases when coordinating parallel work.
4. Use the installed PStack runtime to resolve roles and launch supported delegates. Never invent a model, tool, running coordinator, independent review or test result. Record unavailable capabilities and same-family review honestly.
5. Read all of [CODING_STANDARDS.md](../../CODING_STANDARDS.md) for code changes. Design the domain shape and important interfaces first. Prefer deleting unnecessary machinery over adding wrappers, configuration or speculative abstractions.
6. Prove the required behavior through the real boundaries under [testing strategy](../testing-strategy.md) and [verification](verification.md). No weakened checks, self-confirming fixtures or passing subset can establish full completion.
7. Treat review findings as claims to evaluate. Fix demonstrated defects within scope; justify dismissals. Repeated failures of the same premise require reassessing the design, not another layer of patches or duplicate tests. Preserve the installed runtime's review-round gates.
8. Put lasting authorized decisions in their [documentation home](../documentation.md#recording-changes). Keep local execution trails and coordinator stores separate. Preserve actual decision dates, source evidence and superseded records.
9. Open a PR only when explicitly requested, ready for review. Use `human` landing for authorized publication unless another mode is explicitly granted. Merging, deployment, shared-branch force-push and destructive data operations require their own authorization.
10. Before changing a substantial visual direction, obtain the user's choice among distinct static mocks. After approval, do not reopen routine implementation choices. Continue independent authorized work around genuine product gates.
11. Report the exact branch/head, outcome, requirements covered, commands and results, review independence and remaining limits. Check the final diff, not an earlier candidate. Do not close a specification on partial delivery.
