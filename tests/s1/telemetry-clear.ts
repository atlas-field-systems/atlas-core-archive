import assert from "node:assert/strict";
import { randomUUID } from "node:crypto";
import { createAssetClient, type PreparedCheckin, type PreparedEntityReport } from "../../Atlas SDK/src/index.js";
import { accepted, createSDKAsset, localAction, type S1Config } from "../../examples/s1/workflow.js";

export async function runTelemetryClearWorkflow(config: S1Config) {
  const { assetId, runtime, client, registered } = await createSDKAsset(config);
  accepted(registered);
  accepted(await client.discover());
  const reporter = createAssetClient({ client, assetId, signer: runtime.signer(), authorizer: runtime.authorizer() });
  const generatedAt = "2026-10-10T00:00:00.500+00:00";
  async function submitReport(descriptor: PreparedCheckin | PreparedEntityReport) {
    const pendingId = await runtime.retain(descriptor, reporter.snapshot());
    const response =
      descriptor.operation === "checkin"
        ? accepted(await client.submit(descriptor))
        : accepted(await client.submit(descriptor));
    await runtime.acknowledge(pendingId, reporter.snapshot());
    return response;
  }
  async function history() {
    const response = accepted(
      await client.getMovementHistory(assetId, {
        from: "2020-01-01T00:00:00Z",
        to: "2030-01-01T00:00:00Z",
        time_basis: "received_at",
      }),
    );
    assert.equal(response.data.next_cursor, null);
    return response.data.items;
  }
  const authority = await reporter.prepareAuthority({ expectedGeneration: "0", transferId: randomUUID(), generatedAt });
  reporter.observeAuthority(await submitReport(authority));
  const emptyClear = await reporter.prepareReport({ generatedAt, components: { telemetry: null } });
  const empty = await submitReport(emptyClear);
  assert.equal(empty.data.entity.components.telemetry, undefined);
  assert.deepEqual(empty.data.acceptance.movement_sample_ids, []);
  assert.deepEqual(await history(), []);

  const measurement = await reporter.prepareReport({
    generatedAt,
    components: { telemetry: { position: { latitude: 10, longitude: 20 }, speed_mps: 2 } },
  });
  await submitReport(measurement);
  const original = await history();
  assert.equal(original.length, 1);
  // Retain an earlier report before the clear, but delay its first transmission.
  const delayed = await reporter.prepareReport({
    generatedAt,
    components: { telemetry: { position: { latitude: 11, longitude: 21 }, speed_mps: 3 } },
  });
  const delayedId = await runtime.retain(delayed, reporter.snapshot());
  const clear = await reporter.prepareCheckin({ generatedAt, components: { telemetry: null } });
  const cleared = await submitReport(clear);
  assert.equal(cleared.data.entity.components.telemetry, undefined);
  assert.deepEqual(cleared.data.acceptance.movement_sample_ids, []);
  assert.deepEqual(await history(), original);
  await localAction(config, "restart");
  accepted(await client.discover());
  const duplicate = accepted(await client.submit(clear));
  assert.equal(duplicate.data.acceptance.disposition, "duplicate");
  assert.equal(duplicate.data.acceptance.contact_refreshed, false);
  assert.equal(duplicate.data.entity.components.telemetry, undefined);
  assert.deepEqual(duplicate.data.acceptance.movement_sample_ids, []);
  assert.deepEqual(await history(), original);

  const late = accepted(await client.submit(delayed));
  await runtime.acknowledge(delayedId, reporter.snapshot());
  assert.equal(late.data.entity.components.telemetry, undefined, "an older report cannot restore cleared telemetry");
  assert.equal(late.data.acceptance.movement_sample_ids.length, 1);
  const withLateMeasurement = await history();
  assert.equal(withLateMeasurement.length, 2);
  const recovered = await reporter.prepareReport({
    generatedAt,
    components: { telemetry: null },
    evidence: {
      kind: "historical",
      origin: {
        process_generation: clear.body.report_context.process_generation,
        sequence: clear.body.report_context.sequence,
      },
      retainedEvidenceId: null,
    },
  });
  const recovery = await submitReport(recovered);
  assert.deepEqual(recovery.data.acceptance.movement_sample_ids, []);
  assert.equal(recovery.data.acceptance.contact_refreshed, false);
  assert.deepEqual(await history(), withLateMeasurement);

  const mixed = await reporter.prepareReport({
    generatedAt,
    components: { telemetry: { position: null, speed_mps: 0, altitude: null } },
  });
  assert.equal((await submitReport(mixed)).data.acceptance.movement_sample_ids.length, 1);
  const withZero = await history();
  assert.equal(withZero.length, 3);
  const zero = withZero.at(-1);
  assert(zero);
  assert.equal(zero.quantities.position, undefined);
  assert.equal(zero.quantities.altitude, undefined);
  assert.equal(zero.quantities.speed_mps?.value, 0);
  assert.deepEqual(zero.matched_quantities, ["speed_mps"]);
  const fieldClear = await reporter.prepareReport({ generatedAt, components: { telemetry: { speed_mps: null } } });
  const removedSpeed = await submitReport(fieldClear);
  assert.equal(removedSpeed.data.entity.components.telemetry?.speed_mps, null);
  assert.deepEqual(removedSpeed.data.acceptance.movement_sample_ids, []);
  assert.deepEqual(await history(), withZero);
  console.log("PASS real container SDK telemetry clears: no invented samples, retained history, Restart and ordering");
}
