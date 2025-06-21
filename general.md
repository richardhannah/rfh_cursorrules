# General Repository Management Rules

## Shared Rules Maintenance

### Immediate Push Rule
- **Whenever any shared rules change** - immediately push the subtree to the shared rules repo
- This ensures all projects stay in sync with the latest shared rules
- Command: `git subtree push --prefix=.cursorrules-shared git@github.com:richardhannah/rfh_cursorrules.git main`

### Workflow for Shared Rules Updates
1. Make changes to shared rules in `.cursorrules-shared/`
2. Commit the changes: `git commit -m "Update shared rules: [description]"`
3. **Immediately push to shared repo**: `git subtree push --prefix=.cursorrules-shared git@github.com:richardhannah/rfh_cursorrules.git main`
4. Continue with normal development

### Pulling Updates in Other Projects
- When working on other projects, regularly pull the latest shared rules
- Command: `git subtree pull --prefix=.cursorrules-shared git@github.com:richardhannah/rfh_cursorrules.git main --squash`
- Do this before starting new features to ensure you have the latest rules

## Documentation Organization and Maintenance

### Documentation Structure
- **`README.md`** - Project introduction and documentation index
- **`docs/`** - Main documentation directory
  - **`API.md`** - Complete API reference with endpoints, examples, error codes
  - **`ARCHITECTURE.md`** - System design, patterns, and architectural decisions
  - **`DEVELOPMENT.md`** - Setup, workflow, and development practices
  - **`TESTING.md`** - Comprehensive testing strategy and examples
  - **`decisions/`** - Design decision documentation
    - **`[DECISION_NAME].md`** - Individual decision records

### Documentation Update Rule
- **Whenever code changes are made** - update relevant documentation
- **API changes** require updates to `docs/API.md`
- **Architecture changes** require updates to `docs/ARCHITECTURE.md`
- **New features** require updates to `docs/DEVELOPMENT.md`
- **Testing changes** require updates to `docs/TESTING.md`
- **README.md** should be updated for any user-facing changes

### Major Changes Documentation Rule
- **Any major changes must be recorded in `docs/decisions/`**
- **Major changes include**:
  - Significant architecture changes
  - Breaking API changes
  - New design patterns or frameworks
  - Database schema changes
  - Security model changes
  - Performance optimization strategies

### Decision Documentation Format
```markdown
# [Decision Name]

## Context
Brief description of the problem or opportunity

## Decision
What was decided and why

## Consequences
- Positive consequences
- Negative consequences
- Trade-offs considered

## Implementation
How the decision was implemented

## Related
Links to related decisions or documentation
```

### Documentation Quality Standards
- **Keep documentation up-to-date** with code changes
- **Use clear, concise language** - avoid jargon when possible
- **Include practical examples** for all concepts
- **Cross-reference related documentation** where appropriate
- **Use consistent formatting** and structure
- **Include code examples** where relevant
- **Maintain version compatibility** information

## Automatic Testing and Committing

### Inactivity-Based Auto-Commit Rule
- **After 5 minutes of IDE inactivity** - automatically run unit tests
- **If all tests pass** - commit any pending changes with descriptive message
- **If tests fail** - do not commit, alert user to fix issues first
- This ensures code is regularly saved and tested without manual intervention

### Auto-Commit Workflow
1. Monitor IDE activity for 5-minute periods of inactivity
2. Run `go test ./...` to execute all unit tests
3. If tests pass:
   - Stage all changes: `git add .`
   - Commit with timestamp: `git commit -m "Auto-commit: [timestamp] - All tests passing"`
   - Optionally push if configured for auto-push
4. If tests fail:
   - Log the failure details
   - Alert user to review and fix failing tests
   - Do not commit changes

### Configuration
- Enable/disable auto-commit feature via project settings
- Configure test timeout and retry logic
- Set up notifications for test failures
- Customize commit message format

## Repository Management

### Git Subtree Best Practices
- Always use `--squash` when pulling to keep history clean
- Use descriptive commit messages for shared rule changes
- Keep shared rules focused and general (not project-specific)
- Test shared rules in multiple projects before finalizing

### Version Control
- Tag major versions of shared rules for stability
- Document breaking changes in shared rules
- Consider using semantic versioning for shared rules releases

### Collaboration
- Share the shared rules repo with team members
- Establish review process for shared rule changes
- Document any project-specific overrides or extensions 

## Important: Never Ignore Cursor Rules

### Critical Rule: Never Add to .gitignore
- **NEVER add `.cursorrules` to `.gitignore`** – Contains important project-specific rules and references
- **NEVER add `.cursorrules-shared/` to `.gitignore`** – Will break git subtree workflow
- Both files must remain tracked for the cursor rules system to function properly

### Why This Matters
- `.cursorrules` contains project-specific rules and references to shared rules
- `.cursorrules-shared/` must be tracked for git subtree to push/pull changes
- Ignoring either file will break the shared rules workflow
- These files are part of your project's documentation and should be version controlled 

## Build and Test Automation

### Makefile Test Target
- All projects that use `make` MUST have a `test` target that runs all available tests (unit, integration, etc).
- The `test` target should be the default target when running `make` with no arguments.
