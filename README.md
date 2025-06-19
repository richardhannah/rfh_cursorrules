# Go Development Cursor Rules

Shared Cursor rules for Go development projects. These rules provide consistent development practices across multiple Go projects.

## Files

- `architecture.md` - Clean architecture, repository pattern, dependency injection, package structure, naming conventions
- `database.md` - Database operations, model struct tags, validation
- `api-design.md` - REST conventions, DTOs, error handling, middleware patterns
- `testing.md` - Test organization, patterns, build tags, mocks, table-driven tests
- `development.md` - Error handling, code quality, logging, comments, function focus
- `security.md` - Parameterized queries, authentication, authorization, input sanitization, HTTPS
- `general.md` - Repository management, shared rules maintenance, git subtree best practices

## Usage

Add to your project using git subtree:

```bash
git subtree add --prefix=.cursorrules-shared https://github.com/richardhannah/rfh_cursorrules.git main --squash
```

## Updating Rules

To update rules in your project:
```bash
git subtree pull --prefix=.cursorrules-shared https://github.com/richardhannah/rfh_cursorrules.git main --squash
```

To contribute back to shared rules:
```bash
git subtree push --prefix=.cursorrules-shared https://github.com/richardhannah/rfh_cursorrules.git main
```

## Important: Immediate Push Rule

**Whenever any shared rules change** - immediately push the subtree to the shared rules repo to keep all projects in sync. See `general.md` for detailed workflow.
