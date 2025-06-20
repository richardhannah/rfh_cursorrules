# PowerShell Development Rules

This document outlines the best practices and style guidelines for writing PowerShell scripts within this repository. Adhering to these rules ensures that scripts are consistent, readable, maintainable, and secure.

## 1. Script Structure and Naming

- **Verb-Noun Naming Convention**: All scripts and functions must follow the `Verb-Noun` naming convention (e.g., `Get-Process`, `New-Item`). This is a core PowerShell principle that enhances predictability and discoverability.
- **Comment-Based Help**: Every script and function must include comprehensive comment-based help. This enables `Get-Help` to provide automatic documentation.

  ```powershell
  <#
  .SYNOPSIS
      A brief, one-line summary of what the script or function does.
  .DESCRIPTION
      A more detailed description of the script's purpose and functionality.
  .PARAMETER ParameterName
      A description of what each parameter is for and what kind of value it expects.
  .EXAMPLE
      PS C:\> Get-MyService -Name "wuauserv"
      An example demonstrating how to use the script or function.
  .NOTES
      Any additional notes, including version, author, or specific requirements.
  .LINK
      A URL to a related resource or documentation.
  #>
  ```

## 2. Functions and Parameterization

- **Advanced Functions**: Use `[CmdletBinding()]` to create advanced functions. This provides support for common parameters like `-Verbose`, `-Debug`, and `-ErrorAction`, making your scripts behave like native cmdlets.
- **Mandatory Parameters**: Clearly define parameters that are required for the script to function.
  ```powershell
  [Parameter(Mandatory=$true, HelpMessage="Please provide a valid computer name.")]
  ```
- **Input Validation**: Use validation attributes to ensure parameter data is correct *before* the script's main logic runs. This is crucial for preventing errors and improving security.
  ```powershell
  param(
      [Parameter(Mandatory=$true)]
      [ValidateNotNullOrEmpty()]
      [string]$ComputerName,

      [Parameter(Mandatory=$true)]
      [ValidateSet("Running", "Stopped", "All")]
      [string]$Status
  )
  ```

## 3. Error Handling

- **Use `try/catch` for Terminating Errors**: For any command that could produce a terminating error, wrap it in a `try/catch` block to handle exceptions gracefully.
- **Force Terminating Errors**: Most cmdlets produce non-terminating errors by default. Use `-ErrorAction Stop` on a command to force it to produce a terminating error that a `catch` block can trap.
- **Global Error Preference**: You can set `$ErrorActionPreference = "Stop"` at the beginning of a script to make all errors terminating. Use this with care, as it changes the default behavior for the entire script.

  ```powershell
  # Prefer using -ErrorAction Stop on individual commands for clarity.
  try {
      Get-ChildItem -Path "C:\NonExistentPath" -ErrorAction Stop
  }
  catch {
      Write-Error "An error occurred while accessing the path: $_"
      # Perform cleanup actions or exit gracefully.
  }
  ```

## 4. Security Best Practices

- **No Hardcoded Credentials**: Never store passwords, API keys, or any other secrets directly in a script.
- **Use `PSCredential` for Interactive Scripts**: For functions that require credentials, use the `[System.Management.Automation.PSCredential]` type for parameters. This prompts the user for credentials securely.
- **Use Secure Storage for Automation**: For unattended scripts, retrieve credentials at runtime from a secure, encrypted store like the Windows Credential Manager, Azure Key Vault, or HashiCorp Vault.
- **Principle of Least Privilege**: Ensure that scripts run with only the permissions they absolutely need to perform their function.

## 5. Code Quality and Style

- **Run `Invoke-ScriptAnalyzer`**: Regularly analyze your code using the `PSScriptAnalyzer` module to check for best-practice violations. This should be part of your development and CI/CD workflow.
- **Be Explicit**: Avoid using aliases (`gci`, `?`, `%`) in scripts. Always use the full cmdlet and parameter names (`Get-ChildItem`, `Where-Object`, `ForEach-Object`) for clarity and long-term maintainability.
- **Consistent Formatting**: Maintain a consistent code style for indentation, brace placement (OTBS - One True Brace Style), and line length to improve readability.
- **Avoid `Write-Host` for Output**: `Write-Host` writes directly to the console and the output cannot be redirected or suppressed. To return data from a function, write the object directly to the pipeline. For logging and status messages, use `Write-Verbose`, `Write-Information`, or `Write-Debug`.

```powershell
# WRONG - This output cannot be captured
function Get-Greeting {
    Write-Host "Hello, World"
}

# RIGHT - This object can be captured, redirected, or formatted
function Get-Greeting {
    return "Hello, World"
}
$greeting = Get-Greeting
$greeting | Out-File -FilePath .\greeting.txt
``` 