---
name: verify-atlas
description: Verify an Atlas Core change against its exact revision, authoritative requirements and real boundaries. Use for PStack verification, documentation review or a pre-PR handoff in this repository.
---

# Verify Atlas

Read [AGENTS.md](../../../AGENTS.md) and the complete [verification procedure](../../../docs/development/verification.md). That procedure owns command selection and evidence requirements; this skill is its PStack entry point.

Classify the complete diff before running anything. For documentation-only work, check scope, links and affected authority without claiming an application run. For implementation work, read the [testing strategy](../../../docs/testing-strategy.md) and all [coding standards](../../../CODING_STANDARDS.md), map applicable guarantees, and run the required repository verifier plus relevant real-boundary evidence.

Bind the report to the actual base and head. State what ran, what passed or failed, the reviewer independence actually available and missing evidence. Keep accepted design, tested fixtures and delivered operational behavior distinct. Do not waive a check, invent a result or publish/merge as a side effect of verification.
