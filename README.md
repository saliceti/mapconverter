# Map converter

Utility to convert data from a source and push to another source.

Example:
```
% cat input.json | mapconverter -d yaml -l stdin -s stdout
```

This command will:
- Pull data from stdin
- Load json data
- Convert to yaml data
- Push to stdout
