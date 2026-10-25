# Telemetry removal regression

This follow-up qualifies the existing [S1 reporting contract](../../docs/topics/asset-reporting.md#ordering-and-independent-effects) and [movement capture rules](../../docs/topics/history.md#capture). Explicit null removes a component or clears a nullable field; it is not a movement measurement.

| Requirement | Implementation | Regression coverage |
| --- | --- | --- |
| Clears produce no movement samples, including before any measurements exist | `entities/components.go` retains clear evidence but captures only non-null measurements | `operationaltests/telemetry_clear_test.go`: Entity patch and check-in, absent and populated telemetry, whole-component and individual-field clears |
| Removal ordering and original evidence survive retries and Restart | Existing report acceptance and ordering remain in the same commit; clear evidence has no movement association | Same direct HTTPS suite: duplicate and recovered clears, changed known/unknown-origin evidence rejected without committed effects, older measurements captured without restoring current values, newer measurements still accepted |
| Mixed reports capture only actual supplied measurements | Movement quantities exclude clears while retaining a real zero value | Direct HTTPS mixed report plus `telemetry-clear.ts` through the public SDK |
| Supported SDK and actual container preserve the same behavior | `workflow.ts` runs `runTelemetryClearWorkflow` between isolated Dataset resets | CLI Restart, retained SDK descriptors, check-in and Entity reports, delayed older report, historical clear, full history read-back |

Paths under `entities/` and `operationaltests/` above are relative to `Atlas Core/`. The direct fixtures use real Core, HTTPS and SQLite, with independently authored expectations. The SDK workflow uses the existing host manager, actual container and caller-owned reporting retention. It introduces no new simulator storage or production endpoint.

The required verifier discovers the Go regression tests and runs the added SDK workflow through the existing `TestPublicSDKAndDirectProtocol` container harness:

```sh
python3 scripts/verify.py --bootstrap
```

A successful run must be associated with the final source revision. These coverage descriptions are not a substitute for executed verification. No generated output, storage schema, existing assertion or required check is changed. Previously retained history is not rewritten.
