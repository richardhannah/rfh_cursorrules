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