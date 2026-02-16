import { useWebSocket } from "./hooks/useWebSocket";
import StatsBar from "./components/StatsBar";
import EventsTable from "./components/EventsTable";
import "./App.css";

function App() {
  const { events, isConnected, error, clearEvents } = useWebSocket();

  return (
    <div className="app">
      <header className="app-header">
        <h1>Real-Time Event Dashboard</h1>
        {error && <div className="error-banner">{error}</div>}
      </header>
      <main className="app-main">
        <StatsBar
          events={events}
          isConnected={isConnected}
          onClear={clearEvents}
        />
        <EventsTable events={events} />
      </main>
    </div>
  );
}

export default App;
