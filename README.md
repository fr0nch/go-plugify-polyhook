[![Русский](https://img.shields.io/badge/Русский-%F0%9F%87%B7%F0%9F%87%BA-green?style=for-the-badge)](README_ru.md)

# Go bindings for the PolyHook Plugify plugin

This repository contains Go bindings for the [plugify-plugin-polyhook](https://github.com/untrustedmodders/plugify-plugin-polyhook) plugin. The bindings are automatically generated and kept in sync with the original plugin.

## API Documentation

Full API documentation can be found here on the [API Hub](https://api.plugify.net?file=https://raw.githubusercontent.com/untrustedmodders/plugify-plugin-polyhook/refs/heads/main/plugify-plugin-polyhook.pplugin.in) website.

## Installation

```bash
go get github.com/fr0nch/go-plugify-polyhook/v2
```

## Updates

This repository uses GitHub Actions to automatically check for updates in [plugify-plugin-polyhook](https://github.com/untrustedmodders/plugify-plugin-polyhook) every day at 2:00 UTC. When a new version is detected, the bindings are updated, and a new release is created.

## License

This Go module for Plugify is licensed under the [MIT License](LICENSE).