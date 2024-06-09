from pathlib import Path
from rich.console import Console
import UnityPy
import UnityPy.config
import json

ASSETBUNDLE_DIR = "cache/plain"
IMG_DIR = "cache/img"
ESRGAN_DIR = "cache/esrgan"
DOWNLOADED_FILE_PATH = "cache/catalog_diff.json"
UnityPy.config.FALLBACK_UNITY_VERSION = "2021.3.16f1"
console = Console()


scalingWaitList: list[str] = []


def info(msg: str):
    console.print(f"[bold blue]>>> [Info][/bold blue] {msg}")


def succeed(msg: str):
    console.print(f"[bold green]>>> [Succeed][/bold green] {msg}")


def error(msg: str):
    console.print(f"[bold red]>>> [Error][/bold red] {msg}")


def warn(msg: str):
    console.print(f"[bold yellow]>>> [Warning][/bold yellow] {msg}")


def unpack_to_image(asset_bytes: bytes, dest: str):
    env = UnityPy.load(asset_bytes)
    for obj in env.objects:
        if obj.type.name == "Texture2D":
            try:
                data = obj.read()
                dest_path = Path(dest, data.name).with_suffix(".png")
                dest_path.parent.mkdir(exist_ok=True)
                img = data.image
                img.save(dest_path)
                if determineShouldScale(data.name):
                    scalingWaitList.append(str(dest_path))
                info(f"Converted '{data.name}' to png.")
            except:
                error(f"Failed to convert '{data.name}' to image.")


def determineShouldScale(name: str) -> bool:
    if name.startswith("image_card_full_") and not name.endswith("2"):
        return True
    return False


def unpack_action(diff_dict: dict[str, any]):
    for entry in diff_dict:
        try:
            raw = Path(ASSETBUNDLE_DIR, entry["StrLabelCrc"] + ".assetbundle").read_bytes()
            if raw[:5] != b"Unity":
                name = entry["StrLabelCrc"]
                warn(f"'{name}' is not a unity asset, skip processing.")
                continue
            unpack_to_image(raw, IMG_DIR)
        except:
            name = entry["StrLabelCrc"]
            error(f"Failed to process '{name}'.")


def scale_with_esrgan():
    from esrgan import convert_one
    for pth in scalingWaitList:
        convert_one(
            pth,
            ESRGAN_DIR,
            extension="webp",
            to_size=True,
            c_size=(2560, 1659),
        )


def main():
    Path(IMG_DIR).mkdir(exist_ok=True)
    with open(DOWNLOADED_FILE_PATH) as fp:
        diff_dict: dict[str, any] = json.load(fp)
    diff_dict = filter(lambda x: x["ResourceType"] <= 1, diff_dict)
    unpack_action(diff_dict)
    scale_with_esrgan()


if __name__ == "__main__":
    main()
