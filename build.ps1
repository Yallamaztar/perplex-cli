param(
    [string]$Output = "status.exe"
)

Write-Host "Building $Output..."

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "Go is not found in PATH. Please install Go or add it to PATH."
    exit 1
}

try {
    & go build -trimpath -ldflags="-s -w" -gcflags="all=-B" -o $Output
} catch {
    Write-Error "Build failed: $_"
    exit 1
}

if (-not (Test-Path (Join-Path (Get-Location).Path $Output))) {
    Write-Error "Build did not produce $Output"
    exit 1
}

$dir = (Get-Location).Path
$userPath = [Environment]::GetEnvironmentVariable("Path","User")

if (-not $userPath) { $userPath = "" }

if ($userPath -notmatch [Regex]::Escape($dir)) {
    if ([string]::IsNullOrEmpty($userPath)) { $newPath = $dir } else { $newPath = "$userPath;$dir" }
    setx Path $newPath | Out-Null
    Write-Host "Added $dir to user PATH. Restart terminals to apply."
} else {
    Write-Host "User PATH already contains $dir"
}

Write-Host "Done."
Write-Host "You can now run 'status.exe (-l | -loop)' from any terminal."
