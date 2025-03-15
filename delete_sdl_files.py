# for deleting error copied SDL files

from pathlib import Path
import os

d1 = Path(r"C:\Users\fts-guest-05600\Downloads\SDL2-2.32.2\x86_64-w64-mingw32")

d2 = Path(r"C:\TDM-GCC-64\x86_64-w64-mingw32")


for d in os.listdir(d1 / "lib"):
    pd = d1 / "lib" / d
    if pd.is_file():
        pd2 = d2 / "lib" / d
        if pd2.exists():
            pd2.unlink()
            print("deleted: ", pd2)
