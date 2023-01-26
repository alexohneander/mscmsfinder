# msCMSFinder [![](https://github.com/alexohneander/mscmsfinder/workflows/Go/badge.svg)](https://github.com/alexohneander/mscmsfinder/actions)

### What is msCMSFinder?
msCMSFinder is a small API that should make it possible to quickly scan websites for their technology.

## Install
```bash
git clone https://github.com/alexohneander/mscmsfinder.git

go get .
```

## Usage

```bash
go run .
```

**API Request:**
```bash
curl http://127.0.0.1:3000/api/cms/find?url=http://example.de
```