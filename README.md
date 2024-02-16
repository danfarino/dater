# Dater

Simple CLI tool to parse and display dates and timestamps.

# Examples

### Show the current time

```
$ dater
Local: Fri, 16 Feb 2024 14:22:58 -0800
UTC:   Fri, 16 Feb 2024 22:22:58 +0000
Epoch sec: 1708122178
less than a minute ago; 0s; 0.00 days total
```

### Parse a Unix timestamp

```
$ dater 1702739272
Local: Sat, 16 Dec 2023 07:07:52 -0800
UTC:   Sat, 16 Dec 2023 15:07:52 +0000
Epoch sec: 1702739272
2 months ago; 1495h16m12s; 62.30 days total
```

### Calculate relative to current time

```
$ dater +3h
Local: Fri, 16 Feb 2024 17:24:22 -0800
UTC:   Sat, 17 Feb 2024 01:24:22 +0000
Epoch sec: 1708133062
about 3 hours from now; 3h0m0s; 0.12 days total
```

### Parse a time string

```
$ dater 'Thu, 24 Jun 2021 17:25:23 +0000'
Local: Thu, 24 Jun 2021 10:25:23 -0700
UTC:   Thu, 24 Jun 2021 17:25:23 +0000
Epoch sec: 1624555523
over 2 years ago; 23213h0m18s; 967.21 days total
```
