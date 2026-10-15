# Atlas verification

The [testing strategy](../testing-strategy.md) owns scenario selection, real-boundary evidence and completion criteria. [Coding standards](../../CODING_STANDARDS.md) owns engineering checks. This page maps PStack's verification step to those requirements without replacing them.

## Establish the candidate

Read the request, source requirements and required baseline. Record the comparison base and candidate head. Inspect the changed paths before choosing checks. A documentation-only claim requires that the complete PR diff, including prerequisite commits, changes only documentation.

Do not infer delivered behavior from a spec, README, generated type or an agent's report. The [implementation sequence](../architecture/implementation-sequence.md) and the relevant qualification records identify what the current checkout actually establishes.

## Documentation-only work

Check both the documentation meaning and the references:

1. Inspect the complete base-to-head changed-file list and diff. Confirm no application source, Protocol schema, test source, generated output, dependency, Git configuration or CI configuration changed. Markdown skill definitions are documentation; embedding an instruction does not execute it.
2. Resolve relative links and relevant heading anchors from each changed file's new location. Check incoming links for moved or removed documents. A retained reference page must point to one canonical home and carry no competing policy.
3. Check affected glossary definitions, topic rules, ADR status, implementation status, API/component summaries and evidence claims against the accepted baseline. Keep proposals, historical decisions and delivered behavior distinct.
4. Confirm PStack prompts and standing orders preserve decision and publication authority. Check referenced skill names and runtime conventions against the installed or explicitly cited upstream revision.
5. Inspect whitespace and staged content, including local logs or configuration accidentally selected for commit. Record precisely which files and references were checked and any inaccessible evidence.

Useful repository commands:

```sh
# Replace BASE with the actual comparison revision.
git diff --name-status BASE...HEAD
git diff --check BASE...HEAD
git diff --stat BASE...HEAD
git status --short
```

For uncommitted work, inspect the working and staged diffs as well. A remote API comparison can establish path scope when no local checkout is available; it does not prove a local command ran. Documentation checks do not require an unimplemented application suite. Existing required CI checks still apply and must not be disabled.

## Code, contracts, generators or tests

Read all of `CODING_STANDARDS.md`, map the changed guarantees to applicable required scenarios, and run the repository's required checks from the root:

```sh
python3 scripts/verify.py --bootstrap
```

Use the repository's configured toolchain and inspect the current verifier before reporting what it ran. Run focused real-boundary probes for uncovered combinations as required by the testing strategy. `python3 scripts/format.py` applies the configured formatting for implementation work; review its diff rather than mixing broad formatting into an unrelated change.

Generation, build, type, lint and fixture checks are necessary within their scope. They do not by themselves qualify operational Core/SDK workflows, real container lifecycle, independently built Plugins, field hardware, physical radio capacity or released-version compatibility. Record the exact editions and scenarios exercised. Never relabel the isolated foundation as the operational implementation.

## Independent review

Use the PStack review step for a material change. Supply the exact base/head, requirements, diff and verification evidence. Prefer a reviewer on another model family. A same-family reviewer is not cross-family review; self-review is not independent review. Report an unavailable review capability instead of fabricating a pass.

Check implementation correctness separately from scope and maintainability. Treat automated findings as hypotheses with a concrete consequence to verify. Do not build a new layer or duplicate tests merely to satisfy a stylistic suggestion. Revisit a failed premise instead of carrying it through repeated fix rounds.

## Evidence and landing

Record commands actually executed, their outcomes, revision, applicable requirements and material omissions. A failure claimed to predate the change requires equivalent-base evidence. An unexplained failure remains unresolved.

The candidate must match its evidence. Recheck every guarantee invalidated by later edits or integration. Use the landing queue only under the [publication policy](../../CONTRIBUTING.md#review-and-publication). Required failing, skipped or unavailable checks leave the work unverified; a partial result must remain clearly partial. Finish documentation consistency checks even when application checks are green.
