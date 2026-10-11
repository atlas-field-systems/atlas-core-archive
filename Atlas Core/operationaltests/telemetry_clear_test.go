package operationaltests_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/atlas-field-systems/atlas-core/generated/protocol"
	"github.com/atlas-field-systems/atlas-core/systemoperations"
	"github.com/google/uuid"
)

// The real HTTPS fixtures qualify removal, original evidence and per-unit
// ordering together. The SDK/container workflow exercises the matching surface.
func TestTelemetryClearsDoNotCreateMovementAndRetainOrdering(t *testing.T) {
	routes := []struct{ name, method, kind, suffix string }{
		{"entity", "PATCH", "entity_report", ""},
		{"checkin", "POST", "checkin", "/checkin"},
	}
	for _, route := range routes {
		for _, initial := range []string{"absent", "populated"} {
			for _, removal := range []string{"component", "fields"} {
				t.Run(route.name+"/"+initial+"/"+removal, func(t *testing.T) {
					f := newFixture(t, systemoperations.Options{})
					a, _ := f.register()
					f.checkin(a)
					a.challenge = ""
					path := "/entities/" + a.id + route.suffix
					measurements := map[string]interface{}{"position": map[string]float64{"latitude": 10, "longitude": 20}, "speed_mps": 2}
					body := func(sequence string, telemetry interface{}, kind string, origin interface{}) map[string]interface{} {
						return f.reportBody(a, route.kind, a.id, sequence, map[string]interface{}{"components": map[string]interface{}{"telemetry": telemetry}}, false, kind, origin, nil, "2026-10-10T00:00:00.500+00:00")
					}
					if initial == "populated" {
						f.request(route.method, path, a.secret, body("2", measurements, "current", nil), 200)
					}
					original := readTelemetryClearHistory(f, a)
					var clear interface{}
					if removal == "fields" {
						clear = map[string]interface{}{"position": nil, "speed_mps": nil, "altitude": nil, "heading_deg": nil}
					}
					assertCleared := func(entity protocol.Asset) {
						t.Helper()
						telemetry := entity.Components.Telemetry
						if removal == "component" {
							if telemetry != nil {
								t.Fatal("removed telemetry component was restored")
							}
						} else if telemetry == nil || !telemetry.Position.IsNull() || !telemetry.SpeedMps.IsNull() || !telemetry.Altitude.IsNull() || !telemetry.HeadingDeg.IsNull() {
							t.Fatal("explicit nullable field clears were not retained")
						}
					}
					assertNoMovement := func(response protocol.EntityReportResponse) {
						t.Helper()
						if len(response.Data.Acceptance.MovementSampleIds) != 0 || !reflect.DeepEqual(readTelemetryClearHistory(f, a), original) {
							t.Fatal("telemetry removal created movement or changed retained measurements")
						}
						assertCleared(response.Data.Entity)
					}
					clearBody := body("10", clear, "current", nil)
					cleared := decode[protocol.EntityReportResponse](t, f.request(route.method, path, a.secret, clearBody, 200))
					assertNoMovement(cleared)
					for _, phase := range []string{"same_run", "after_restart"} {
						if phase == "after_restart" {
							f.restart(systemoperations.Options{})
						}
						before, err := f.core.InspectEvidence(context.Background())
						if err != nil {
							t.Fatal(err)
						}
						replay := decode[protocol.EntityReportResponse](t, f.request(route.method, path, a.secret, clearBody, 200))
						assertNoMovement(replay)
						after, err := f.core.InspectEvidence(context.Background())
						if err != nil {
							t.Fatal(err)
						}
						if before != after || replay.Data.Acceptance.Disposition != "duplicate" || replay.Data.Acceptance.ContactRefreshed {
							t.Fatalf("%s removal retry repeated committed effects", phase)
						}
					}
					origin := map[string]string{"process_generation": "1", "sequence": "10"}
					recovered := decode[protocol.EntityReportResponse](t, f.request(route.method, path, a.secret, body("11", clear, "historical", origin), 200))
					assertNoMovement(recovered)
					if recovered.Data.Acceptance.ContactRefreshed {
						t.Fatal("recovered clear established fresh Contact")
					}
					before, err := f.core.InspectEvidence(context.Background())
					if err != nil {
						t.Fatal(err)
					}
					conflict := decode[protocol.Error](t, f.request(route.method, path, a.secret, body("12", measurements, "historical", origin), 409))
					after, err := f.core.InspectEvidence(context.Background())
					if err != nil {
						t.Fatal(err)
					}
					if conflict.Error.Code != "evidence_identity_conflict" || before != after {
						t.Fatal("changed recovery facts bypassed retained clear evidence or committed effects")
					}
					// A distinct older measurement can enter history, but not restore
					// a cleared current quantity. The rejected sequence is reusable.
					older := map[string]string{"process_generation": "1", "sequence": "9"}
					late := decode[protocol.EntityReportResponse](t, f.request(route.method, path, a.secret, body("12", measurements, "historical", older), 200))
					assertCleared(late.Data.Entity)
					if len(late.Data.Acceptance.MovementSampleIds) != 1 || len(readTelemetryClearHistory(f, a)) != len(original)+1 {
						t.Fatal("valid older measurement was lost with the clear")
					}
					newer := decode[protocol.EntityReportResponse](t, f.request(route.method, path, a.secret, body("13", measurements, "current", nil), 200))
					if newer.Data.Entity.Components.Telemetry == nil || newer.Data.Entity.Components.Telemetry.Position.IsNull() || newer.Data.Entity.Components.Telemetry.SpeedMps.GetOrEmpty() != 2 || len(newer.Data.Acceptance.MovementSampleIds) != 1 {
						t.Fatal("ordering tombstone prevented a genuinely newer measurement")
					}
				})
			}
		}
	}
}

