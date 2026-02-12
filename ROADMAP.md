# Spectra Roadmap

This roadmap is a living document that tracks the planned scope and
evolution of Spectra.

---

## MVP (Scope)

The initial MVP focuses strictly on:

-   [ ] Define atomic observation event schema
-   [ ] Ingest observation events from edge devices (camera or adapter)
-   [ ] Maintain short-lived presence state at area level
-   [ ] Support multiple cameras per area (including overlapping feeds)
-   [ ] Compute basic area-level metrics:
    -   [ ] Concurrent viewers
    -   [ ] Dwell time
-   [ ] Persist aggregated metric events (append-only commit log)
-   [ ] Expose read-only analytics APIs (debug, tooling)

---

## Design Principles (MVP)

-   **Event-first architecture** - no image or video handling
-   **Replayable pipeline** - derived metrics can be recomputed from
    stored events
-   **Privacy-preserving by design** - no identity persistence in MVP
-   **Failure-aware** - tolerates duplicate, late, and out-of-order
    events
-   **Area-level truth** - cameras are signal sources; areas are
    aggregation boundaries
-   **Bounded concurrency** - backpressure over unbounded growth (slow
    down or reject events)

