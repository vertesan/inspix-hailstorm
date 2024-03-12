# AES

Padding = PKCS7 Mode = CBC (Cipher Block Chaining) KeySize = 128 BlockSize = 128

SHA256() the result of SHA256 (32 bytes) used for Key: 0~16 bytes IV: 16~32
bytes

# binWriter Rules

1. qword_7D13E80

2.

AssetInfoRow.Seed & 0x7FFFFFFFFFFFFFFF (8 bytes
Catalog.Manifest.Seed (8 bytes

3.

CRC64 calculated from assetName raw string (raw bytes?) (8 bytes
CRC64 calculated from String.Concat(UnityEngine.Application.Version(), ":", simpleResver)

4.

CRC32 calculated from "raw" raw string (raw bytes?) (4 bytes
CRC32 calculated from "android" raw string (raw bytes?) (4 bytes

5.

VLQ (Variable Length Quantity) calculated from AssetInfoRow.Size (8 bytes
VLQ (Variable Length Quantity) calculated from Catalog.Manifest.Size (8 bytes

6. 
sha256 of all previous written bytes
[:16] key
[16:] iv

7. 
aes decrypt

8. 
last step: LZ4FrameDecoder -> DecompressLZ4Bytes

# binary example

binWriter flushed result

c34ea77df4976cd8907096fa47bb97e61852305892494e3692ba0c7eb434f022c549c96cf7ca0ee1b6ba7f203b6c76e8679699ce9c44af7b1cb000173a515938c151bf0f407bd74a3743595b935578cf92bcb62dc0f01f


1 (constant arbitrary bytes 64b): C3 4E A7 7D F4 97 6C D8 90 70 96 FA 47 BB 97
E6 18 52 30 58 92 49 4E 36 92 BA 0C 7E B4 34 F0 22 C5 49 C9 6C F7 CA 0E E1 B6 BA
7F 20 3B 6C 76 E8 67 96 99 CE 9C 44 AF 7B 1C B0 00 17 3A 51 59 38

2 (Seed 8b): C1 51 BF 0F 40 7B D7 4A

3 (CRC64 appver + ":" + Manifest.SimpleResver 8b): 37 43 59 5B 93 55 78 CF

4 (CRC32 android 4b): 92 BC B6 2D

5 (vlq size variable bytes): C0 F0 1F

sha256 bdb9cc91acb1376b92e78ddc16e9e3abcd17ec779c195dbb8df3b8ffdecbef7a

key bytes: bdb9cc91acb1376b92e78ddc16e9e3ab iv bytes:
cd17ec779c195dbb8df3b8ffdecbef7a

# headers

x-res-version: R2402010 x-client-version: 1.10.30 x-device-type: android
x-idempotency-key: 5adf7d8440eb4dfd8797c20c7c513485 inspix-user-api-version:
1.0.0 Accept: application/json x-api-key:
4e769efa67d8f54be0b67e8f70ccb23d513a3c841191b6b2ba45ffc6fb498068 User-Agent:
inspix-android/1.10.30

# get real name

signature like: x-res-version: `R2402010@B/FicABV0d3BUb8PQHvXSsDwHw==`

spiltter: `@`(0x40u)

SimpleResver = `R2402010`

Hailstorm_Catalog_Manifest___ctor @ 0x286D4B8

v8 = Base64Decode(`B/FicABV0d3BUb8PQHvXSsDwHw==`) = `07F162700055D1DDC151BF0F407BD74AC0F01F`
checsum = BinReader__ReadULong(v8) = `07F162700055D1DD`
Seed = BinReader__ReadULong(v8) = `C151BF0F407BD74A`
Size = BinReader__ReadVLQ(v8) = Read(`C0F01F`) = `7F840`
IsQualified = 1

checksum = Manifest.Checksum = `07F162700055D1DD`
labelCrc = Crc64(Manifest.SimpleResver) = `E0D8673BCC7D7A04`
size = Manifest.Size = `07F840`

x = checksum + labelCrc + VLQ(size) = `07F162700055D1DDE0D8673BCC7D7A04C0F01F`
y = MD5(x) = `8C50F36F436CCE88E99EB741CF9E690A`
result = Base32__FromBytes(y) = `rripg32dnthir2m6w5a47htjbi`


# parseTransposedArray

size vlq: 16396 = 0x400c

CA 01 00 03 8C 80 01 00 00 00 00 00 00 00 00 00

magic number
CA 01

idk
00 03

vlq number
8C 80 01

data
0000...


# LocalDataStorageProvider__GenPrefsKeyForDomain:

```
string domain = "app"
bytes = Encoding.UTF8.GetBytes(domain)
sha1 = Cryptography.HMACSHA1.ComputeHash(bytes)
base64 = System.Convert.ToBase64String(sha1)
return base64
```

# Random File Name

1. Use `RNGCryptoServiceProvider` to generate an arbitrary 256 length byte array
2. Base64 encode these bytes, save it into persist storage `UnityEngine.PlayerPrefs`
3. LocalDataStorageProvider.Seed = random256Bytes
4. Use `Seed` to generate storageName:
  LocalDataStorageProvider.StorageName = Base32.FromBytes(MD5(Seed))    (md5 bytes, not string representation)


# 3 files in M directory

CC: CatalogCacheStorage
ID: IdentityStorage
Metadata SQLite



2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Entered decodeRaw
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  Decoding AssetName: 'raidrewards.csv'
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Calling decodeRawBk
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Entered Crc64Calc
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  crc seed is: [0]
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  crc length is (at): [F]
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  crc64 span str at: [0x6e8c123314]
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  crc64 str is: [raidrewards.tsv]
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  crc64 result is: [4592AD3507A9965D]
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Entered binWriteBytes
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  roSpanBytes at [70620BA3C8]
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Calling binWriteBytesBk
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Entered binWriteBytes
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Calling binWriteFlushBk
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  Flushed bytes: [c34ea77df4976cd8907096fa47bb97e61852305892494e3692ba0c7eb434f022c549c96cf7ca0ee1b6ba7f203b6c76e8679699ce9c44af7b1cb000173a51593853d0377e31d7b94c4592ad3507a9965d1ab3db55d004]
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Entered computeHash
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               I  Calling computeHashBk
2024-02-19 17:03:14.506  8377-31834 UNITYHOOK               usap64                               D  Computed hash bytes: [8f5ab496058c8dacff3bf703d6e0a0ac99923ba91626a44676d665e6e045c10c]
2024-02-19 17:03:14.507  8377-31834 UNITYHOOK               usap64                               I  Exist decodeRaw


8650505719232016025

seed
53d0377e31d7b94c
53 d0 37 7e 31 d7 b9 4c
crc64
4592ad3507a9965d
crc32
1ab3db55
size
d004
