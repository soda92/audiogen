from fastapi import FastAPI
from pathlib import Path
from fastapi.responses import RedirectResponse
from fastapi.responses import FileResponse

CURRENT = Path(__file__).resolve().parent

app = FastAPI()


@app.get("/")
def index():
    return RedirectResponse("/index.html")


@app.get("/index.html")
def index_html():
    return FileResponse(path=CURRENT.parent.joinpath("index.html"))


@app.get("/index.js")
def index_js():
    return FileResponse(path=CURRENT.parent.joinpath("index.js"))


@app.get("/play")
def play():
    p = CURRENT.parent.joinpath("sine_wave.wav")
    from sodatools import str_path
    pv = str_path(p)
    import winsound

    winsound.PlaySound(pv, winsound.SND_FILENAME)
    return ""
