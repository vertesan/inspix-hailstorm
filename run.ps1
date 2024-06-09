New-Item -ItemType "directory" -Path "cache" -Force
New-Item -ItemType "directory" -Path "masterdata" -Force

./hailstorm.exe
if (Test-Path "cache/updated") {
  python unpack.py
}
