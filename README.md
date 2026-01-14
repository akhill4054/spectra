# Spectra

Spectra is an **in-progress, privacy-preserving backend system** for computing real-time audience metrics for digital signage using face detection events from edge devices.

---

## Project Status

🚧 **Actively in progress**

Spectra is currently being built as a learning exercise focused on streaming-style data processing in Go. It may evolve into a full-fledged face analytics engine for real-world signage systems over time.

---

## MVP (Scope)

The initial MVP focuses strictly on:

- [ ] Ingest face detection events from edge devices (camera)
- [ ] Maintain short-lived viewer presence state
- [ ] Compute basic (per screen) metrics:
  - [ ] Concurrent viewers
  - [ ] Dwell time
- [ ] Persist aggregated metrics (as commit-log)
- [ ] Expose read-only analytics APIs (debug, tooling)

---

## Design Principles (MVP)

- **Event-first architecture** - no image or video handling
- **Privacy-preserving by design** - no identity persistence
- **Failure-aware** - tolerates duplicate, late, and out-of-order events
- **Bounded concurrency** - backpressure over unbounded growth (slow down or reject events)

---

## Non-Goals

- Face recognition or biometric identification
- Image or video storage
- Frontend dashboards
- Model training or ML experimentation

---

## License

MIT
