import os
from datetime import datetime

import click

from data_classes.team_management import TeamScheduleEntry
from utils.async_helpers import coro
import gspread

print(os.environ.items())
#
#
# gc = gspread.service_account_from_dict({
#     'type': os.getenv('GOOGLE_CREDENTIALS_TYPE'),
#     'project_id': os.getenv('GOOGLE_PROJECT_ID'),
#     'private_key_id': os.getenv('GOOGLE_PRIVATE_KEY_ID'),
#     'private_key': os.getenv('GOOGLE_PRIVATE_KEY').replace("\\n", "\n").replace("\"", ""),
#     'client_email': os.getenv('GOOGLE_CLIENT_EMAIL'),
#     'client_id': os.getenv('GOOGLE_CLIENT_ID'),
#     'auth_uri': os.getenv('GOOGLE_AUTH_URI'),
#     'token_uri': os.getenv('GOOGLE_TOKEN_URI'),
#     'auth_provider_x509_cert_url': os.getenv('GOOGLE_AUTH_PROVIDER_X509_CERT_URL'),
#     'client_x509_cert_url': os.getenv('GOOGLE_CLIENT_X509_CERT_URL'),
# })
#
# SHEET_ID = os.getenv('SCHEDULE_SHEET_ID')
#
#
# @click.group()
# def cli():
#     pass
#
#
# @cli.command()
# @click.option("--team", "-t",
#               type=click.Choice(['worship']),
#               required=True)
# @coro
# async def who_is_serving(team: str):
#     spreadsheet = gc.open_by_key(SHEET_ID)
#     matching_sheet = [sheet for sheet in spreadsheet.worksheets() if sheet.title.lower() == team.lower()]
#     if len(matching_sheet) == 0:
#         print(f'Sorry about that! I had trouble finding the schedule for the {team} team...')
#         return
#
#     team_sheet = matching_sheet[0]
#     entries = [TeamScheduleEntry(date=datetime.strptime(entry.get('Date'), '%m/%d/%Y'),
#                                  team_members=entry.get('Team Members'),
#                                  notes=entry.get('Notes')) for entry in team_sheet.get_all_records()]
#     future_entries = [entry for entry in entries if entry.date >= datetime.now()]
#     if len(future_entries) == 0:
#         print(f'Sorry about that! The spreadsheet needs to be updated with new schedules.')
#         return
#
#     serving = min(future_entries, key=lambda x: x.date)
#     print(
#         f'_Here\'s who\'s serving this Sunday, {serving.date.strftime("%b %d")}:_\n\n*{with_slack_mentions(serving.team_members)}*')
#
#
# @cli.command()
# @click.option("--team", "-t",
#               type=click.Choice(['worship']),
#               required=True)
# @coro
# async def get_schedule(team: str):
#     spreadsheet = gc.open_by_key(SHEET_ID)
#     matching_sheet = [sheet for sheet in spreadsheet.worksheets() if sheet.title.lower() == team.lower()]
#     if len(matching_sheet) == 0:
#         print(f'Sorry about that! I had trouble finding the schedule for the {team} team...')
#         return
#
#     team_sheet = matching_sheet[0]
#     entries = [TeamScheduleEntry(date=datetime.strptime(entry.get('Date'), '%m/%d/%Y'),
#                                  team_members=entry.get('Team Members'),
#                                  notes=entry.get('Notes')) for entry in team_sheet.get_all_records()]
#     future_entries = [entry for entry in entries if entry.date >= datetime.now()]
#     if len(future_entries) == 0:
#         print(f'Sorry about that! The spreadsheet needs to be updated with new schedules.')
#         return
#
#     schedule_text = '\n'.join(
#         [f'{entry.date.strftime("%b %d")}: *{with_slack_mentions(entry.team_members)}*' for entry in
#          future_entries])
#     sheet_text = f'_Feel free to make changes on the google sheet:_ https://docs.google.com/spreadsheets/d/{SHEET_ID}'
#     print(
#         f'_Here\'s the upcoming schedule for the {team} team:_\n\n{schedule_text}\n\n{sheet_text}')
#
#
# def with_slack_mentions(text: str) -> str:
#     slack_user_map = {
#         'peppy': 'U010DDYSGDN',
#         'shalom': 'U0106QYAU03',
#         'ida': 'U010DGBGMUG',
#         'angelo': 'U010HN4HH60',
#         'angel': 'U010HN4HH60',
#         'bikki': 'U010HUMG7TQ',
#         'caleb': 'U0105LQ01A5',
#         'jackie': 'U010KU8141M',
#         'benu': 'U01JTK7SEM9'
#     }
#
#     for name, slack_user_id in slack_user_map.items():
#         text = text.replace(name, f'<@{slack_user_id}>')
#         text = text.replace(name.capitalize(), f'<@{slack_user_id}>')
#
#     return text
#
#
# if __name__ == "__main__":
#     cli()
