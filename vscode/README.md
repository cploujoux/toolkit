# Visual studio Code Beamlit

Extension for developers building agents to run on Beamlit platform.

## Features

Add an explorer view where you can:
- List Beamlit resources
  - Agents
  - Functions
  - Models
  - Environments
  - Policies
  - Integration
- Describe every resources

## Requirements

1. Install Beamlit CLI
```bash
curl -fsSL https://raw.githubusercontent.com/beamlit/toolkit/main/install.sh | BINDIR=$HOME/.local/bin sh
```
2. Login to Beamlit's Platform
```
bl login YOUR_WORKSPACE
```

## Known Issues

- Have to login with CLI before using the extension
- Cannot switch workspace

## Following extension guidelines

Ensure that you've read through the extensions guidelines and follow the best practices for creating your extension.

* [Extension Guidelines](https://code.visualstudio.com/api/references/extension-guidelines)
