# Development environment

This repo uses 'just' as a 'make' replacement.

## Reset environment

In order to make these settings permanent, logout or reboot your PC.

## Installing tools

Execute go tools:

```bash
just tools
```

# Development workflow

## On a rebase

Initialise tools and modules

```bash
just tools
just qa
```
## Changing code

Edit or add code or other development activity.

Quality check code:

```bash
just qa
```

