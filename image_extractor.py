from pathlib import Path
from image_converter import convert_one
import UnityPy
import json

AST_SRC_DIR = "cache/plain"
IMG_DST_DIR = "cache/image"
ESRGAN_DST_DIR = "cache/esrgan"
SIGNATURE = b"Unity"


scalingWaitList: list[str] = []


def scale_with_esrgan():
    for entry in scalingWaitList:
        pth = Path(entry)
        convert_one(
            str(pth), ESRGAN_DST_DIR, to_size=True, c_size=(2560, 1659)
        )  # 1,659.259259...


def determineShouldScale(name: str) -> bool:
    if name.startswith("image_card_full_") and not name.endswith("2"):
        return True
    return False


def unpack_to_image(asset_bytes: bytes, dest_dir: str):
    if asset_bytes[:5] != SIGNATURE:
        print(f"File: '{data.name}'.")
        print(f"Signature mismatched, expect: {SIGNATURE}, given {asset_bytes[:5]}.")
        return
    env = UnityPy.load(asset_bytes)
    for obj in env.objects:
        if obj.type.name == "Texture2D":
            try:
                data = obj.read()
                dest_path = Path(dest_dir, data.name).with_suffix(".png")
                # if dest_path.exists():
                #     print(f"'{dest_path}' already exists.")
                #     continue
                dest_path.parent.mkdir(exist_ok=True)
                img = data.image
                img.save(dest_path)
                if determineShouldScale(data.name):
                    scalingWaitList.append(f"{dest_dir}/{data.name}.png")
                print(f"Successfully converted '{data.name}' to image.")
            except:
                print(f"Failed to convert '{data.name}' to image.")


def main():
    with open("cache/catalog_diff.json") as fp:
        entries = json.load(fp)
    entries = filter(lambda x: x["ResourceType"] <= 1, entries)
    for entry in entries:
        pth = AST_SRC_DIR + "/" + entry["StrLabelCrc"] + ".assetbundle"
        pth = Path(pth)
        contents = pth.read_bytes()
        unpack_to_image(contents, IMG_DST_DIR)

    # for pth in Path(AST_SRC_DIR).glob("*.assetbundle"):
    #     contents = pth.read_bytes()
    #     unpack_to_image(contents, IMG_DST_DIR)
    print("Convertion completed.")
    print("Start scaling images with esrgan.")
    scale_with_esrgan()
    print("Scaling completed.")

    print("All tasks completed.")


if __name__ == "__main__":
    main()
