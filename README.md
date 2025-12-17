# plugin-github

Official GitHub plugin for [Relicta](https://github.com/relicta-tech/relicta) - Create GitHub releases and upload assets.

## Installation

```bash
relicta plugin install github
```

Or install from source:

```bash
git clone https://github.com/relicta-tech/plugin-github
cd plugin-github
go build -o github
relicta plugin install ./github
```

## Configuration

Add to your `release.config.yaml`:

```yaml
plugins:
  - name: github
    enabled: true
    config:
      # Optional: override repository (auto-detected from git remote)
      owner: "your-org"
      repo: "your-repo"

      # Optional: create as draft release
      draft: false

      # Optional: mark as prerelease
      prerelease: false

      # Optional: use GitHub's auto-generated release notes
      generate_release_notes: false

      # Optional: files to upload as release assets
      assets:
        - "dist/*.tar.gz"
        - "dist/*.zip"
        - "dist/checksums.txt"

      # Optional: create a discussion for the release
      discussion_category: "Releases"
```

## Authentication

The plugin requires a GitHub token with `repo` permissions. Set it via:

1. Environment variable (recommended):
   ```bash
   export GITHUB_TOKEN=ghp_xxxx
   # or
   export GH_TOKEN=ghp_xxxx
   ```

2. Configuration:
   ```yaml
   plugins:
     - name: github
       config:
         token: "ghp_xxxx"  # Not recommended for version control
   ```

## Hooks

This plugin responds to the following hooks:

| Hook | Behavior |
|------|----------|
| `post-publish` | Creates GitHub release and uploads assets |
| `on-success` | Logs success message |
| `on-error` | Acknowledges failure |

## Outputs

After execution, the plugin provides these outputs:

| Output | Description |
|--------|-------------|
| `release_id` | GitHub release ID |
| `release_url` | URL to the release page |
| `tag_name` | Git tag name |

## Development

### Building

```bash
go build -o github
```

### Testing with Relicta

```bash
# Install locally
relicta plugin install ./github
relicta plugin enable github

# Test with dry-run
relicta publish --dry-run
```

### Dev mode with live reload

```bash
relicta plugin dev --watch
```

## License

MIT License - see [LICENSE](LICENSE) for details.
