import type { FC } from "react";
import type { RealTimeEvent } from "../types/types";

interface StatsBarProps {
  events: RealTimeEvent[];
  isConnected: boolean;
  onClear: () => void;
}

const StatsBar: FC<StatsBarProps> = ({ events, isConnected, onClear }) => {
  const late = events.filter((e) => e.status === "late").length;
  const onTime = events.filter((e) => e.status === "on-time").length;

  return (
    <div className="stats-bar">
      <div className="stats-info">
        <span
          className={`connection-status ${isConnected ? "connected" : "disconnected"}`}
        >
          {isConnected ? "● Connected" : "○ Disconnected"}
        </span>
        <span className="stat-item">
          <strong>Total:</strong> {events.length}
        </span>
        <span className="stat-item on-time">
          <strong>On-time:</strong> {onTime}
        </span>
        <span className="stat-item late">
          <strong>Late:</strong> {late}
        </span>
        {events.length > 0 && (
          <span className="stat-item">
            <strong>Late Rate:</strong>{" "}
            {((late / events.length) * 100).toFixed(1)}%
          </span>
        )}
      </div>
      <button onClick={onClear} className="clear-btn">
        Clear
      </button>
    </div>
  );
};

export default StatsBar;
