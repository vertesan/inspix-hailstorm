import inference_realesrgan
import sys
import contextlib
from pathlib import Path
from PIL import Image
from rich.console import Console

console = Console()


def info(msg: str):
    console.print(f"[bold blue]>>> [Info][/bold blue] {msg}")


def succeed(msg: str):
    console.print(f"[bold green]>>> [Succeed][/bold green] {msg}")


def error(msg: str):
    console.print(f"[bold red]>>> [Error][/bold red] {msg}")


def warning(msg: str):
    console.print(f"[bold yellow]>>> [Warning][/bold yellow] {msg}")


@contextlib.contextmanager
def _redirect_argv(args: list):
    arg0 = sys.argv[0]
    args.insert(0, arg0)
    sys.argv = args


def _convert_with_esrgan(inputs: str, output: str, suffix: str, scale: str, tile: str):
    """Original Doc:
    Usage: python inference_realesrgan.py -n RealESRGAN_x4plus -i infile -o outfile [options]...

    A common command: python inference_realesrgan.py -n RealESRGAN_x4plus -i infile --outscale 3.5 --half --face_enhance

    -h                   show this help
    -i --input           Input image or folder. Default: inputs
    -o --output          Output folder. Default: results
    -n --model_name      Model name. Default: RealESRGAN_x4plus
    -s, --outscale       The final upsampling scale of the image. Default: 4
    --suffix             Suffix of the restored image. Default: out
    -t, --tile           Tile size, 0 for no tile during testing. Default: 0
    --face_enhance       Whether to use GFPGAN to enhance face. Default: False
    --half               Whether to use half precision during inference. Default: False
    --ext                Image extension. Options: auto | jpg | png, auto means using the same extension as inputs. Default: auto
    """
    arg_list = [
        "-n",
        "RealESRGAN_x4plus_anime_6B",
        "--suffix",
        suffix,
        "-s",
        scale,
        "-i",
        inputs,
        "-o",
        output,
        "--ext",
        "png",
        "-t",
        tile,
    ]
    _redirect_argv(arg_list)
    inference_realesrgan.main()


def convert_one(
    input: str,
    output: str,
    scale: str = "2",
    tile: str = "512",
    extension: str = "webp",
    suffix: str = "esrgan",
    to_size: bool = False,
    c_size: tuple = (0, 0),
):
    # input = str(Path(input).absolute())
    # output = str(Path(output).absolute())
    try:
        info(f"Start converting {input}.")
        _convert_with_esrgan(
            inputs=input,
            output=output,
            suffix=suffix,
            scale=scale,
            tile=tile,
        )
        if to_size:
            info("Begin after-converting.")
            esrgan_file = f"{output}/{Path(input).stem}_{suffix}.png"
            save_file = f"{output}/{Path(input).stem}_{suffix}.{extension}"
            img = Image.open(esrgan_file)
            img = img.resize(c_size, resample=Image.LANCZOS, reducing_gap=3)
            img.save(fp=save_file, format=extension, lossless=True, quality=80)
            if esrgan_file != save_file:
                img.close()
                Path(esrgan_file).unlink(missing_ok=True)

    except Exception as err:
        error(err)
        raise
    succeed(f"Convert {input} completed.")
