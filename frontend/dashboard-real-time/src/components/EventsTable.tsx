import type { FC } from "react";
import type { RealTimeEvent } from "../types/types";

interface EventsTableProps {
  events: RealTimeEvent[];
}

export const EventsTable: FC<EventsTableProps> = ({ events }) => {
  const formatDate = (date: string | undefined) => {
    if (!date) return "-";
    return new Date(date).toLocaleTimeString();
  };

  return (
    <div className="table-container">
      <table className="events-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Status</th>
            <th>Created At</th>
            <th>Processed At</th>
            <th>Deadline (ms)</th>
          </tr>
        </thead>
        <tbody>
          {events.length === 0 ? (
            <tr>
              <td colSpan={5} className="empty-state">
                Waiting for events...
              </td>
            </tr>
          ) : (
            events.map((e, i) => (
              <tr
                key={e.id + i}
                className={e.status === "late" ? "late-row" : ""}
              >
                <td className="id-cell">{e.id.slice(0, 8)}</td>
                <td>
                  <span className={`status-badge ${e.status}`}>{e.status}</span>
                </td>
                <td>{formatDate(e.created_at)}</td>
                <td>{formatDate(e.processed_at)}</td>
                <td>{e.deadline_ms / 1000000}</td>
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
};

export default EventsTable;
