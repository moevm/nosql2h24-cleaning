interface StatusLog {
  id: string;
  prev_status: string;
  new_status: string;
  created_at: Date;
}

export default StatusLog;