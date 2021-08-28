from dataclasses import dataclass
from datetime import datetime


@dataclass
class TeamScheduleEntry:
    date: datetime
    team_members: str
    notes: str
