import eel
import yaml
import datetime
import os

eel.init('www')
SETTINGS_FILE_PATH = os.path.join(os.path.expanduser('~'), "settings.yaml")

@eel.expose
def get_user_name():
    pass

@eel.expose
def is_user_set():
    data = read_settings()
    is_min_date = data["DOB"] == datetime.date(datetime.MINYEAR, 1, 1) 
    return not is_min_date


@eel.expose
def write_settings(data):
    with open(SETTINGS_FILE_PATH, 'w') as file:
        data = yaml.dump(data, file)

@eel.expose
def read_settings():
    try:
        with open(SETTINGS_FILE_PATH) as file:
            data = yaml.load(file, Loader=yaml.FullLoader)
            return data
    except FileNotFoundError:
        write_settings({
            "Name": "",
            "DOB": datetime.date(datetime.MINYEAR, 1, 1),
            "Colors": ""
            })


@eel.expose
def update_dob(name, dob, color_scheme):
    write_settings({"Name": name, "DOB": dob, "Colors": color_scheme})

@eel.expose
def get_weeks_done():
    data = read_settings()
    start = data["DOB"]
    start = datetime.date.fromisoformat(start)
    now = datetime.date.today()
    days = now - start
    return int(days.days/7)

read_settings()
eel.start('index.html', size=(850, 600))