func TestUnknownOriginClearRetainsConflictEvidenceWithoutMovement(t *testing.T) {
	f := newFixture(t, systemoperations.Options{})
	a, _ := f.register()
	f.checkin(a)
	path := "/entities/" + a.id
	seed := f.reportBody(a, "entity_report", a.id, "2", map[string]interface{}{"components": map[string]interface{}{"telemetry": map[string]float64{"speed_mps": 2}}}, false, "current", nil, nil, "2026-10-10T00:00:00Z")
	current := decode[protocol.EntityReportResponse](t, f.request("PATCH", path, a.secret, seed, 200))
	original := readTelemetryClearHistory(f, a)
	retained := uuid.NewString()
	for _, sequence := range []string{"3", "4"} {
		if sequence == "4" {
			f.restart(systemoperations.Options{})
		}
		clear := f.reportBody(a, "entity_report", a.id, sequence, map[string]interface{}{"components": map[string]interface{}{"telemetry": nil}}, false, "historical", nil, retained, nil)
		result := decode[protocol.EntityReportResponse](t, f.request("PATCH", path, a.secret, clear, 200))
		if len(result.Data.Acceptance.MovementSampleIds) != 0 || result.Data.Acceptance.ContactRefreshed || !reflect.DeepEqual(result.Data.Entity.Components.Telemetry, current.Data.Entity.Components.Telemetry) || !reflect.DeepEqual(readTelemetryClearHistory(f, a), original) {
			t.Fatal("unknown-origin removal changed current telemetry, Contact or movement")
		}
	}
	before, err := f.core.InspectEvidence(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	altered := f.reportBody(a, "entity_report", a.id, "5", map[string]interface{}{"components": map[string]interface{}{"telemetry": map[string]float64{"speed_mps": 3}}}, false, "historical", nil, retained, nil)
	conflict := decode[protocol.Error](t, f.request("PATCH", path, a.secret, altered, 409))
	after, err := f.core.InspectEvidence(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if conflict.Error.Code != "evidence_identity_conflict" || before != after {
		t.Fatal("unknown-origin clear forgot its original facts")
	}
}

func TestMixedClearCapturesOnlySuppliedMeasurementsIncludingZero(t *testing.T) {
	f := newFixture(t, systemoperations.Options{})
	a, _ := f.register()
	f.checkin(a)
	body := f.reportBody(a, "entity_report", a.id, "2", map[string]interface{}{"components": map[string]interface{}{"telemetry": map[string]interface{}{"position": nil, "speed_mps": 0, "altitude": nil}}}, false, "current", nil, nil, "2026-10-10T00:00:00Z")
	result := decode[protocol.EntityReportResponse](t, f.request("PATCH", "/entities/"+a.id, a.secret, body, 200))
	samples := readTelemetryClearHistory(f, a)
	if len(result.Data.Acceptance.MovementSampleIds) != 1 || len(samples) != 1 {
		t.Fatal("mixed report did not capture exactly one movement sample")
	}
	quantities := samples[0].Quantities
	if quantities.Position != nil || quantities.Altitude != nil || quantities.SpeedMps == nil || quantities.SpeedMps.Value.IsNull() || quantities.SpeedMps.Value.GetOrEmpty() != 0 {
		t.Fatal("mixed report included cleared quantities or dropped a real zero measurement")
	}
	if !reflect.DeepEqual(samples[0].MatchedQuantities, []protocol.MovementSampleMatchedQuantities{"speed_mps"}) {
		t.Fatal("history matched quantities included a removal")
	}
}

func readTelemetryClearHistory(f *fixture, a *asset) []protocol.MovementSample {
	f.t.Helper()
	path := "/entities/" + a.id + "/movement-history?from=2020-01-01T00:00:00Z&to=2030-01-01T00:00:00Z&time_basis=received_at"
	page := decode[protocol.MovementPageResponse](f.t, f.request("GET", path, f.admin, nil, 200))
	if !page.Data.NextCursor.IsNull() {
		f.t.Fatal("small movement fixture unexpectedly paginated")
	}
	return page.Data.Items
}
