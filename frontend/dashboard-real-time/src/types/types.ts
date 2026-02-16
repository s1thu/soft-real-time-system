export interface RealTimeEvent {
  id: string;
  created_at: string;
  deadline_ms: number;
  status: string;
  processed_at?: string;
}
