import string
from pathlib import Path
import random

def get_name() -> str:
    return ''.join(random.sample(string.ascii_lowercase, random.randint(4, 20))).title()

def main():
    out = Path(__file__).parent / "data.txt"
    with out.open("w") as f:
        for _ in range(1000):
            name = f"{get_name()} {get_name()}"
            f.write(name + "\n")


if __name__ == '__main__':
    main()