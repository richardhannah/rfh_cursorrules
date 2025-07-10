param(
    [Parameter(Mandatory=$true)]
    [string]$RemoteDir
)

$localRulesDir = ".cursor/rules"
$localScriptsDir = ".cursor/scripts"

# Ensure local .cursor/rules directory exists
if (-not (Test-Path $localRulesDir)) {
    New-Item -ItemType Directory -Path $localRulesDir | Out-Null
}

# Ensure local .cursor/scripts directory exists
if (-not (Test-Path $localScriptsDir)) {
    New-Item -ItemType Directory -Path $localScriptsDir | Out-Null
}

$remoteRulesDir = Join-Path $RemoteDir ".cursor/rules"
$remoteScriptsDir = Join-Path $RemoteDir ".cursor/scripts"

# Check if remote .cursor/rules exists
if (-not (Test-Path $remoteRulesDir)) {
    Write-Error "Remote .cursor/rules directory not found at $remoteRulesDir"
    exit 1
}

# Helper function to sync shared folders
function Sync-SharedFolders {
    param(
        [string]$RemoteParent,
        [string]$LocalParent
    )
    $categories = Get-ChildItem -Path $RemoteParent -Directory -ErrorAction SilentlyContinue
    foreach ($cat in $categories) {
        $remoteShared = Join-Path $cat.FullName "shared"
        if (Test-Path $remoteShared) {
            $localCat = Join-Path $LocalParent $cat.Name
            $localShared = Join-Path $localCat "shared"
            if (-not (Test-Path $localCat)) {
                New-Item -ItemType Directory -Path $localCat | Out-Null
            }
            if (-not (Test-Path $localShared)) {
                New-Item -ItemType Directory -Path $localShared | Out-Null
            }
            # Symlink all files from remote shared to local shared
            $files = Get-ChildItem -Path $remoteShared -File
            foreach ($file in $files) {
                $localLink = Join-Path $localShared $file.Name
                if (Test-Path $localLink) {
                    Remove-Item $localLink -Force
                }
                New-Item -ItemType SymbolicLink -Path $localLink -Target $file.FullName | Out-Null
            }
        }
    }
}

# Sync shared folders in rules and scripts
echo "Syncing shared folders in .cursor/rules..."
Sync-SharedFolders -RemoteParent $remoteRulesDir -LocalParent $localRulesDir
echo "Syncing shared folders in .cursor/scripts..."
if (Test-Path $remoteScriptsDir) {
    # Special handling for .cursor/scripts/shared
    $remoteSharedScripts = Join-Path $remoteScriptsDir "shared"
    $localSharedScripts = Join-Path $localScriptsDir "shared"
    if (Test-Path $remoteSharedScripts) {
        if (-not (Test-Path $localSharedScripts)) {
            New-Item -ItemType Directory -Path $localSharedScripts | Out-Null
        }
        $files = Get-ChildItem -Path $remoteSharedScripts -File
        foreach ($file in $files) {
            $localLink = Join-Path $localSharedScripts $file.Name
            if (Test-Path $localLink) {
                Remove-Item $localLink -Force
            }
            New-Item -ItemType SymbolicLink -Path $localLink -Target $file.FullName | Out-Null
        }
    }
    Sync-SharedFolders -RemoteParent $remoteScriptsDir -LocalParent $localScriptsDir
}

Write-Output "Shared rule folders and symlinks have been synchronized." 