param(
    [Parameter(Mandatory=$true)]
    [string]$RuleFilePath,
    [Parameter(Mandatory=$true)]
    [ValidateSet('always','auto-attached','agent-requested','manual')]
    [string]$RuleType,
    [string]$Description = "A short description of the rule",
    [string[]]$Globs = @()
)

# Define YAML frontmatter templates
$frontmatters = @{
    'always' = "---`ndescription: $Description`nglobs:`nalwaysApply: true`n---`n"
    'auto-attached' = "---`ndescription: $Description`nglobs: [$(if ($Globs) { $Globs -join ', ' } else { '"*.ext"' })]`nalwaysApply: false`n---`n"
    'agent-requested' = "---`ndescription: $Description`nglobs:`nalwaysApply: false`n---`n"
    'manual' = "---`ndescription: $Description`nglobs:`nalwaysApply: false`n---`n"
}

if (-not (Test-Path $RuleFilePath)) {
    Write-Error "Rule file not found: $RuleFilePath"
    exit 1
}

# Read the existing content
$content = Get-Content $RuleFilePath -Raw

# Remove any existing YAML frontmatter (between --- lines at the top)
if ($content -match "(?s)^---.*?---\s*") {
    $content = $content -replace "(?s)^---.*?---\s*", ''
}

# Prepend the correct frontmatter
$frontmatter = $frontmatters[$RuleType]
Set-Content -Path $RuleFilePath -Value ("$frontmatter`n$content")

Write-Output "Prepended YAML frontmatter for rule type '$RuleType' to $RuleFilePath" 