interface StatusLog {
  id: string;
  prevStatus: string;
  newStatus: string;
  createdAt: Date;
}

export default StatusLog;